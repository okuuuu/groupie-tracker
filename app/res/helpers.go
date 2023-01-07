package res

import (
	"strings"
)

func capitalize(s string) string {
	rns := []rune(s)
	for i, r := range rns {
		if i == 0 || s[i-1] == ' ' {
			rns[i] = r - 32
		}
	}
	res := string(rns)
	res = strings.ReplaceAll(res, "Usa", "USA")
	res = strings.ReplaceAll(res, "Uk", "UK")
	return res
}

func formatConcerts(relation **Relation) {
	updmap := make(map[string][]string)
	for key, val := range (*relation).DatesLocations {
		key = strings.ReplaceAll(key, "_", " ")
		key = strings.ReplaceAll(key, "-", ", ")
		key = capitalize(key)
		updmap[key] = val
	}
	(*relation).DatesLocations = updmap
}
