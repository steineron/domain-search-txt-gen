package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	/*commonPhrases := []string{
		"for sale",
		"for sale in",
		"for rent",
		"for rent in",
		"sold",
		"sold in",
	}*/

	commonSearchTerms := createCommonSearchTerms()

	fmt.Print(commonSearchTerms)
	fmt.Printf("Total phrases %v", len(commonSearchTerms))
}

type countableFeatures struct {
	min      int
	max      int
	keywords []string
}

func createCommonSearchTerms() []string {

	beds := []string{"bed", "beds", "bedroom", "bedrooms"}
	baths := []string{"bath", "baths", "bathroom", "bathrooms"}
	cars := []string{"car", "cars", "carparks", "parkings"}

	var phrases []string

	phrases = generateFeaturePhrases(
		&countableFeatures{1, 3, beds},
		&countableFeatures{1, 2, baths},
		&countableFeatures{1, 2, cars},
	)
	phrases = append( phrases, generateFeaturePhrases(
		&countableFeatures{4, 6, beds},
		&countableFeatures{2, 3, baths},
		&countableFeatures{2, 4, cars},
	)...)
	return phrases

}

func generatePhrases(prefix []string, features *countableFeatures) []string {

	var phrases []string
	for _, prefix := range prefix {
		for _, kw := range features.keywords {
			for k := features.min; k <= features.max; k++ {
				builder := strings.Builder{}
				if len(prefix) > 0 {

					builder.WriteString(prefix)
					builder.WriteString(" ")
				}
				builder.WriteString(strconv.Itoa(k))
				builder.WriteString(" ")
				builder.WriteString(kw)

				phrases = append(phrases, builder.String())
			}
		}
	}
	return phrases
}

func generateFeaturePhrases(major *countableFeatures, minor *countableFeatures, option *countableFeatures) []string {

	var phrases []string

	majorFeatures := append(generatePhrases([]string{""}, major)) // x beds
	majorMinor := generatePhrases(majorFeatures, minor)           // x beds y baths

	phrases = append(phrases, majorFeatures...)
	phrases = append(phrases, majorMinor...)

	phrases = append(phrases, generatePhrases(majorFeatures, option)...)
	phrases = append(phrases, generatePhrases(majorMinor, option)...)

	return phrases
}
