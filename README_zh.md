# base-error

流水线模式设计，一种主动/集中错误处理和链式调用封装方案。

## 特性

- 主动/集中错误处理
- 链式调用

## 适用场景

有大量组合操作的中间件封装，方便中间件的调用者集中处理错误，以及使用链式调用，提升代码可读性和开发效率。

## 效果对比

普通模式和base-error模式的对比示例，可以参考[examples](https://github.com/yuanboshe/base-error/tree/master/examples)中的例子。

普通模式下，调用者一般是单方法调用，被动错误处理。

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

base-error模式下，调用者可以主动/集中处理错误，可以实现链式调用，增强可读性和提升开发效率。

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

## 使用规范

参考：[berr_example_test.go](https://github.com/yuanboshe/base-error/blob/master/berr/berr_example_test.go)

### 中间件封装

引包：`import "github.com/yuanboshe/base-error/berr"`

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

要点：
1. 所有中间件封装都继承自`BaseErr[T]`，其中`T`为中间件struct类型。
2. 所有方法都要以`if p.Err() != nil { return p }`开头，用于检查错误，跳过后面的操作。
3. 所有内部错误处理，都以`p.SetErr()`导入到`err`成员变量暂存。

### 中间件调用

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

要点：
1. 所有中间件对象都需要`InitAddr()`进行对象地址初始化。
2. 通过`Err()`检查错误，并对错误进行处理后，如果后续还要继续使用中间件对象，则需要`SetErr(nil)`清除错误。

## 设计思想

一切从用户（调用者）使用便利的角度出发，通过用户主动/集中错误处理，减少错误处理代码的编写，提升可读性和开发效率；通过链式调用，进一步增强代码的可读性和开发效率。

更详细的主动/集中错误处理设计思想参考[一种优雅的Golang error设计模式](https://blog.pz1.top/article/base-error-elegant-golang-error-design-pattern)，流水线模式和链式调用设计思想参考[Golang流水线模式与链式调用](https://blog.pz1.top/article/base-error-pipeline-pattern-chaining-calls)。
