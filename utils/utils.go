package utils

func IsValidPassword(password string) bool {
	return len(password) >= 6
}
