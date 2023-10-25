package test

import (
	"github.com/mkdtemplar/SimpleProblems/Elementary/02-Greetings/pkg"
	"io"
	"os"
	"strings"
	"testing"
)

func TestGreeting(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	greeting := pkg.NewGreetings("Ivan")
	greeting.Greeting()

	err := w.Close()
	if err != nil {
		return
	}

	out, _ := io.ReadAll(r)
	os.Stdout = old
	got := string(out)

	want := "Hello Ivan\n"
	if strings.Compare(got, want) != 0 {
		t.Errorf("incorrect greeting got: %v", string(out))
	}
	if got != want {
		t.Errorf("incorrect greeting got: %v", string(out))
	}
}
