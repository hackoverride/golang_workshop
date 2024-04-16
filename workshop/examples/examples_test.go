package examples

import (
	"testing"
)

func TestAddingOnExample(t *testing.T) {
	want := 3
	got, err := ExampleForTest(1, 2)
	if err != nil {
		t.Fatalf("ExampleForTest(1, 2) returned an unexpected error: %v", err)
	}
	if want != got {
		t.Fatalf("ExampleForTest(1, 2) = %d, want %d", got, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	got, err := ExampleForTest(1, -1)
	if err == nil {
		t.Fatalf("ExampleForTest(1, -1) expected an error but got nil")
	}
	if got != 0 {
		t.Fatalf("ExampleForTest(1, -1) = %d, want 0", got)
	}
}
