package utils

import "regexp"

func ValidatePhoneNumber(str *string) bool {
	// 日本の電話番号の形式に一致する正規表現
	// 例: 03-1234-5678 もしくは +81 90-1234-5678
	regex := `^(\+\d{1,2}\s?)?(\d{1,4}-?)?\d{1,4}-\d{1,4}$`

	match, err := regexp.MatchString(regex, *str)
	if err != nil {
		return false
	}

	return match
}

func ValidateEmail(str *string) bool {
	// 簡易なメールアドレスの正規表現
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, err := regexp.MatchString(regex, *str)
	if err != nil {
		return false
	}

	return match
}
