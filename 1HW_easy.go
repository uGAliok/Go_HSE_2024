package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func HelloWorld() {
	fmt.Println("Hello World")
}

func add(x int, y int) int {
	return x + y
}

func isEven(x int) bool {
	return x%2 == 0
}

func MaxOfThree(x int, y int, z int) int {
	if x > y {
		if x > z {
			return x
		}
	} else {
		if y > z {
			return y
		}
	}
	return z
}

func factorial(x int) (res int) {
	res = 1
	if x == 0 {
		return
	}
	for i := 2; i <= x; i++ {
		res *= i
	}
	return
}

func isVowel(x string) bool {
	a := unicode.ToLower(rune(x[0]))
	// да, не очень оптимально, но хотелось побаловаться с свитчем
	switch a {
	case 'o':
		return true
	case 'e':
		return true
	case 'y':
		return true
	case 'u':
		return true
	case 'i':
		return true
	case 'a':
		return true
	default:
		return false
	}
}

func PrintPlain(x int) {
	var j, a int
	var flag bool
	for a = 2; a < x; a++ {
		flag = false
		for j = 2; j <= a/2; j++ {
			if a%j == 0 {
				flag = true
				continue
			}
		}
		if flag == false {
			fmt.Println(a)
		}
	}
}

func reverse(x string) string {
	runes := []rune(x)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func summ(s []int) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += s[i]
	}
	return res
}

type Rectangle struct {
	Height, Width int
}

func (r *Rectangle) Area() int {
	return r.Height * r.Width
}

type celsius float64
type fahrenheit float64

func (c celsius) Fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func countdown(n int) {
	for i := n; i > 0; i-- {
		fmt.Println(i)
	}
}

func stringLength(s string) int {
	length := 0
	for range s {
		length++
	}
	return length
}

func contains(arr []int, elem int) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}

func average(arr []int) float64 {
	if len(arr) == 0 {
		return 0
	}
	sum := summ(arr) // из предыдущих заданий
	return float64(sum) / float64(len(arr))
}

func multitable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", n, i, n*i)
	}
}

func isPalindrom(x string) bool {
	x = strings.ToLower(x)
	x = strings.ReplaceAll(x, " ", "")
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		if x[i] != x[j] {
			return false
		}
	}
	return true
}

func findMinMax(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}
	min := math.MaxInt64
	max := math.MinInt64
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func removeElem(arr *[]int, index int) {
	if index < 0 || index >= len(*arr) {
		return
	}
	copy((*arr)[index:], (*arr)[index+1:])
	*arr = (*arr)[:len(*arr)-1]
}

func linSearch(arr []int, elem int) int {
	for i, num := range arr {
		if num == elem {
			return i
		}
	}
	return -1
}

func main() {
	HelloWorld()

	var a, b = 1, 2
	fmt.Println(add(a, b))

	fmt.Println(isEven(a))
	fmt.Println(isEven(b))

	c := 3
	fmt.Println(MaxOfThree(a, b, c))

	fmt.Println(factorial(c))

	fmt.Println(isVowel("w"))
	fmt.Println(isVowel("o"))

	PrintPlain(9)

	word := "Hello World"
	fmt.Println(reverse(word))

	array := []int{1, 4, 3, 7, 5}
	fmt.Println(summ(array))

	r := Rectangle{Height: 2, Width: 2}
	fmt.Println(r.Area())

	var temp celsius = 30
	fmt.Println(temp.Fahrenheit())

	countdown(7)

	fmt.Println(stringLength(word))

	fmt.Println(contains(array, c))

	fmt.Println(average(array))

	multitable(7)

	palindrom := "A rosa upala na lapu Asora"
	fmt.Println(isPalindrom(palindrom))

	min, max := findMinMax(array)
	fmt.Printf("Минимум: %d \nМаксмум: %d \n", min, max)

	fmt.Println("До ", array)
	removeElem(&array, 1)
	fmt.Println("Удалили 2-й элемент(например) ", array)

	fmt.Println(linSearch(array, 7), "  ", linSearch(array, 10))
}
