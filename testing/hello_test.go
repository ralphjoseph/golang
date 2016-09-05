// testing project main.go
package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "hello world!"
	actual := hello("world")
	if actual != expected {
		t.Errorf("test failed, expected %s got %s", expected, actual)
	}
}
