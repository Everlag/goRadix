package radix

import (
	"testing"

	"sort"

	"math/rand"
)

var seed int = 99
var random []int = nRands(1000000, seed)

func TestSort(t *testing.T) {

	// Stdlib
	std:= make([]int, len(random))
	copy(std, random)
	sort.Ints(std)

	// Custom radix
	radix:= make([]int, len(random))
	copy(radix, random)
	Ints(radix)

	// Make sure Ints is aligned with std
	for i, n:= range std{
		if radix[i] != n {
			t.Fatal("Unequal results")
		}
	}

}

func BenchmarkInts(b *testing.B) {
	radix:= make([]int, len(random))

	for n := 0; n < b.N; n++ {
		copy(radix, random)
		Ints(radix)
	}
}

func BenchmarkStd(b *testing.B) {
	std:= make([]int, len(random))

	for n := 0; n < b.N; n++ {
		copy(std, random)
		sort.Ints(std)
	}
}

// Returns an array of random integers of length n
// with source seeded with the provided seed
func nRands(n, seed int) [] int {

	r := rand.New(rand.NewSource(int64(seed)))

	random:= make([]int, n)

	for i := 0; i < n; i++ {
		random[i] = r.Int()
	}

	return random
}
