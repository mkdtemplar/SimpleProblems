package pkg_for_arithmetic

import (
	"strconv"
	"strings"
)

type Numbers struct {
	num1 string
	num2 string
}

type INumbers interface {
	Addition() string
	Subtraction() string
}

func NewNumbers(num1 string, num2 string) INumbers {
	return &Numbers{num1: num1, num2: num2}
}

func StringToDigitsSlice(num1 string) []int64 {
	var numSlice []int64
	for _, e := range num1 {
		char, _ := strconv.Atoi(string(e))

		numSlice = append(numSlice, int64(char))
	}

	return numSlice
}

func RemoveLeadingZeros(res []int64) []int64 {
	indicator := -1
	var output []int64
	for i, e := range res {
		if e != 0 {
			indicator = i
			break
		}
	}

	for i := 0; i < len(res)-indicator; i++ {
		output = append(output, res[indicator+i])
	}

	return output
}

func OptimizeSlices(num1 []int64, num2 []int64) ([]int64, []int64) {
	var dif int
	var num1New []int64
	var num2New []int64
	if len(num1) > len(num2) {
		dif = len(num1) - len(num2)
		num2New = make([]int64, len(num1))
		for i := len(num2) - 1; i >= 0; i-- {
			num2New[i+dif] = num2[i]
		}
		num2 = num2New
	} else if len(num2) > len(num1) {
		dif = len(num2) - len(num1)
		num1New = make([]int64, len(num2))
		for i := len(num1) - 1; i >= 0; i-- {
			num1New[i+dif] = num1[i]
		}
		num1 = num1New
	}

	return num1, num2
}

func AddingDigits(num1 []int64, num2 []int64) []int64 {
	resultSlice := make([]int64, len(num1)+1)
	var carry int
	for i := len(num1) - 1; i >= 0; i-- {
		res := num1[i] + num2[i] + int64(carry)
		if res >= 10 {
			temp := res
			res = int64(int(res % 10))
			carry = int(temp / 10)
			resultSlice[i+1] = res
			if carry > 0 && i-2 >= 0 {
				resultSlice[i-2] += int64(carry)
			}
		} else {
			resultSlice[i+1] = res
		}
	}

	return resultSlice
}

func NumSliceToStringSlice(resultSlice []int64) []string {

	var numberString []string

	for _, e := range resultSlice {
		numberString = append(numberString, strconv.Itoa(int(e)))
	}

	return numberString
}

func (n *Numbers) Addition() string {
	num1 := StringToDigitsSlice(n.num1)
	num2 := StringToDigitsSlice(n.num2)

	num1, num2 = OptimizeSlices(num1, num2)

	resultSlice := AddingDigits(num1, num2)

	resultSlice = RemoveLeadingZeros(resultSlice)

	resultString := NumSliceToStringSlice(resultSlice)

	numberString := strings.Join(resultString, "")

	return numberString
}

func (n *Numbers) Subtraction() string {
	//TODO implement me
	panic("implement me")
}
