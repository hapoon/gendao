package generator

import (
	"regexp"
	"strings"
)

func ToCamelCaseFromSnakeCase(s string) string {
	rep := regexp.MustCompile(`(_|^)(.)`)
	str := rep.ReplaceAllStringFunc(s, replaceSnakeCaseToUpper)
	return str
}

func replaceSnakeCaseToUpper(s string) string {
	return strings.ToUpper(strings.Trim(s, "_"))
}
