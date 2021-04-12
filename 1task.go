package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

type ErrorWithTrace struct {
	text string
	trace string
	timePoint string
}

func NewErr(text string) error {
	return &ErrorWithTrace{
		text: text,
		trace: string(debug.Stack()),
		timePoint: time.Now().String(),
	}
}

func (e *ErrorWithTrace) Error() string {
	return fmt.Sprintf("error: %s\ntrace:\n%s\ntime:%s", e.text, e.trace, e.timePoint)
}

func recoverPanic() {
	if err := recover(); err !=nil {
		err = NewErr("my error description below")
		fmt.Println(err)
	}
}

func makePanicSituation() {
	a,b:=5,0
	fmt.Print(a/b)
}

func main() {

	defer recoverPanic()

	makePanicSituation()

}