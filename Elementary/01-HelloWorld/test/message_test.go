package test

import (
	"github.com/mkdtemplar/SimpleProblems/Elementary/01-HelloWorld/pkg"
	"io"
	"os"
	"strings"
	"testing"
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
