package unitTesting

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd1(t *testing.T) {

	got := Add(4, 6)
	want := 10
	fmt.Println(got, want)
	if got != want {
		t.Error("got {}, wanted {}", got, want)
	}
}

func TestAdd2(t *testing.T) {

	got := Add(4, 6)
	want := 11
	fmt.Println(got, want)
	assert.NotEqual(t, got, want)

}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"
	fmt.Println(repeated, expected)
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
