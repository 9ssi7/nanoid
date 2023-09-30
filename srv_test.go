package nanoid

import (
	"testing"

	"github.com/google/uuid"
)

func TestNanoID(t *testing.T) {
	shortURL1, _ := New()
	shortURL2, _ := New()

	if shortURL1 == shortURL2 {
		t.Fatalf("Expected unique short URLs, but got %s and %s", shortURL1, shortURL2)
	}
}

func BenchmarkNanoID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = New()
	}
}

func BenchmarkGoogleUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = uuid.New().String()
	}
}
