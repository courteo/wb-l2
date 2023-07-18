package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)


var f, d string 
var separated bool

func regStringVar(p *string, name string, value string, usage string) {
    if flag.Lookup(name) == nil {
        flag.StringVar(p, name, value, usage)
    }
}

func regBoolVar(p *bool, name string, value bool, usage string) {
    if flag.Lookup(name) == nil {
        flag.BoolVar(p, name, value, usage)
    }
}

func getStringFlag(name string) string {
    return flag.Lookup(name).Value.(flag.Getter).Get().(string)
}

func getBoolFlag(name string) bool {
    return flag.Lookup(name).Value.(flag.Getter).Get().(bool)
}


func init() {
    regStringVar(&f, "f", "", "fields")
    regStringVar(&d, "d", " ", "delimiter")
    regBoolVar(&separated, "s", false, "separated")
}

func initFlags() {
    f = getStringFlag("f")
    d = getStringFlag("d")
    separated = getBoolFlag("s")
}

func main() {

	flag.Parse()
	initFlags()	

	// log.Println(f,d, separated)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		s := scanner.Text()
		if s == "" {
			return
		}

		res := make([]string, 0)
		temp := ""

		for _, val := range s {
			if val != rune(d[0]) {							// if not delimiter then add to temp
				temp += string(val)
				continue
			}	
			if f == "" {									// if no fields then print else add to slice of strings
				fmt.Print(temp, " ")
			} else {
				res = append(res, temp)
			}
			
			temp = ""
		}

		if temp != "" {										// check last temp
			if temp == s && separated {
				continue
			}
			
			if f == "" {
				fmt.Print(temp, " ")
			} else {
				res = append(res,  temp)
			}
		}

		if f == "" {
			fmt.Println()
		}
		
		if f != "" {
			var numbers []string
			flag := strings.Contains(f, ",") 
			if flag {
				numbers= strings.Split(f, ",")
			} else {
				numbers= strings.Split(f, "-")
			}
			

			if len(numbers) == 1 {
				index, err := strconv.Atoi(numbers[0])
				if err != nil {
					panic("fields error " + err.Error())
				}
				fmt.Println(res[index - 1])
			} else if flag {
				for i, val := range numbers {
					index, err := strconv.Atoi(val)
					
					if err != nil {
						panic("fields error " + err.Error())
					}
					fmt.Print(res[index - 1])
					if i != len(numbers) - 1{
						fmt.Print(" ")
					} else {
						fmt.Print(" ")
					}
				}
			} else {

				first, err := strconv.Atoi(numbers[0])
				if err != nil {
					panic("fields error " + err.Error())
				}

				second , err := strconv.Atoi(numbers[1])
				if err != nil {
					panic("fields error " + err.Error())
				}


				for i:=first -1; i < second ;i++ {
					fmt.Print(res[i])
					
					if i != second - 1 {
						fmt.Print(" ")
					} else {
						fmt.Println()
					}
				}
			}
		}
	}
	f= ""
	d = " "
	separated = false
}