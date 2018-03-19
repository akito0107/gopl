package ex03

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func randomPalindrome(rng *rand.Rand) string {
	ss := []rune{' ', ',', '?', '!', '.'}
	n := rng.Intn(12)
	src := make([]rune, n*2)
	for i := 0; i < n; i += 2 {
		r := rune(rng.Intn(0x1000))
		src[i] = r
		s := rng.Intn(len(ss) - 1)
		src[i+1] = ss[s]
	}
	d := make([]rune, n*12)
	copy(src, d)

	for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
		d[i], d[j] = d[j], d[i]
	}
	runes := append(src, d...)
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
