package util

type Role string
var (
	HackerRole Role = "hacker"
	CompanyRole Role = "company"
)

func CastRole(role string) (Role, bool) {
	switch role {
	case string(HackerRole):
		return HackerRole, true
	case string(CompanyRole):
		return CompanyRole, true
	}

	return "", false
}