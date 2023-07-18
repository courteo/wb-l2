package pattern

/*
Паттерн "Строитель" (Builder pattern) является порождающим паттерном проектирования, который позволяет создавать сложные объекты пошагово. 
Он отделяет процесс создания объекта от его представления, позволяя использовать один и тот же процесс конструирования для создания различных представлений объекта.

Применимость паттерна "Строитель" возникает в случаях, когда у объекта может быть сложная структура или множество опциональных параметров для его создания. 
Он позволяет создавать объекты пошагово, концентрируясь на каждой составляющей и отделяя их от основного объекта.

Преимущества использования паттерна "Строитель" включают:

1. Упрощение процесса создания сложных объектов.
2. Улучшение читаемости и поддерживаемости кода. 
3. Возможность создания различных представлений объекта. 

Некоторые возможные недостатки паттерна "Строитель" включают:

1. Увеличение сложности кода. 
2. Возможность нарушения инкапсуляции. 
Примеры использования паттерна "Строитель" в реальных приложениях включают:

Создание объектов с большим количеством параметров. Например, при создании конфигурации комплексного сетевого сервера, 
где каждый параметр (IP-адрес, порт, протокол, безопасность и т. д.) может быть настроен отдельно с использованием класса строителя.
*/

import "fmt"

// Какая-то структура с поялми
type Product struct {
	name string
	count int
	isExist bool
}

// Наш строитель
type Builder interface {
	BuildName()
	BuildCount()
	BuildIsExist()
	GetProduct() *Product
}

// Структура которая будет использоваться для строителя
type ConcreteBuilder struct {
	product *Product
}

// Создание
func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{
		product: &Product{},
	}
}

// Определение функций чтобы соответствовать интерфейсу
func (b *ConcreteBuilder) BuildName() {
	b.product.name = "Part 1"
}

func (b *ConcreteBuilder) BuildCount() {
	b.product.count = 42
}

func (b *ConcreteBuilder) BuildIsExist() {
	b.product.isExist = true
}

func (b *ConcreteBuilder) GetProduct() *Product {
	return b.product
}

// Структура которая содержит интерфейс
type Director struct {
	builder Builder
}


func (d *Director) Construct() {
	d.builder.BuildName()
	d.builder.BuildCount()
	d.builder.BuildIsExist()
}

func BuilderFunc() {
	builder := NewConcreteBuilder()
	director := Director{builder: builder}

	director.Construct()

	product := builder.GetProduct()

	fmt.Printf("Product: %+v\n", product)
}