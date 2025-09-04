package auth

import (
	"fmt"
	"testing"
)

func TestHashAndVerify(t *testing.T) {
	hash, err := HashPassword("a1sdf234", 10)

	fmt.Printf("hashed pw: %s\n", hash)

	if err != nil {
		t.Fatalf("hash error: %v", err)
	}

	if err := VerifyPassword("a1sdf234", hash); err != nil {
		t.Fatalf("expected err == nil")
	}

	if err := VerifyPassword("wrong", hash); err == nil {
		t.Fatalf("expected error != nil")
	}
}
