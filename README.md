# base-error

Pipeline pattern design, an active/centralized error handling and chaining call encapsulation solution.

## Features

- Active/Centralized Error Handling
- Chaining Calls

## Applicable Scenarios

Encapsulation of middleware with a large number of composite operations, facilitating centralized error handling for middleware callers, and using chaining calls to enhance code readability and development efficiency.

## Comparison of Effects

For a comparison example between the normal mode and the base-error mode, please refer to the examples in [examples](https://github.com/yuanboshe/base-error/tree/master/examples)。

In the normal mode, callers generally make single-method calls with passive error handling.

```go
func main() {
	robot := examples.NewRobot("robot_without_berr")

	err := robot.SetWheels(4)
	if err != nil {
		// do something error handling
		panic(err)
	}

	err = robot.Paint("white")
	if err != nil {
		// do something error handling
		panic(err)
	}

	err = robot.Charge(0.8)
	if err != nil {
		// do something error handling
		panic(err)
	}

	report, err := robot.GetReport()
	if err != nil {
		// do something error handling
		panic(err)
	}

	fmt.Println(report)
}
```

In the base-error mode, callers can actively/centrally handle errors, and can implement chaining calls to enhance readability and improve development efficiency.

```go
func main() {
	robot := examples.NewRobot("robot_with_berr")

	report := robot.SetWheels(4).Paint("white").Charge(0.8).GetReport()
	if robot.Err() != nil {
		// do something error handling
		panic(robot.Err())
	}

	fmt.Println(report)
}
```

## Usage Specifications

Reference: [berr_example_test.go](https://github.com/yuanboshe/base-error/blob/master/berr/berr_example_test.go)

### Middleware Encapsulation

Import package: `import "github.com/yuanboshe/base-error/berr"`

```go
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
```

Key Points:
1. All middleware encapsulations inherit from `BaseErr[T]`, where `T` is the middleware struct type.
2. All methods should start with `if p.Err() != nil { return p }` to check for errors and skip subsequent operations.
3. All internal error handling should be imported into the `err` member variable for temporary storage using `p.SetErr()`.

### Middleware Invocation

```go
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
```

Key Points:
1. All middleware objects need to be initialized with `InitAddr()` for object address initialization.
2. Check for errors using `Err()`, and after handling the error, if you want to continue using the middleware object, clear the error with `SetErr(nil)`.

## Design Philosophy

Everything starts from the convenience of the user (caller). By actively/centrally handling errors, it reduces the writing of error handling code, enhances readability, and improves development efficiency. Chaining calls further enhance code readability and development efficiency.

For more detailed active/centralized error handling design philosophy, please refer to [An Elegant Golang Error Design Pattern](https://blog.pz1.top/en/article/base-error-elegant-golang-error-design-pattern), and for the design philosophy of pipeline pattern and chaining calls, please refer to [Golang Pipeline Pattern and Chaining Calls](https://blog.pz1.top/en/article/base-error-pipeline-pattern-chaining-calls).
