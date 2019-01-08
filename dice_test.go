package dice

import (
	"testing"
)

func TestComparePairs(t *testing.T) {
	tables := []struct {
		input1 string
		input2 string
		result float32
	}{
		{"french", "quebec", 0},
		{"france", "france", 1},
		{"fRaNce", "france", 0.2},
		{"healed", "sealed", 0.8},
		{"web applications", "applications of the web", 0.7878787878787878},
		{"this will have a typo somewhere", "this will huve a typo somewhere", 0.92},
		{"Olive-green table for sale, in extremely good condition.", "For sale: table in very good  condition, olive green in colour.", 0.6060606060606061},
		{"Olive-green table for sale, in extremely good condition.", "For sale: green Subaru Impreza, 210,000 miles", 0.2558139534883721},
		{"Olive-green table for sale, in extremely good condition.", "Wanted: mountain bike with at least 21 gears.", 0.1411764705882353},
		{"this has one extra word", "this has one word", 0.7741935483870968},
		{"a", "a", 1},
		{"a", "b", 0},
		{"", "", 1},
		{"a", "", 0},
		{"", "a", 0},
		{"apple event", "apple    event", 1},
		{"iphone", "iphone x", 0.9090909090909091},
	}

	for _, table := range tables {
		score := ComparePair(table.input1, table.input2)
		if score != table.result {
			t.Errorf("Incorrect score when comparing `%s` and `%s` - expected %f, got %f", table.input1, table.input2, table.result, score)
		}
	}
}

func TestFindBest(t *testing.T) {
	result := FindBest("healed", []string{"mailed", "edward", "sealed", "theatre"})

	if result.Source != "healed" {
		t.Errorf("Result source is incorrect, expected `healed` got `%s` ", result.Source)
	}

	if result.BestMatchIndex != 2 {
		t.Errorf("Wrong best match index, expected `2` got `%d`", result.BestMatchIndex)
	}

	if result.BestMatch.Target != "sealed" {
		t.Errorf("Wrong best match target, expected `sealed` got `%s`", result.BestMatch.Target)
	}

	if result.BestMatch != result.Ratings[result.BestMatchIndex] {
		t.Error("Results best match is not the same as results ratings item at specified index")
	}

	tables := make(map[string]float32, 4)
	tables["mailed"] = 0.4
	tables["edward"] = 0.2
	tables["sealed"] = 0.8
	tables["theatre"] = 0.36363636363636365

	for _, rating := range result.Ratings {
		score, ok := tables[rating.Target]
		if !ok {
			t.Errorf("Unknown rating result occured (`%s`, `%f`)", rating.Target, rating.Rating)
		}

		if score != rating.Rating {
			t.Errorf("Invalid comparision result (expected `%f` got `%f` when processing `%s`)", score, rating.Rating, rating.Target)
		}
	}
}
