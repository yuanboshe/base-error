package berr

import (
	"fmt"
)

type MyErr struct {
	BaseErr[MyErr]
}

func (p *MyErr) DoSomethingNoError() *MyErr {
	// if err is not nil, return itself with err message
	if p.Err() != nil {
		return p
	}

	// do something no error ...
	fmt.Println("do something no error ...")

	// if success, return itself without err message
	return p
}

func (p *MyErr) DoSomethingWithError() *MyErr {
	// if err is not nil, return itself with err message
	if p.Err() != nil {
		return p
	}

	// do something with error ...
	fmt.Println("do something with error ...")

	// error handle in do something ...
	if err := fmt.Errorf("if some error need handle"); err != nil {
		return p.SetErr(err)
	}

	// if success, return itself without err message
	return p
}

func ExampleBaseErr() {

	var myErr MyErr
	myErr.InitAddr(&myErr)

	if err := myErr.DoSomethingNoError().DoSomethingWithError().DoSomethingNoError().Err(); err != nil {
		fmt.Println("error in example main:", err)
	}

	// output:
	// do something no error ...
	// do something with error ...
	// error in example main: if some error need handle
}
