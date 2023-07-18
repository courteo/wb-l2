package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

/*
	Сделано для двух типов: чисел и строк
	Для чисел в случае флага -n, иначе для строк
	Обьединено с помощью интерфейса Slices
*/



type Slices interface {
	sortColumn(slice Slices, k int, r bool) Slices
	MakeSort(scanner *bufio.Scanner, k int, r, u bool) error 
	removeDuplicate(slice Slices) Slices
}


type SliceInt	[][]int

func (s SliceInt) removeDuplicate(slice SliceInt) SliceInt {			// убирает дубликаты для флага -u
	sort.Slice(slice[:], func(i, j int) bool {							// сортируем матрицу
		for x := range slice[i] {
			if slice[i][x] == slice[j][x] {
				continue
			}
			return slice[i][x] < slice[j][x]
		}
		return false
	})
	
	res := [][]int{}
	res = append(res, slice[0])
	for i:=1; i < len(slice); i++ {	
		if !reflect.DeepEqual(slice[i], slice[i -1]) {					// сравниваем строки матрицы и если не совпали значит добавляем в результат
			res = append(res, slice[i])
		}
	}
    return res
}

func (s SliceInt) sortColumn(slice SliceInt, k int, r bool) SliceInt {						// сортирует по колонне
	sort.SliceStable(slice, func (i, j int) bool {
		if k >= len(slice[i]) && k >= len(slice[j]){								
			return slice[i][0] < slice[j][0]  && !r || slice[i][0] > slice[j][0]  && r		// если нужной колонны нет в обоих строках, то смотрим по первой колонне
		} else if k >= len(slice[i]) {
			return slice[i][0]  < slice[j][k] && !r || slice[i][0]  > slice[j][k] && r		// если нет нужной в одной из строк, то сравниваем с нужную с первой колонной
		}else if k >= len(slice[j]) {
			return slice[i][k] < slice[j][0] && !r || slice[i][k] > slice[j][0] && r
		}
		return slice[i][k] < slice[j][k] && !r ||  slice[i][k] > slice[j][k] && r			// если есть нужная колонна в обоих строках
	})
	return slice
}

func (s SliceInt) MakeSort(scanner *bufio.Scanner, k int, r, u bool) error{					// основной код + запись в файл
	sliceInt := make(SliceInt, 0)
    for scanner.Scan() {																	// из массива строк делаем матрицу чисел, если есть не числа то выкидываем ошибку
		s := scanner.Text()
		slice := make([]int, 0)
		temp:=""
		for _, val := range s {
			if val != ' '{
				temp += string(val)
			} else if temp != "" {
				i, err := strconv.Atoi(temp)
				if err != nil {
					return err
				}

				slice = append(slice, i)
				temp = ""
			}
		}
		if temp != "" {
			i, err := strconv.Atoi(temp)
			if err != nil {
				return err
			}

			slice = append(slice, i)
		}
		sliceInt = append(sliceInt, slice)
    }

	if u {																					// убираем дубликаты если нужно
		sliceInt = s.removeDuplicate(sliceInt)
	}

	// fmt.Println(sliceInt)
	sliceInt = sliceInt.sortColumn(sliceInt, k, r)														// сортируем по колоннам
	// fmt.Println(sliceInt)
	output, errCreate := os.Create("outputInt.txt")
	if errCreate != nil {
		log.Fatal(errCreate)
	}
	defer output.Close()

	for _, word := range sliceInt {															// записываем в файлы
		str := make([]string, 0)
		for _, val := range word {
			str = append(str, strconv.Itoa(val))
		}
        _, err := output.WriteString(strings.Join(str, " ") + "\n")

        if err != nil {
            log.Fatal(err)
        }
    }
	return nil
}



type SliceString [][]string															// для строк аналогично

func (s SliceString) removeDuplicate(slice SliceString) SliceString {
	sort.Slice(slice[:], func(i, j int) bool {
		for x := range slice[i] {
			if slice[i][x] == slice[j][x] {
				continue
			}
			return slice[i][x] > slice[j][x]
		}
		return false
	})
	
	res := [][]string{}
	res = append(res, slice[0])
	for i:=1; i < len(slice); i++ {
		if !reflect.DeepEqual(slice[i], slice[i -1]) {
			res = append(res, slice[i])
		}
	}
    return res
}

func (s SliceString) MakeSort(scanner *bufio.Scanner, k int, r, u bool) (error) {
	sliceStrings := make(SliceString, 0)
    for scanner.Scan() {
		s := scanner.Text()
		slice := make([]string, 0)
		temp:=""
		for _, val := range s {
			if val != ' '{
				temp += string(val)
			} else if temp != "" {
				slice = append(slice, temp)
				temp = ""
			}
		}
		if temp != "" {
			slice = append(slice, temp)
		}
		sliceStrings = append(sliceStrings, slice)
    }

	if u {
		sliceStrings = s.removeDuplicate(sliceStrings)
	}


	sliceStrings.sortColumn(sliceStrings, k, r)

	output, errCreate := os.Create("outputString.txt")
	if errCreate != nil {
		log.Fatal(errCreate)
	}
	defer output.Close()

	for _, word := range sliceStrings {

        _, err := output.WriteString(strings.Join(word, " ") + "\n")

        if err != nil {
            log.Fatal(err)
        }
    }


	return nil
}

func (s SliceString) sortColumn(slice SliceString, k int, r bool) SliceString {
	sort.SliceStable(slice, func (i, j int) bool {
		if k >= len(slice[i]) && k >= len(slice[j]){								
			return slice[i][0] < slice[j][0]  && !r || slice[i][0] > slice[j][0]  && r		// если нужной колонны нет в обоих строках, то смотрим по первой колонне
		} else if k >= len(slice[i]) {
			return slice[i][0]  < slice[j][k] && !r || slice[i][0]  > slice[j][k] && r		// если нет нужной в одной из строк, то сравниваем с нужную с первой колонной
		}else if k >= len(slice[j]) {
			return slice[i][k] < slice[j][0] && !r || slice[i][k] > slice[j][0] && r
		}
		return slice[i][k] < slice[j][k] && !r ||  slice[i][k] > slice[j][k] && r			// если есть нужная колонна в обоих строках
	})
	return slice
}

func main() {
	var  n, r, u bool										// нужные нам флаги
	var k int
	flag.IntVar(&k, "k", 0, "by column")
	flag.BoolVar(&n, "n", false, "only numeric")
	flag.BoolVar(&r, "r", false, "reverse")
	flag.BoolVar(&u, "u", false, "without duplicate")
	flag.Parse()


	input, errOpen := os.Open(flag.Arg(0))					// имя файла 
	if errOpen != nil {
		log.Fatal(errOpen)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	
	if n {													// если для чисел
		var slice SliceInt
		err := slice.MakeSort(scanner, k, r, u)
		if err != nil {
			log.Fatal("not int ", err.Error())
		}
	} else {												// если для строк
		var slice SliceString
		slice.MakeSort(scanner, k, r, u)
	}
}