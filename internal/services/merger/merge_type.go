package merger

func nonZeroFloat(floats ...float64) float64 {
	for _, float := range floats {
		if float != 0.0 {
			return float
		}
	}
	return 0.0
}

func nonEmptyString(strings ...string) string {
	for _, val := range strings {
		if val != "" {
			return val
		}
	}
	return ""
}

func nonZeroInt(ints ...int) int {
	for _, val := range ints {
		if val != 0 {
			return val
		}
	}
	return 0
}

// merge all string with a space in between
func mergeString(strings ...string) string {
	result := ""
	for _, str := range strings {
		result += " " + str
	}
	return result
}

// select only longest string as result
func longestString(strings ...string) string {
	longestStr := ""
	for _, str := range strings {
		if len(str) > len(longestStr) {
			longestStr = str
		}
	}
	return longestStr
}

// merge list of string slices to a slice which contains unique values
func uniqueString(stringSlices ...[]string) []string {
	uniqueMap := map[string]bool{}

	for _, intSlice := range stringSlices {
		for _, number := range intSlice {
			uniqueMap[number] = true
		}
	}

	// Create a slice with the capacity of unique items
	// This capacity make appending flow much more efficient
	result := make([]string, 0, len(uniqueMap))

	for key := range uniqueMap {
		result = append(result, key)
	}

	return result
}
