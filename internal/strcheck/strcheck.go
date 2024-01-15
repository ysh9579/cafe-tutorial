package strcheck

import "regexp"

const (
	// RegexPhone 핸드폰 번호 검증
	RegexPhone = `^01([0|1|6|7|8|9])-(\d{3,4})-(\d{4})$`

	// RegexPassword 비밀번호 검증
	// 최소 8자 이상의 영문이나 숫자로 이루어져야 한다
	RegexPassword = `[a-zA-Z0-9]{8,}$`
)

var (
	PhoneRegexp    = regexp.MustCompile(RegexPhone)
	PasswordRegexp = regexp.MustCompile(RegexPassword)
)

func ValidatePhone(phone string) bool {
	return PhoneRegexp.MatchString(phone)
}

func ValidatePassword(pwd string) bool {
	return PasswordRegexp.MatchString(pwd)
}
