package phoneutil

import "strings"

func HasOnlyDigits(phone string) bool {
	if phone == "" {
		return false
	}

	isNotDigit := func(c rune) bool { return c < '0' || c > '9' }

	return strings.IndexFunc(phone, isNotDigit) == -1
}

func NormalizePhone(phone string) string {
	r := strings.NewReplacer("\\", "", "\"", "", ")", "", "(", "", "^", "", "%", "", "$", "", "#", "", "*", "", "'", "",
		"!", "", "<", "", ">", "", ";", "", ":", "", "/", "", "[", "", "]", "", "{", "", "}", "", "=", "", "~", "", "`", "", "+", "",
		" ", "", "_", "", "-", "", "&", "")
	phone = r.Replace(phone)

	return strings.TrimSpace(phone)
}
