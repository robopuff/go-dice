package dice

import "regexp"

// ComparePair Compares pair of strings and produces similarity score (float32)
func ComparePair(input1, input2 string) float32 {
	rx := regexp.MustCompile(`\s+`)
	input1 = rx.ReplaceAllString(input1, "")
	input2 = rx.ReplaceAllString(input2, "")

	len1 := len(input1)
	len2 := len(input2)

	// Both are empty
	// Or strings are the same
	if len1 == 0 && len2 == 0 || input1 == input2 {
		return 1
	}

	// Both are 1-letter or
	// at least one of them is
	if len1 < 2 || len2 < 2 {
		return 0
	}

	bigrams := make(map[string]int)
	for i := 0; i < len1-1; i++ {
		bigram := input1[i : i+2]
		count := 1
		if hit, ok := bigrams[bigram]; ok {
			count = hit + 1
		}

		bigrams[bigram] = count
	}

	intersection := float32(0)
	for i := 0; i < len2-1; i++ {
		bigram := input2[i : i+2]
		count := 0
		if hit, ok := bigrams[bigram]; ok {
			count = hit
		}

		if count > 0 {
			bigrams[bigram]--
			intersection++
		}
	}

	return (2.0 * intersection) / (float32(len1+len2) - 2.0)
}

// FindBest Compares haystack to provided string (a needle)
// Returns an index of best matched string and an array of scores
func FindBest(source string, haystack []string) (int, []float32) {
	ratings := make([]float32, len(haystack))

	var bestIndex int
	highRating := float32(0)
	for index, item := range haystack {
		score := ComparePair(source, item)

		if score > highRating {
			highRating = score
			bestIndex = index
		}

		ratings[index] = score
	}

	return bestIndex, ratings
}
