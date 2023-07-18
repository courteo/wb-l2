package pattern

/*
	Паттерн "Команда" (Command pattern) является поведенческим паттерном проектирования, который инкапсулирует запрос как объект, 
	позволяя передавать его в качестве аргумента, сохранять его в виде объекта и выполнять операции с ним. Он позволяет отделить 
	отправителя запроса от получателя, обеспечивая гибкость и расширяемость взаимодействия между ними.

Применимость паттерна "Команда" возникает в ситуациях, когда требуется параметризовать объекты или операции, выполняемые над ними. 
Паттерн "Команда" позволяет инкапсулировать операцию как объект, который можно передавать, сохранять и использовать для выполнения 
различных операций.

Преимущества использования паттерна "Команда" включают:

1.	Расширяемость и гибкость. 
2.	Отделение отправителя и получателя. 
3.	Поддержка отмены и повтора операций. 

Некоторые возможные недостатки паттерна "Команда" включают:

1.	Увеличение сложности кода.
2.	Увеличение накладных расходов на память. 

Примеры использования паттерна "Команда" в реальных приложениях включают:

Реализация истории операций в текстовом редакторе. Каждая команда представляет отдельную операцию, 
такую как вставка текста, удаление текста или форматирование. Команды могут быть сохранены в истории и отменены/повторены 
при необходимости.
*/


import "fmt"

type Command interface {
	Execute()
}

// Создаем структуру StartCommand удовлетворяющую интерфейсу
type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

func (c *StartCommand) Execute() {
	c.mb.Start()
}

// Создаем структуру RebootCommand удовлетворяющую интерфейсу
type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}


// Создаем структуру, в которой будут функцию для стуктур Rebbot Start
type MotherBoard struct{}

func (*MotherBoard) Start() {
	fmt.Print("system starting\n")
}

func (*MotherBoard) Reboot() {
	fmt.Print("system rebooting\n")
}

// В этой структуре содержится две какие-то команды, которые можно вызвать
type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}


// вызов этих комманд
func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}

func CommandFunc() {
	mb := &MotherBoard{}
	startCommand := NewStartCommand(mb)
	rebootCommand := NewRebootCommand(mb)

	box1 := NewBox(startCommand, rebootCommand)
	box1.PressButton1()
	box1.PressButton2()

	box2 := NewBox(rebootCommand, startCommand)
	box2.PressButton1()
	box2.PressButton2()
	// Output:
	// system starting
	// system rebooting
	// system rebooting
	// system starting
}