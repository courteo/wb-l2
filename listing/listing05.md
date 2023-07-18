Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error

Интерфейс хранит в себе тип интерфейса и тип самого значение.

Значение любого интерфейса  является nil в случае когда И значение И тип являются nil.

тк возвращается nil типа *customError, результат мы сравниваем с nil типа nil, откуда и следует их неравенство.

```