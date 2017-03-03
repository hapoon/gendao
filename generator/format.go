package generator

import (
	"strings"
)

func FormatDAOName(daoName string) (result string) {
	result = daoName
	// TODO: If daoName is plural, it is singularized by using dictionary API.
	if strings.HasSuffix(daoName, "s") || strings.HasSuffix(daoName, "es") {
		trimmed := strings.TrimSuffix(daoName, "es")
		if trimmed == daoName {
			trimmed = strings.TrimSuffix(daoName, "s")
		}
		result = trimmed
	}
	return
}
