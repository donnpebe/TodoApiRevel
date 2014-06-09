package services

import "github.com/Unknwon/com"

func SanitizeString(value string) string {
	return com.Trim(com.HtmlEncode(value))
}
