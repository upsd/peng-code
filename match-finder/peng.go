package match_finder

func CheckForMatches(terms []string, slangDictionary map[string]string) map[string]string {
	matches := make(map[string]string)

	for _, term := range terms {
		for k, v := range slangDictionary {
			if term == k {
				matches[term] = v
			}
			if term == v {
				matches[term] = k
			}
		}
	}

	return matches
}