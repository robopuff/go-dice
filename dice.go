package dice

import "regexp"

// FindBestResults Basic results structure
type FindBestResults struct {
	Source         string
	Ratings        []FindBestResultsRating
	BestMatch      FindBestResultsRating
	BestMatchIndex int
}

// FindBestResultsRating A rating used in FindBestResults
type FindBestResultsRating struct {
	Target string
	Rating float32
}

// ComparePair Compares pair of strings and produces similarity score (float32)
func ComparePair(input1, input2 string) float32 {
	rx := regexp.MustCompile(`\s+`)
	input1 = rx.ReplaceAllString(input1, "")
	input2 = rx.ReplaceAllString(input2, "")

	// Both are empty
	if len(input1) == 0 && len(input2) == 0 {
		return 1
	}

	// At least one is empty
	if len(input1) == 0 || len(input2) == 0 {
		return 0
	}

	// Strings are the same
	if input1 == input2 {
		return 1
	}

	// Both are 1-letter or
	// at least one of them is
	if len(input1) < 2 || len(input2) < 2 {
		return 0
	}

	firstBigrams := make(map[string]int)
	for i := 0; i < len(input1)-1; i++ {
		bigram := input1[i : i+2]
		count := 1
		if hit, ok := firstBigrams[bigram]; ok {
			count = hit + 1
		}

		firstBigrams[bigram] = count

	}

	intersectionSize := 0
	for i := 0; i < len(input2)-1; i++ {
		bigram := input2[i : i+2]
		count := 0
		if hit, ok := firstBigrams[bigram]; ok {
			count = hit
		}

		if count > 0 {
			firstBigrams[bigram]--
			intersectionSize++
		}
	}

	return (2.0 * float32(intersectionSize)) / (float32(len(input1)) + float32(len(input2)) - 2.0)
}

// FindBest Compares haystack to provided string (a needle) and returns FindBestResults
func FindBest(source string, haystack []string) FindBestResults {
	results := FindBestResults{
		Source:  source,
		Ratings: make([]FindBestResultsRating, len(haystack)),
	}

	highRating := float32(0)
	highIndex := 0
	for index, item := range haystack {
		itemRating := FindBestResultsRating{
			Target: item,
			Rating: ComparePair(source, item),
		}

		if itemRating.Rating > highRating {
			highRating = itemRating.Rating
			highIndex = index
		}

		results.Ratings[index] = itemRating
	}

	results.BestMatch = results.Ratings[highIndex]
	results.BestMatchIndex = highIndex

	return results
}
