```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

Ответ:
```
часть из массива а вырезали и сделали срезом b
так же можно из массива а сделать срез путем a[:]

```