package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)


func compare(text, word string, v, i bool) bool {
	response := false
	if !v {
		response = true
	}

	first, second := text, word
	if i {
		first, second = strings.ToLower(text), strings.ToLower(word)
	}

	if first == second {
		return response
	}
	return !response
}


func main() {
	

	var  count, i, v, f, n bool
	var a, b, context int
	flag.IntVar(&a, "a", 0, "after")
	flag.IntVar(&b, "b", 0, "before")
	flag.IntVar(&context, "C", 0, "context")
	flag.BoolVar(&count, "c", false, "count")
	flag.BoolVar(&i, "i", false, "ignore-case")
	flag.BoolVar(&v, "v", false, "invert")
	flag.BoolVar(&f, "F", false, "fixed")
	flag.BoolVar(&n, "n", false, "line num")
	flag.Parse()
	// fmt.Println(a, b, context, count, i, v, f, n)
	args := flag.Args()
	var word string
	for i:=0; i < len(args) -1; i++ {
		word += args[i]
		if i != len(args)- 2 {
			word += " "
		}
	}
	// fmt.Println("word", word)
	input, errOpen := os.Open(args[len(args) - 1])					// имя файла 
	if errOpen != nil {
		log.Fatal(errOpen)
	}
	defer input.Close()

	countOfWord := 0
	firstFindedIndex := -1
	sliceOfInput := make([][]string, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if f {
			s := scanner.Text()
			if compare(s, word, v, i)  {
				if firstFindedIndex == -1 {
					firstFindedIndex = len(sliceOfInput) + 1
				}
				// fmt.Println(word, firstFindedIndex)
				countOfWord++
			}
			sliceOfInput = append(sliceOfInput, []string{s})
			continue
		}
		flag := false
		s := scanner.Text()
		slice := make([]string, 0)
		temp:=""
		for _, val := range s {
			if val != ' '{
				temp += string(val)

				// if compare(temp, word, v, i) {
				// 	if firstFindedIndex == -1 {
				// 		firstFindedIndex = len(sliceOfInput) +1
				// 	}
				// 	fmt.Println(temp ,word, firstFindedIndex)
				// 	flag = true
				// 	temp = ""
				// }
				

			} else if temp != "" {
				slice = append(slice, temp)
				// fmt.Println("temp ", temp)
				if compare(temp, word, v, i) {
					if firstFindedIndex == -1 {
						firstFindedIndex = len(sliceOfInput) +1
					}
					// fmt.Println(temp ,word, firstFindedIndex)
					flag = true
				}
				temp = ""
			}
		}
		if temp != "" {
			// fmt.Println("temp ", temp)
			if temp == word {
				if firstFindedIndex == -1 {
					firstFindedIndex = len(sliceOfInput) +1
				}
				flag = true
				// fmt.Println(temp ,word, firstFindedIndex)
			}
			slice = append(slice, temp)
		}
		if flag {
			countOfWord++
		}
		sliceOfInput = append(sliceOfInput, slice)
	}

	if firstFindedIndex == -1 {
		fmt.Println("Not found")
		return
	}

	if n {
		fmt.Println("Line num=",firstFindedIndex)
	}

	if count {
		fmt.Println("Count ", countOfWord)
	}

	var after, before int
	
	if context > 0 || a > 0 && b > 0{
		if context > 0 {
			after, before = context, context
		} else {
			after, before = a, b
		}	
	} else if a > 0 {
		after = a
	} else if b > 0 {
		before = b
	}
	// fmt.Println(firstFindedIndex, before, after)
	for i:=firstFindedIndex -1 - before; i < firstFindedIndex + after; i++ {
		// fmt.Println("i ", i, len(sliceOfInput))
		if i >= 0  && i < len(sliceOfInput) {
			fmt.Printf("Line %d - %s\n", i, sliceOfInput[i])
		}
	}
}