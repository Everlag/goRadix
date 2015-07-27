package radix

func Ints(a []int) {

	// Scratch space to work in.
	scratch:= make([]int, len(a))

	// Bucket on 0 or 1
	bucketLen:= 2
	buckets:= make([]int, bucketLen)

	// Assume 64 bit integers. Can do early exit later.
	var pos int = 0
	var i uint = 0
	var n int
	var pow2 int
	for ; i < 64; i++ {
	
		pow2 = 1 << i

		// Populate buckets
		for _, n:= range a{

			// Find if the bit is set then shift
			// right to avoid a conditional
			pos = (n & pow2) >> i
			buckets[pos]+= 1

		}

		// Sum the buckets with a running rightward count
		buckets[1] = buckets[0] + buckets[1]

		// Read off positions from the buckets backwards
		for p:= (len(a) - 1); p >= 0; p-- {
			
			n = a[p]
			pos = (n & pow2) >> i
			buckets[pos]--
			scratch[buckets[pos]] = n

		}

		copy(a, scratch)

		buckets[0] = 0
		buckets[1] = 0

	}

	return
}