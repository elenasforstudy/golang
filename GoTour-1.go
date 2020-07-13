package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

// Упражнение: циклы и функции

func Sqrt(num float64) (int, float64) {
	var oldNum float64
	const diff = 0.00000000001
	iter := 0

	for i := 0; math.Abs(oldNum-num) > diff; i++ {
		oldNum = num
		num = num - float64(num*num-num)/float64(2*num)
		iter++
	}
	return iter, num

}

// Упражнение: срезы

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for x := range pic {
		pic[x] = make([]uint8, dx)

		for y := range pic[x] {
			pic[x][y] = uint8(x ^ y)
		}
	}
	return pic
}

//Упражнение: карты

func WordCount(s string) map[string]int {
	setStr := strings.Fields(s)
	mapStr := make(map[string]int)
	for str := range setStr {
		mapStr[setStr[str]]++
	}
	return mapStr
}

//Упражнение: замыкание Фибоначчи

func fibonacci() func() int {
	prev := -1
	cur := 1
	return func() int {
		now := prev + cur
		prev = cur
		cur = now
		return now
	}
}
func main() {
	//упражнение на циклы
	var x, y = Sqrt(5)
	fmt.Printf("Вычисление остановлено на %v итерации со значением %v", x, y)
	// упражнение на срезы
	pic.Show(Pic)
	// упражнение на map
	wc.Test(WordCount)
	f := fibonacci()
	//Упражнение: замыкание Фибоначчи
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
