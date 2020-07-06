/*
 Copyright 2020 The GoPlus Authors (goplus.org)

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/qiniu/goplus/cmd/qexp/gopkg"
)

var (
	flagExportDir string
	buildPlugin   bool
)

func init() {
	flag.StringVar(&flagExportDir, "outdir", "", "optional set export lib path, default is $GoPlusRoot/lib path")
	flag.BoolVar(&buildPlugin, "plugin", false, "optional build lib to plugin, install to $HOME/.goplus/lib/")
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "Usage: qexp [-outdir <outRootDir>] [packages]\n")
		flag.PrintDefaults()
		return
	}
	var root string
	if flagExportDir != "" {
		root = flagExportDir
	} else {
		var err error
		root, err = goplusRoot()
		if err != nil {
			fmt.Fprintln(os.Stderr, "find goplus root failed:", err)
			os.Exit(-1)
		}
	}
	for _, pkgPath := range flag.Args() {
		buildPkg(pkgPath, root, buildPlugin)
	}
}

func buildPkg(pkgPath string, root string, plugin bool) {
	pkg, err := gopkg.Import(pkgPath)
	if err != nil {
		log.Printf("import pkg %q error: %v\n", pkgPath, err)
		return
	}
	outdir, err := exportPkg(pkg, root)
	if err != nil {
		log.Printf("export pkg %q error: %v\n", pkgPath, err)
		return
	}
	log.Printf("export pkg %q success: %v", pkgPath, outdir)
	if !plugin {
		return
	}
	ofile, err := exportPlugin(pkg, root)
	if err != nil {
		log.Printf("insall plugin %q error: %v\n", pkgPath, err)
		return
	}
	log.Printf("insall plugin %q success: %v", pkgPath, ofile)
}

func exportPkg(pkg *types.Package, root string) (string, error) {
	var buf bytes.Buffer
	err := gopkg.ExportPackage(pkg, &buf)
	if err != nil {
		return "", err
	}
	data, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(buf.String())
		return "", err
	}
	dir := filepath.Join(root, "lib", pkg.Path())
	os.MkdirAll(dir, 0777)
	err = ioutil.WriteFile(filepath.Join(dir, "gomod_export.go"), data, 0666)
	if err != nil {
		return "", err
	}
	return dir, nil
}

const gopkgExportHeader = `// Package %s provide Go+ "%s" package, as "%s" package in Go.
package main

import (
`

const gopkgExportFooter = `)
`

var (
	pkglib = "github.com/qiniu/goplus/lib"
)

func exportPlugin(pkg *types.Package, root string) (string, error) {
	var buf bytes.Buffer
	pkgName, pkgPath := pkg.Name(), pkg.Path()
	fmt.Fprintf(&buf, gopkgExportHeader, pkgName, pkgPath, pkgPath)
	fmt.Fprintf(&buf, "\t_ %q\n", pkglib+"/"+pkgPath)
	fmt.Fprintf(&buf, gopkgExportFooter)
	fmt.Fprintf(&buf, "func main(){}\n")
	data, err := format.Source(buf.Bytes())
	if err != nil {
		return "", err
	}
	dir := filepath.Join(root, "plugin", pkg.Path())
	os.MkdirAll(dir, 0777)
	err = ioutil.WriteFile(filepath.Join(dir, "gomod_export.go"), data, 0666)
	if err != nil {
		return "", err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	pkgDir, _ := filepath.Split(pkgPath)
	installDir := filepath.Join(home, ".goplus/lib", pkgDir)
	os.MkdirAll(installDir, 0777)
	gobin, err := exec.LookPath("go")
	if err != nil {
		return "", err
	}
	ofile := filepath.Join(installDir, pkgName+".dylib")
	cmd := exec.Command(gobin, "build", "-v", "-buildmode", "plugin", "-o", ofile)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	return ofile, nil
}

func goplusRoot() (root string, err error) {
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	for {
		modfile := filepath.Join(dir, "go.mod")
		if hasFile(modfile) {
			if isGoplus(modfile) {
				return dir, nil
			}
			return "", errors.New("current directory is not under goplus root")
		}
		next := filepath.Dir(dir)
		if dir == next {
			return "", errors.New("go.mod not found, please run under goplus root")
		}
		dir = next
	}
}

func isGoplus(modfile string) bool {
	b, err := ioutil.ReadFile(modfile)
	return err == nil && bytes.HasPrefix(b, goplusPrefix)
}

var (
	goplusPrefix = []byte("module github.com/qiniu/goplus")
)

func hasFile(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.Mode().IsRegular()
}

func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}
	if "windows" == runtime.GOOS {
		return homeWindows()
	}
	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}
	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}
	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}
	return home, nil
}

// -----------------------------------------------------------------------------
