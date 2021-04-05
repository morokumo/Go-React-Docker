package handler

import (
	"html/template"
)

func EscapeString(str string) string {
	str = template.HTMLEscapeString(str)
	str = template.JSEscapeString(str)
	return str
}
