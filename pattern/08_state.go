package pattern

/*
Паттерн "Состояние" (State pattern) является поведенческим паттерном проектирования, который позволяет объекту 
изменять свое поведение в зависимости от своего внутреннего состояния. Он позволяет инкапсулировать различные состояния 
объекта в отдельные классы и делает их взаимозаменяемыми.

Применимость паттерна "Состояние" возникает, когда у объекта есть множество различных состояний, и его поведение зависит 
от текущего состояния. Паттерн "Состояние" позволяет объекту динамически изменять свое состояние и связанное с ним поведение, 
без изменения самого объекта.

Преимущества использования паттерна "Состояние" включают:

1.	Упрощение кода.
2.	Гибкость и расширяемость. 
3.	Улучшение поддержки кода. 

Некоторые возможные недостатки паттерна "Состояние" включают:

1.	Увеличение числа классов. 
2.	Возможность несогласованного состояния. 

Примеры использования паттерна "Состояние" в реальных приложениях включают:

Обработка заказов. В системе управления заказами паттерн "Состояние" может быть использован для моделирования 
жизненного цикла заказа, с различными состояниями, такими как "новый", "обрабатывается", "отгружен" и "завершен". 
Каждое состояние представляет различные действия и переходы между состояниями.
*/

import "fmt"


type State interface {
	Handle()
}

// ConcreteStateA удовлетворяет интерфейсу State
type ConcreteStateA struct{}


func (s *ConcreteStateA) Handle() {
	fmt.Println("Handling state A")
}

// ConcreteStateB удовлетворяет интерфейсу State
type ConcreteStateB struct{}


func (s *ConcreteStateB) Handle() {
	fmt.Println("Handling state B")
}

// Context содержит интерфейс
type Context struct {
	state State
}

// SetState изменяет состояние
func (c *Context) SetState(state State) {
	c.state = state
}

// HandleState использует функции интерфейса
func (c *Context) HandleState() {
	c.state.Handle()
}

func StateFunc() {
	// Create context
	context := &Context{}

	// Create state A
	stateA := &ConcreteStateA{}

	// Set state A
	context.SetState(stateA)

	// Handle state A
	context.HandleState()

	// Create state B
	stateB := &ConcreteStateB{}

	// Set state B
	context.SetState(stateB)

	// Handle state B
	context.HandleState()
}