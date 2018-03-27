package main

var plurals = []string{"FiEntries"}

func IsPlural(s string) bool {
	for _, p := range plurals {
		if p == s {
			return true
		}
	}

	return false
}
