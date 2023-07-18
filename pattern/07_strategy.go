package pattern

/*
Паттерн "Стратегия" (Strategy pattern) является поведенческим паттерном проектирования, который позволяет определить 
семейство алгоритмов, инкапсулировать каждый из них и делать их взаимозаменяемыми. Этот паттерн позволяет изменять алгоритмы 
независимо от клиентов, которые их используют.

Применимость паттерна "Стратегия" возникает в ситуациях, когда у вас есть несколько алгоритмов или вариаций операций, 
которые могут быть применены в зависимости от контекста. Паттерн "Стратегия" позволяет инкапсулировать каждую стратегию 
в отдельном классе и делает их взаимозаменяемыми, что обеспечивает гибкость и удобство в изменении алгоритмов.

Преимущества использования паттерна "Стратегия" включают:

1.	Гибкость и расширяемость. 
2.	Избегание условных операторов. 
3.	Упрощение тестирования. 
Некоторые возможные недостатки паттерна "Стратегия" включают:

1.	Увеличение числа классов. 
2.	Возможная сложность конфигурации. 

Примеры использования паттерна "Стратегия" в реальных приложениях включают:

Сортировка данных. Паттерн "Стратегия" может быть использован для реализации различных алгоритмов сортировки, 
таких как быстрая сортировка, сортировка слиянием или пузырьковая сортировка. Клиентский код может выбирать нужную 
стратегию сортировки в зависимости от требуемого результата или размера данных.
*/

import "fmt"


// структура которая использует нашу стратегию и данные в зависимости от контекста
type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}



func NewPayment(name, cardid string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardid,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

// данные с которыми оперируем
type PaymentContext struct {
	Name, CardID string
	Money        int
}


// интерфейс наших стратегий
type PaymentStrategy interface {
	Pay(*PaymentContext)
}


// структуры с конкертными реализациями Pay
type Cash struct{}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

type Bank struct{}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardID)

}

func StrategyFunc() {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()
	// Output:
	// Pay $123 to Ada by cash

	payment = NewPayment("Bob", "0002", 888, &Bank{})
	payment.Pay()
	// Output:
	// Pay $888 to Bob by bank account 0002
}