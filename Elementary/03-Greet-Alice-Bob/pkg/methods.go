package pkg

import (
	"fmt"
)

type Names struct {
	sliceOfNames []string
}

type INames interface {
	PrintNames() string
}

func NewNames(sliceOfNames []string) INames {
	return &Names{sliceOfNames: sliceOfNames}
}

func (n *Names) PrintNames() string {
	return fmt.Sprintf("Hello %s and %s", n.sliceOfNames[0], n.sliceOfNames[1])
}
