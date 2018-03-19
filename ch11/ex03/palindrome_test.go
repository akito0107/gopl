package ex03

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 2
	runes := make([]rune, n)
	for i := 0; i < n; {
		r := rune(rng.Intn(0x1000))
		if unicode.IsLetter(r) {
			runes[i] = r
			i++
		}
	}
	i := len(runes)
	for j, r := range runes {
		if r != runes[i-j-1] {
			return string(runes)
		}
	}
	return randomNonPalindrome(rng)
}

func TestIsPalindrome(t *testing.T) {
	t.Run("palindrome = false", func(t *testing.T) {
		seed := time.Now().UTC().UnixNano()
		t.Logf("Random seed: %d", seed)
		rng := rand.New(rand.NewSource(seed))
		for i := 0; i < 1000; i++ {
			p := randomPalindrome(rng)
			if !IsPalindrome(p) {
				t.Errorf("isPalindrome(%q) = false", p)
			}
		}
	})
	t.Run("palindrome = true", func(t *testing.T) {
		seed := time.Now().UTC().UnixNano()
		t.Logf("Random seed: %d", seed)
		rng := rand.New(rand.NewSource(seed))
		for i := 0; i < 1000; i++ {
			p := randomNonPalindrome(rng)
			if IsPalindrome(p) {
				t.Errorf("isPalindrome(%q) = true", p)
			}
		}
	})
}
