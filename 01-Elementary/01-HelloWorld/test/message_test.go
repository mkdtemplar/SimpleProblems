package test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/mkdtemplar/SimpleProblems/Elementary/01-HelloWorld/pkg"
)

func TestPrintGreetMessage(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	greet := pkg.NewGreetMessage("Hello World")
	greet.PrintGreetMessage()
	err := w.Close()
	if err != nil {
		return
	}

	out, _ := io.ReadAll(r)
	os.Stdout = old
	got := string(out)

	want := "Hello World\n"
	if strings.Compare(got, want) != 0 {
		t.Errorf("incorrect greeting got: %v", string(out))
	}
	if got != want {
		t.Errorf("incorrect greeting got: %v", string(out))
	}
}

func TestGreetMessage1(t *testing.T) {
	greet := pkg.NewGreetMessage("Hello World")
	got := greet.PrintGreetMessage1()
	want := "Hello World"

	if strings.Compare(got, want) != 0 {
		t.Errorf("incorrect greeting got: %v, want: %v", got, want)
	}
	if got != want {
		t.Errorf("incorrect greeting got: %v, want: %v", got, want)
	}
}
