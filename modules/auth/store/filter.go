package authstore

func GetEmailFilter(email string) map[string]interface{} {
	return map[string]interface{}{
		"email": email,
	}
}
