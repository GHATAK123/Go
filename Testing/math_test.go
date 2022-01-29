package unitTesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd1(t *testing.T) {

	got := Add(4, 6)
	want := 10

	if got != want {
		t.Error("got {}, wanted {}", got, want)
	}
}

func TestAdd2(t *testing.T) {

	got := Add(4, 6)
	want := 11
	assert.NotEqual(t, got, want)

}
