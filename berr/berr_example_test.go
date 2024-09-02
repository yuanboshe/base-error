package berr

import (
	"fmt"
)

// ------------------- Middleware Encapsulation Example -------------------

type Product struct {
	message string

	// Point1: base on BaseErr
	BaseErr[Product]
}

func (p *Product) DoSomethingWithError() *Product {
	// Point2: check error in the head of method
	if p.Err() != nil {
		return p
	}

	// do something with error ...
	fmt.Println("do something with error ...")
	p.message = "message: do something with error ..."

	// error handle in do something ...
	if err := fmt.Errorf("if some error need handle"); err != nil {
		// Point3: use SetErr() to set error
		return p.SetErr(err)
	}

	// if success, return itself without err message
	return p
}

func (p *Product) DoSomethingNoError() *Product {
	if p.Err() != nil {
		return p
	}

	// do something no error ...
	fmt.Println("do something no error ...")
	p.message = "message: do something no error ..."

	return p
}

func (p *Product) GetMessage() string {
	if p.Err() != nil {
		return ""
	}

	return p.message
}

// ------------------- Middleware Invocation Example -------------------

func ExampleBaseErr() {
	var product Product
	// Point1: user InitAddr() to init the address
	product.InitAddr(&product)

	if product.DoSomethingNoError().DoSomethingWithError().DoSomethingNoError().Err() != nil {
		// deal with error
		fmt.Println("error in example main1:", product.Err())

		// Point2: use SetErr(nil) clear the error after handle the error, if you want to continue use the object
		product.SetErr(nil)
	}

	message := product.DoSomethingNoError().GetMessage()
	fmt.Println(message)
	if product.Err() != nil {
		fmt.Println("error in example main2:", product.Err())
	}

	// output:
	// do something no error ...
	// do something with error ...
	// error in example main1: if some error need handle
	// do something no error ...
	// message: do something no error ...
}
