package main

import (
	"fmt"
	"strings"
)

func unPacking(str string) (string, error) {
	if str == "" {
		return "", fmt.Errorf("incorrect string\n")
	}

	res := ""
	flag := 0
	count := 0
	temp := ""
	for _,symbol := range str {
		if flag == 1 {									// предыдущий символ \
			temp += string(symbol)
			flag = 0 
			// fmt.Println("flag ", string(symbol), temp)
			continue
		}
		if symbol >= '0' && symbol <= '9' {
			count = count*10 + int(symbol - '0')
			// fmt.Println("count ", count)
			continue
		}

		if symbol == 92 {								// symbol == \
			flag = 1 
			if temp != "" {
				res += temp
				temp = ""
			}
			if count > 0 {							    // повторить строку
				// fmt.Println("repeat ",temp, count)
				res += strings.Repeat(temp, count)
				
				count = 0	
				temp  = ""
			}
			continue
		}
		if symbol < 'a' || symbol > 'z' {				// неправильная строка 
			// fmt.Println("return ", string(symbol))
			return "", fmt.Errorf("incorrect string\n")
		}
		if count > 0 {									// повторить строку 
			//  fmt.Println("repeat ",temp, count)
			res += strings.Repeat(temp, count)
			
			count = 0	
			temp  = ""
		}
		if temp != "" {									// просто добавить букву	
			res += temp
			temp = ""
		}
		temp += string(symbol)	
	}
	
	if count > 0 {									// если строка заканчивалась числом
		//  fmt.Println("last repeat ", temp, count, res)
		if temp == "" {
			return "",  fmt.Errorf("incorrect string\n")
		}
		res += strings.Repeat(temp, count)
		temp = ""
		count = 0
	}
	return res + temp, nil
}

func main() {
	s := `qwe\45`
	fmt.Println(unPacking(s))
}