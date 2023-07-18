package pattern

/*
Паттерн "Цепочка вызовов" (Chain of Responsibility pattern) является поведенческим паттерном проектирования, 
который позволяет передавать запросы последовательно по цепочке объектов-обработчиков, пока один из них не обработает запрос. 
Этот паттерн позволяет избежать явной привязки отправителя запроса к получателю и дает возможность обработать запрос нескольким 
объектам.

Применимость паттерна "Цепочка вызовов" возникает, когда в системе имеется набор объектов-обработчиков, и каждый из них может 
либо обработать запрос, либо передать его следующему объекту в цепочке. Паттерн "Цепочка вызовов" позволяет динамически 
конфигурировать объекты в цепочке и определить их порядок обработки.

Преимущества использования паттерна "Цепочка вызовов" включают:

1.	Гибкость и расширяемость. 
2.	Избежание привязки отправителя запроса к получателю. 
3.	Возможность обработки запросов несколькими объектами. 

Некоторые возможные недостатки паттерна "Цепочка вызовов" включают:

1.	Возможность неполной обработки запроса.
2.	Увеличение сложности отладки. 

Примеры использования паттерна "Цепочка вызовов" в реальных приложениях включают:

Обработка запросов веб-сервера. Паттерн "Цепочка вызовов" может быть использован для обработки запросов веб-сервера, 
где каждый объект-обработчик проверяет и обрабатывает определенные аспекты запроса, например, проверка аутентификации, 
проверка авторизации или обработка маршрутизации.
*/

import "fmt"


type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request string)
}

// ConcreteHandlerA удовлетворяет интерфейсу Handler
type ConcreteHandlerA struct {
	next Handler
}

// SetNext устанавливает след handler в цепи
func (h *ConcreteHandlerA) SetNext(handler Handler) {
	h.next = handler
}


func (h *ConcreteHandlerA) HandleRequest(request string) {
	if request == "A" {
		fmt.Println("Handled by handler A")
	} else if h.next != nil {
		h.next.HandleRequest(request)
	} else {
		fmt.Println("Unable to handle the request")
	}
}

// ConcreteHandlerB удовлетворяет интерфейсу Handler
type ConcreteHandlerB struct {
	next Handler
}

// SetNext устанавливает след handler в цепи
func (h *ConcreteHandlerB) SetNext(handler Handler) {
	h.next = handler
}


func (h *ConcreteHandlerB) HandleRequest(request string) {
	if request == "B" {
		fmt.Println("Handled by handler B")
	} else if h.next != nil {
		h.next.HandleRequest(request)
	} else {
		fmt.Println("Unable to handle the request")
	}
}

func ChainofRespFunc() {
	// Create handlers
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	// Set the next handler
	handlerA.SetNext(handlerB)

	// Handle requests
	handlerA.HandleRequest("A")
	handlerA.HandleRequest("B")
	handlerA.HandleRequest("C")
}