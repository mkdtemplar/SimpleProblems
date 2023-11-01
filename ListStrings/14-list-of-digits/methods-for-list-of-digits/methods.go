package methodsforlistofdigits

type ListOfDigits struct {
	n int
}

type IListOfDigits interface {
	PrintListOfDigits() []int
}

func NewListOfDigits(n int) IListOfDigits {
	return &ListOfDigits{n: n}
}

// PrintListOfDigits implements IListOfDigits.
func (l *ListOfDigits) PrintListOfDigits() []int {
	var listOfDigits []int
	for l.n > 0 {
		listOfDigits = append(listOfDigits, l.n%10)
		l.n = l.n / 10
	}

	for i, j := 0, len(listOfDigits)-1; i < j; i, j = i+1, j-1 {
		listOfDigits[i], listOfDigits[j] = listOfDigits[j], listOfDigits[i]
	}

	return listOfDigits
}
