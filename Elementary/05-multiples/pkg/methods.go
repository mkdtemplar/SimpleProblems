package pkg

type Multiplies struct {
	multipliesSlice []int
}

type IMultiplies interface {
	PrintMultiplies(n int) []int
}

func NewMultiplies(multipliesSlice []int) IMultiplies {
	return &Multiplies{multipliesSlice: multipliesSlice}
}

func (m *Multiplies) PrintMultiplies(n int) []int {
	for i := 3; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			m.multipliesSlice = append(m.multipliesSlice, i)
		}
	}

	return m.multipliesSlice
}
