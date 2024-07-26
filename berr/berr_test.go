package berr

import (
	"errors"
	"fmt"
	"testing"
)

type MyTestErr struct {
	BaseErr[MyTestErr]
}

func myPrint(prefix string, t any) {
	fmt.Printf("%v: type %T add %p\n", prefix, t, t)

}
func TestBase(t *testing.T) {
	var myErr MyTestErr

	myPrint("1", &myErr)
	myPrint("2", &myErr.BaseErr)
	myPrint("3", myErr.BaseErr.t)

	errOld := errors.New("this is old error")
	errNew := errors.New("this is new error")
	errRev := myErr.InitAddr(&myErr).SetErr(errOld).SetErr(errNew).Err()

	myPrint("errOld", errOld)
	myPrint("errNew", errNew)
	myPrint("errRev", errRev)
}
