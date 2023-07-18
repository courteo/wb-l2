package pattern

/*
Паттерн "Фабричный метод" (Factory Method pattern) является порождающим паттерном проектирования, который предоставляет интерфейс 
для создания объектов, но позволяет субклассам выбирать класс создаваемого объекта. Этот паттерн делегирует процесс создания объектов 
субклассам, что обеспечивает гибкость и локальное изменение поведения создания объектов.

Применимость паттерна "Фабричный метод" возникает, когда у вас есть суперкласс, определяющий общий интерфейс создания объектов, 
и несколько субклассов, реализующих этот интерфейс и предоставляющих различные реализации создания объектов. Паттерн "Фабричный метод"
 позволяет выбрать подходящий субкласс для создания объектов в зависимости от контекста.

Преимущества использования паттерна "Фабричный метод" включают:

1.	Гибкость и расширяемость. 
2.	Локальная модификация создания объектов. 
3.	Разделение ответственности. 

Некоторые возможные недостатки паттерна "Фабричный метод" включают:

1.	Усложнение структуры системы. 
2.	Зависимость от субклассов.

Примеры использования паттерна "Фабричный метод" в реальных приложениях включают:

Создание различных типов продуктов. Например, в системе заказа продуктов онлайн можно использовать фабричный метод 
для создания различных типов продуктов, таких как одежда, электроника или книги, с использованием соответствующих субклассов-фабрик.
*/

import "fmt"

// интерфейс с нужным функционалом
type SomeProduct interface {
	Use()
}

// ConcreteSomeProductA удовлетворяет интерфейсу SomeProduct
type ConcreteSomeProductA struct{}


func (p *ConcreteSomeProductA) Use() {
	fmt.Println("Using SomeProduct A")
}

// ConcreteSomeProductB удовлетворяет интерфейсу SomeProduct
type ConcreteSomeProductB struct{}


func (p *ConcreteSomeProductB) Use() {
	fmt.Println("Using SomeProduct B")
}

// интерфейс создателей
type Creator interface {
	CreateSomeProduct() SomeProduct
}

// ConcreteCreatorA  удовляетворяет интерфейсу Creator
type ConcreteCreatorA struct{}


func (c *ConcreteCreatorA) CreateSomeProduct() SomeProduct {
	return &ConcreteSomeProductA{}
}

// ConcreteCreatorB  удовляетворяет интерфейсу Creator
type ConcreteCreatorB struct{}


func (c *ConcreteCreatorB) CreateSomeProduct() SomeProduct {
	return &ConcreteSomeProductB{}
}

func FactoryMethodFunc() {

	creatorA := &ConcreteCreatorA{}

	// Create SomeProduct A
	SomeProductA := creatorA.CreateSomeProduct()
	SomeProductA.Use()


	creatorB := &ConcreteCreatorB{}

	// Create SomeProduct B
	SomeProductB := creatorB.CreateSomeProduct()
	SomeProductB.Use()
}