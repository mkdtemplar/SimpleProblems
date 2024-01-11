package test

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/mkdtemplar/SimpleProblems/Elementary/02-Greetings/pkg"
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

func TestGreetingMethod(t *testing.T) {
	greeting := pkg.NewGreetings("Ivan")
	got := greeting.Greeting1()

	want := "Hello Ivan\n"

	if strings.Compare(got, want) != 0 {
		t.Errorf("incorrect greeting got: %v, want: %v", got, want)
	}
	if got != want {
		t.Errorf("incorrect greeting got: %v, want: %v", got, want)
	}

}
