package util

import (
	"testing"
	"fmt"
)

func TestGetEnv(t *testing.T) {
	t.Setenv("PORT", "3000")
	var PORT string = GetEnv("PORT")
	fmt.Print(PORT)
	if PORT == "" {
		t.Fatal("Port should not be empty")
	}
}
