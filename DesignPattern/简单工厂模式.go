package main

import "fmt"

//首先将需要创建的各种不同对象（例如各种不同的 Chart 对象）的相关代码封装到不同的类中，这些类称为具体产品类，而将它们公共的代码进行抽象和提取后封装在一个抽象产品类中，每一个具体产品类都是抽象产品类的子类；然后提供一个工厂类用于创建各种产品，在工厂类中提供一个创建产品的工厂方法，该方法可以根据所传入的参数不同创建不同的具体产品对象；客户端只需调用工厂类的工厂方法并传入相应的参数即可得到一个产品对象。

//简单工厂模式定义如下：
//简单工厂模式（Simple Factory Pattern）：定义一个工厂类，它可以根据参数的不同返回不同类的实例，被创建的实例通常都具有共同的父类。因为在简单工厂模式中用于创建实例的方法是静态（static）方法，因此简单工厂模式又被称为静态工厂方法（Static Factory Method）模式，它属于类创建型模式。

type Factory struct {
}

type Product interface {
	Create()
}

type Product1 struct {
}

func (p *Product1) Create() {
	fmt.Println("this is product1")
}

type Product2 struct {
}

func (p *Product2) Create() {
	fmt.Println("this is product1")
}

func (f *Factory) Generate(name string) Product {
	switch name {
	case "product1":
		return &Product1{}
	case "product2":
		return &Product2{}
	default:
		return nil
	}
}

func main() {
	factory := Factory{}

	p1 := factory.Generate("product1")
	p1.Create()

	p2 := factory.Generate("product2")
	p2.Create()
}
