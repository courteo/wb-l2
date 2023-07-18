package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	slice := []string{"пятак","пятка","тяпка","листок","слиток","столик", "априве"}
	fmt.Println(*GetAnagram(&slice))
}

func GetAnagram(slice *[]string) *map[string][]string {
	res := make(map[string][]string)
	mapSortedWord := make(map[string]string)			// мапа, где ключ - отсортированная строка, значение - первое слово
	for _, val := range *slice {
		val = strings.ToLower(val)
		str := sortString(val)							// сортирует символы в строке 
		if v, ok := mapSortedWord[str]; !ok {			// проверяет есть ли такая анаграма, если нет создаем ключ в мапе, и добавляем в результирующей мапе
			mapSortedWord[str] = val
			res[val] = make([]string, 1)
			res[val][0] = val
		} else {
			res[v] = insert(res[v], val)				// если есть, то вставляем в результат бинарным поиском
		}
	}
	for key, val := range res {							// удаляем множества из одного элемента
		if len(val) == 1 {
			delete(res, key)
		}
	}
	return &res
}

func sortString(val string) string {					// сортирует символы в строке
	runes := []rune(val)
	sort.SliceStable(runes, func(i,j int) bool{
		return runes[i] < runes[j]
	})
	return string(runes)
}

func insert(slice []string, val string) []string {		// вставка бинарным поиском в отсортированный массив
	if slice[0] >= val {								// если надо в начало вставить
		return append([]string{val}, slice...)
	}

	if slice[len(slice)-1] <= val {						// если надо в конец вставить
		return append(slice, val)
	}

	start, end := 0, len(slice)
	for end > start {									// ищем место в вставки
		mid := (start + end) / 2
		if slice[mid] > val {
			end = mid
		} else {
			start = mid + 1
		}
	}
	slice = append(slice, "")
	for i := len(slice) - 1; i > end; i-- {				// освобождаем место вставки 
		slice[i] = slice[i-1]
	}
	slice[end] = val
	return slice
}