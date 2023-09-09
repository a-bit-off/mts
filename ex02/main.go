/*
Условие:
Дан массив строк. Необходимо вернуть “true”, если все строки в массиве по вертикали и горизонтали
образуют одинаковый набор символов.

Пример 1:
Массив:

	["abcd","bnrt","crmy","dtye"]

Результат:

	true

Объяснение:

	1-я строка и 1-я колонка составляют "abcd".
	2-я строка и 2-я колонка составляют "bnrt".
	3-я строка и 3-я колонка составляют "crmy".
	4-я строка и 4-я колонка составляют "dtye".

Ограничения:

	1 <= длинна массива строк <= 500
	1 <= длинна строки <= 500
	Все символы в строках - только строчные латинские буквы.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	data := generateString(500)
	fmt.Println(CheckArrayConsistency(&data))
}

// generateString ...
// Принимает размер матрицы
// Возвращает массив строк заполненный символами 'a'
func generateString(n int) []string {
	str := make([]string, n)
	for i := 0; i < n; i++ {
		str[i] = strings.Repeat("a", n)
	}
	return str
}

// CheckArrayConsistency ...
// Принимает указатель на массив строк
// Возвращает true, если символы относитльно главной диагонали равны
func CheckArrayConsistency(data *[]string) bool {
	n := len(*data)
	for i := 0; i < n; i++ {
		if len((*data)[i]) != n {
			return false
		}
		for j := i; j < n; j++ {
			if (*data)[i][j] != (*data)[j][i] {
				return false
			}
		}
	}
	return true
}
