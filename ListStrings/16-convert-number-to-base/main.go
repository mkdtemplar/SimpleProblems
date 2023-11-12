package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func convertListToNumber(numSlice []string) int {
	var numString string

	for _, i := range numSlice {
		numString += i
	}

	convertedNumber, err := strconv.Atoi(numString)
	if err != nil {
		log.Printf("Can not convert %s to number\n", numString)
		return 0
	}

	return convertedNumber
}

func convertFromHighToLower(num int, newBase int) string {
	var resultSlice []string

	for num != 0 {
		rem := num % newBase
		remConvert := strconv.Itoa(rem)
		num = num / newBase
		resultSlice = append(resultSlice, remConvert)
	}

	resultSlice = reverseResult(resultSlice)

	return strings.Join(resultSlice, "")
}

func convertFromLowToHighBase(num int, newBase int) int {

	var result int

	t := strconv.Itoa(num)
	exponent := len([]rune(t)) - 1
	for _, i := range t {
		convertedNumber, err := strconv.Atoi(string(i))
		if err != nil {
			log.Printf("Can not convert %s to number\n", string(i))
			return 0
		}
		convertedNumber *= int(math.Pow(float64(newBase), float64(exponent)))
		result += convertedNumber
		exponent--
	}

	return result
}

func reverseResult(input []string) []string {

	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]

	}
	return input
}

func main() {

	var numberToConvert string
	var base int
	var newBase int

	fmt.Print("Enter the number to convert: ")
	_, err := fmt.Scan(&numberToConvert)
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Print("Enter the base of the number: ")
	_, err = fmt.Scan(&base)
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Print("Enter new base: ")
	_, err = fmt.Scan(&newBase)
	if err != nil {
		log.Println(err.Error())
		return
	}

	var inputNumberToInt int

	if newBase != 10 || base != 10 {
		inputNumberToInt, err = strconv.Atoi(numberToConvert)
		if err != nil {
			log.Println(err.Error())
			return
		}
		inputNumberToInt = convertFromLowToHighBase(inputNumberToInt, base)
		fmt.Println(convertFromHighToLower(convertListToNumber(strings.Split(strconv.Itoa(inputNumberToInt), "")), newBase))
	} else {
		if base < newBase {
			fmt.Println(convertFromLowToHighBase(convertListToNumber(strings.Split(strconv.Itoa(inputNumberToInt), "")), base))
		} else {
			fmt.Println(convertFromHighToLower(convertListToNumber(strings.Split(strconv.Itoa(inputNumberToInt), "")), newBase))
		}
	}

}
