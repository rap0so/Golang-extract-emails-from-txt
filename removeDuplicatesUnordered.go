package main

func removeDuplicatesUnordered(elements []string) (result []string) {
	encountered := map[string]bool{}

	for _, elem := range elements {
		encountered[elem] = true
	}

	for uniqEmail := range encountered {
		result = append(result, uniqEmail)
	}

	return
}
