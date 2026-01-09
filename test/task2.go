package main

import "fmt"

type BrainrotMeme struct { name string; trend int; category string; views float64 }

func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
	var r []BrainrotMeme
	for i := 0; i < len(memes); i++ {
		if memes[i].views > minViews {
			r = append(r, memes[i])
		}
	}
	for i := 0; i < len(r); i++ {
		for j := i + 1; j < len(r); j++ {
			if r[j].trend > r[i].trend {
				r[i], r[j] = r[j], r[i]
			}
		}
	}
	return r
}

func CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 {
	m := make(map[string]float64)
	for i := 0; i < len(memes); i++ {
		m[memes[i].category] += memes[i].views
	}
	return m
}

func FilterByComplexCondition(memes []BrainrotMeme) []string {
	var r []string
	for i := 0; i < len(memes); i++ {
		if memes[i].trend >= 7 || (memes[i].views > 50 && memes[i].category == "Sigma") {
			r = append(r, memes[i].name)
		}
	}
	return r
}

func main() {
	memes := []BrainrotMeme{
		{"sigma1", 9, "Sigma", 120.5},
		{"tuntun", 6, "TUNTUNTUNSAHUR", 30.2},
		{"subo", 8, "Subo Bratik", 45.0},
		{"skibidi", 10, "Skibidi", 200.1},
		{"mewing", 7, "Mewing", 60.3},
		{"sigma2", 5, "Sigma", 80.0},
		{"other", 4, "Other", 10.7},
	}

	top := FindTopTrending(memes, 40)
	fmt.Println("топ:")
	for i := 0; i < len(top); i++ {
		fmt.Println(top[i].name, top[i].trend, top[i].views)
	}

	impact := CalculateCategoryImpact(memes)
	fmt.Println("всего:")
	for k, v := range impact {
		fmt.Println(k, v)
	}

	filter := FilterByComplexCondition(memes)
	fmt.Println("фильтр:")
	for i := 0; i < len(filter); i++ {
		fmt.Println(filter[i])
	}
}
