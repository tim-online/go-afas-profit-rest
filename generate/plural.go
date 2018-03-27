package main

var plurals = map[string]string{
	"FiEntries":    "FIEntry",
	"FiDimEntries": "FiDimEntry",
	"FiPrjEntries": "FiPrjEntry",
}

func IsPlural(s string) bool {
	_, ok := plurals[s]
	return ok
}

func GetSingular(s string) string {
	if s, _ := plurals[s]; s != "" {
		return s
	}
	return s
}
