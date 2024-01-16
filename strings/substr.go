package strings

func FindAll(s string, t string) []int {
	occurrences := make([]int, 0, len(s)-len(t)+1)
	for i, p := range pref(t + "\x00" + s)[len(s)+1:] {
		if p == len(t) {
			occurrences = append(occurrences, i)
		}
	}
	return occurrences[:len(occurrences):len(occurrences)]
}

func FindN(s string, t string, n int) []int {
	occurrences := make([]int, 0, n)
	for i, p := range pref(t + "\x00" + s)[len(s)+1:] {
		if p == len(t) {
			occurrences = append(occurrences, i)
			if len(occurrences) == n {
				return occurrences
			}
		}
	}
	return occurrences[:len(occurrences):len(occurrences)]
}

func pref(s string) []int {
	p := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		k := p[i-1]
		for k > 0 && s[k] != s[i] {
			k = p[k-1]
		}
		if s[k] == s[i] {
			k++
		}
		p[i] = k
	}
	return p
}
