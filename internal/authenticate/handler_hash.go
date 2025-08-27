package authenticate

import "golang.org/x/crypto/bcrypt"

func CreateHash(input string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CompareHash(input, hashedInput string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedInput), []byte(input))
	return err == nil
}
