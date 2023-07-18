package pattern

/*
	Паттерн "Посетитель" (Visitor pattern) является поведенческим паттерном проектирования, который позволяет добавлять 
	новые операции к объектам без изменения их классов. Он позволяет разделить алгоритмы от объектной структуры, на которой 
	они оперируют.

Применимость паттерна "Посетитель" возникает в ситуациях, когда у объектной структуры есть множество различных операций, 
которые могут быть добавлены или изменены независимо от самих объектов. Паттерн "Посетитель" позволяет добавлять новые операции, 
создавая новые классы-посетители, без изменения классов объектов.

Преимущества использования паттерна "Посетитель" включают:

1.	Расширяемость операций. 
2.	Сохранение инкапсуляции объектов. 
3.	Увеличение гибкости системы. 

Некоторые возможные недостатки паттерна "Посетитель" включают:

1.	Усложнение структуры системы. 
2.  Нарушение инкапсуляции объектов. 

Примеры использования паттерна "Посетитель" в реальных приложениях включают:


Визуализация объектной структуры. Паттерн "Посетитель" может быть использован для обхода и отображения объектной структуры, 
где каждый класс-посетитель представляет различные методы визуализации для разных типов объектов.
*/

import "fmt"


type Customer interface {
	Accept(Visitor)
}

type Visitor interface {
	Visit(Customer)
}


// структура соответствует интерфейсу Customer
type CustomerCol struct {
	customers []Customer
}

func (c *CustomerCol) Add(customer Customer) {
	c.customers = append(c.customers, customer)
}

func (c *CustomerCol) Accept(visitor Visitor) {
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}

// структура соответствует интерфейсу Customer
type EnterpriseCustomer struct {
	name string
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{
		name: name,
	}
}

func (c *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

// структура соответствует интерфейсу Customer
type IndividualCustomer struct {
	name string
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{
		name: name,
	}
}

func (c *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

// структура соответствует интерфейсу Visitor
type ServiceRequestVisitor struct{}

// В зависимости от входящей структуры выводим в stdIn
func (*ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("serving enterprise customer %s\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("serving individual customer %s\n", c.name)
	}
}

func VisitorFunc() {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Add(NewIndividualCustomer("bob"))
	c.Accept(&ServiceRequestVisitor{})

}