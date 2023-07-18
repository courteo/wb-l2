package main

import (
	"reflect"
	"testing"
)

func TestAlotAnagram(t *testing.T) {
	slice := []string{"пятак","пятка","тяпка","листок","слиток","столик", "априве"}

	expected := &map[string][]string{
		"листок":{"листок", "слиток", "столик"},
		"пятак":{"пятак", "пятка", "тяпка"},
	}

	result := GetAnagram(&slice)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(`Ошибка в AlotAnagram. 
		Ожидаемый результат: 	%v
		Полученный результат:	%v`, 
		expected, result)
	}
}

func TestWithOutAnagram(t *testing.T) {
	slice := []string{"пятак", "априве"}

	expected := &map[string][]string{}

	result := GetAnagram(&slice)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(`Ошибка в AlotAnagram. 
		Ожидаемый результат: 	%v
		Полученный результат:	%v`, 
		expected, result)
	}
}