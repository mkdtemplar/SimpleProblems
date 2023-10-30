package pkg

import "fmt"

type Slice struct {
	newSlice []int
}

func (s *Slice) Run() {
	var element int

	fmt.Print("Enter number to check if it is in the slice: ")
	_, err := fmt.Scan(&element)
	if err != nil {
		return
	}

	if s.FindElement(element) {
		fmt.Printf("Entered number %d, is in the slice\n", element)
	} else {
		fmt.Printf("Entered number %d, is not present in the slice\n", element)
	}
}

func (s *Slice) FindElement(input int) bool {

	for _, element := range s.newSlice {
		if input == element {
			return true

		}
	}
	return false
}

func NewSlice(newSlice []int) ISlice {
	return &Slice{newSlice: newSlice}
}

type ISlice interface {
	FindElement(input int) bool
	Run()
}
