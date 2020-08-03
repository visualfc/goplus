// Package fmt provide Go+ "fmt" package, as "fmt" package in Go.
package fmt

import (
	fmt "fmt"
	io "io"

	gop "github.com/goplus/gop"
)

func execErrorf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0 := fmt.Errorf(args[0].(string), args[1:]...)
	p.Ret(arity, ret0)
}

func toType0(v interface{}) fmt.State {
	if v == nil {
		return nil
	}
	return v.(fmt.State)
}

func execmFormatterFormat(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	args[0].(fmt.Formatter).Format(toType0(args[1]), args[2].(rune))
	p.PopN(3)
}

func toType1(v interface{}) io.Writer {
	if v == nil {
		return nil
	}
	return v.(io.Writer)
}

func execFprint(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0, ret1 := fmt.Fprint(toType1(args[0]), args[1:]...)
	p.Ret(arity, ret0, ret1)
}

func execFprintf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0, ret1 := fmt.Fprintf(toType1(args[0]), args[1].(string), args[2:]...)
	p.Ret(arity, ret0, ret1)
}

func execFprintln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0, ret1 := fmt.Fprintln(toType1(args[0]), args[1:]...)
	p.Ret(arity, ret0, ret1)
}

func toType2(v interface{}) io.Reader {
	if v == nil {
		return nil
	}
	return v.(io.Reader)
}

func execFscan(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0, ret1 := fmt.Fscan(toType2(args[0]), args[1:]...)
	p.Ret(arity, ret0, ret1)
}

func execFscanf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0, ret1 := fmt.Fscanf(toType2(args[0]), args[1].(string), args[2:]...)
	p.Ret(arity, ret0, ret1)
}

func execFscanln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0, ret1 := fmt.Fscanln(toType2(args[0]), args[1:]...)
	p.Ret(arity, ret0, ret1)
}

func execmGoStringerGoString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret0 := args[0].(fmt.GoStringer).GoString()
	p.Ret(1, ret0)
}

func execPrint(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0, ret1 := fmt.Print(args...)
	p.Ret(arity, ret0, ret1)
}

func execPrintf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0, ret1 := fmt.Printf(args[0].(string), args[1:]...)
	p.Ret(arity, ret0, ret1)
}

func execPrintln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0, ret1 := fmt.Println(args...)
	p.Ret(arity, ret0, ret1)
}

func execScan(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0, ret1 := fmt.Scan(args...)
	p.Ret(arity, ret0, ret1)
}

func execmScanStateRead(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret0, ret1 := args[0].(fmt.ScanState).Read(args[1].([]byte))
	p.Ret(2, ret0, ret1)
}

func execmScanStateReadRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret0, ret1, ret2 := args[0].(fmt.ScanState).ReadRune()
	p.Ret(1, ret0, ret1, ret2)
}

func execmScanStateSkipSpace(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	args[0].(fmt.ScanState).SkipSpace()
	p.PopN(1)
}

func execmScanStateToken(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret0, ret1 := args[0].(fmt.ScanState).Token(args[1].(bool), args[2].(func(rune) bool))
	p.Ret(3, ret0, ret1)
}

func execmScanStateUnreadRune(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret0 := args[0].(fmt.ScanState).UnreadRune()
	p.Ret(1, ret0)
}

func execmScanStateWidth(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret0, ret1 := args[0].(fmt.ScanState).Width()
	p.Ret(1, ret0, ret1)
}

func execScanf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0, ret1 := fmt.Scanf(args[0].(string), args[1:]...)
	p.Ret(arity, ret0, ret1)
}

func execScanln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0, ret1 := fmt.Scanln(args...)
	p.Ret(arity, ret0, ret1)
}

func toType3(v interface{}) fmt.ScanState {
	if v == nil {
		return nil
	}
	return v.(fmt.ScanState)
}

func execmScannerScan(_ int, p *gop.Context) {
	args := p.GetArgs(3)
	ret0 := args[0].(fmt.Scanner).Scan(toType3(args[1]), args[2].(rune))
	p.Ret(3, ret0)
}

func execSprint(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0 := fmt.Sprint(args...)
	p.Ret(arity, ret0)
}

func execSprintf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0 := fmt.Sprintf(args[0].(string), args[1:]...)
	p.Ret(arity, ret0)
}

func execSprintln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0 := fmt.Sprintln(args...)
	p.Ret(arity, ret0)
}

func execSscan(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0, ret1 := fmt.Sscan(args[0].(string), args[1:]...)
	p.Ret(arity, ret0, ret1)
}

func execSscanf(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0, ret1 := fmt.Sscanf(args[0].(string), args[1].(string), args[2:]...)
	p.Ret(arity, ret0, ret1)
}

func execSscanln(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	ret0, ret1 := fmt.Sscanln(args[0].(string), args[1:]...)
	p.Ret(arity, ret0, ret1)
}

func execmStateFlag(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret0 := args[0].(fmt.State).Flag(args[1].(int))
	p.Ret(2, ret0)
}

func execmStatePrecision(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret0, ret1 := args[0].(fmt.State).Precision()
	p.Ret(1, ret0, ret1)
}

func execmStateWidth(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret0, ret1 := args[0].(fmt.State).Width()
	p.Ret(1, ret0, ret1)
}

func execmStateWrite(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret0, ret1 := args[0].(fmt.State).Write(args[1].([]byte))
	p.Ret(2, ret0, ret1)
}

func execmStringerString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret0 := args[0].(fmt.Stringer).String()
	p.Ret(1, ret0)
}

// I is a Go package instance.
var I = gop.NewGoPackage("fmt")

func init() {
	I.RegisterFuncs(
		I.Func("(Formatter).Format", (fmt.Formatter).Format, execmFormatterFormat),
		I.Func("(GoStringer).GoString", (fmt.GoStringer).GoString, execmGoStringerGoString),
		I.Func("(ScanState).Read", (fmt.ScanState).Read, execmScanStateRead),
		I.Func("(ScanState).ReadRune", (fmt.ScanState).ReadRune, execmScanStateReadRune),
		I.Func("(ScanState).SkipSpace", (fmt.ScanState).SkipSpace, execmScanStateSkipSpace),
		I.Func("(ScanState).Token", (fmt.ScanState).Token, execmScanStateToken),
		I.Func("(ScanState).UnreadRune", (fmt.ScanState).UnreadRune, execmScanStateUnreadRune),
		I.Func("(ScanState).Width", (fmt.ScanState).Width, execmScanStateWidth),
		I.Func("(Scanner).Scan", (fmt.Scanner).Scan, execmScannerScan),
		I.Func("(State).Flag", (fmt.State).Flag, execmStateFlag),
		I.Func("(State).Precision", (fmt.State).Precision, execmStatePrecision),
		I.Func("(State).Width", (fmt.State).Width, execmStateWidth),
		I.Func("(State).Write", (fmt.State).Write, execmStateWrite),
		I.Func("(Stringer).String", (fmt.Stringer).String, execmStringerString),
	)
	I.RegisterFuncvs(
		I.Funcv("Errorf", fmt.Errorf, execErrorf),
		I.Funcv("Fprint", fmt.Fprint, execFprint),
		I.Funcv("Fprintf", fmt.Fprintf, execFprintf),
		I.Funcv("Fprintln", fmt.Fprintln, execFprintln),
		I.Funcv("Fscan", fmt.Fscan, execFscan),
		I.Funcv("Fscanf", fmt.Fscanf, execFscanf),
		I.Funcv("Fscanln", fmt.Fscanln, execFscanln),
		I.Funcv("Print", fmt.Print, execPrint),
		I.Funcv("Printf", fmt.Printf, execPrintf),
		I.Funcv("Println", fmt.Println, execPrintln),
		I.Funcv("Scan", fmt.Scan, execScan),
		I.Funcv("Scanf", fmt.Scanf, execScanf),
		I.Funcv("Scanln", fmt.Scanln, execScanln),
		I.Funcv("Sprint", fmt.Sprint, execSprint),
		I.Funcv("Sprintf", fmt.Sprintf, execSprintf),
		I.Funcv("Sprintln", fmt.Sprintln, execSprintln),
		I.Funcv("Sscan", fmt.Sscan, execSscan),
		I.Funcv("Sscanf", fmt.Sscanf, execSscanf),
		I.Funcv("Sscanln", fmt.Sscanln, execSscanln),
	)
}
