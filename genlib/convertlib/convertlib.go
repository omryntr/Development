package convertlib

import (
	"strings"
)

func Converter(file_scope string, Or_value string, Pg_Value string) string {
	newContents := strings.Replace(string(file_scope), Or_value, Pg_Value, -1)
	return newContents
}
