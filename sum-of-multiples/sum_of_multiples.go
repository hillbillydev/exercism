package summultiples

func SumMultiples(n int, divisors ...int) int {
	var (
		result int
		duplicates   = make(map[int]bool)
	)

	for _, div := range divisors {
        if div == 0 {
            continue
        }

		for i := 0; div*i < n; i++ {
            product := div * i

			if duplicates[product] {
				continue
			}

			result += product
			duplicates[product] = true
		}
	}

	return result
}
