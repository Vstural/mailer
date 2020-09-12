package outlook

import (
	"fmt"
	"strings"
)

const contentType = "Content-Type: text/plain" + "; charset=UTF-8"

func ToStr(to []string) string {
	return strings.Join(to, ";")
}

func OutLookBody(from string, to []string, subject, content string) string {
	tostr := ToStr(to)
	return fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\n%s\r\n\r\n%s",
		tostr, from, subject, contentType, content)
}
