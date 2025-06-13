package test

import (
	"github.com/pkg/errors"
	"testing"
)

func TestWelcome(t *testing.T) {
	var l []int
	for i := 0; i < 10; i++ {
		l = append(l, i)
	}
}

// func TestWelcome2(t *testing.T) {
// 	t.Errorf("Error: %v", errors.New("Жесть ошибка"))
// }
