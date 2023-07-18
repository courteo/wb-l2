Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
будет выведено какое-то количество чисел из нужных, а потом нули

потому что в select не проверяется закрыты ли каналы a,b и просто получается дефолтное значение 0, цикл по каналу с d в main будет бесконечен тк канал не закрыт 

я бы предложил эту реализацию функции

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
        flagA := false
        flagB := false
		for {
            if flagA && flagB {
                close(c)
                return
            }
			select {
			case v, ok := <-a:
            if !ok {
                flagA = true
            } else {
                c <- v
            }
				
			case v, ok:= <-b:
            if !ok {
                flagB = true
            } else {
                c <- v
            }
			}
		}
	}()
	return c
}
```