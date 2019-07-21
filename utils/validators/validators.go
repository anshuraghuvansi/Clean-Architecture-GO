package validators

import (
	"errors"
	"regexp"
)

const (
	nameMinLength = 3
	nameMaxLength = 50

	passwordMinLength = 6
	passwordMaxLength = 12
)

var (
	errProvideValidName = errors.New("Provide a valid name")
	errNameMinLengthMSG = errors.New("Name should be a minimum of 3 characters")
	errNameMaxLengthMSG = errors.New("Name should be less than 50 characters")

	errProvideValidPassword = errors.New("Provide valid password")
	errPasswordMinLengthMSG = errors.New("Password length should not be less than 6 characters")
	errPasswordMaxLengthMSG = errors.New("Password length should not be more 12 than characters")

	errInvalidEmail = errors.New("Provide a valid email")
)

//ValidateName :
func ValidateName(name string) error {

	nameLen := len(name)

	if nameLen == 0 {
		return errProvideValidName
	} else if nameLen < nameMinLength {
		return errNameMinLengthMSG
	} else if nameLen > nameMaxLength {
		return errNameMaxLengthMSG
	}
	return nil
}

//ValidatePassword :
func ValidatePassword(password string) error {

	pwdLen := len(password)

	if pwdLen == 0 {
		return errProvideValidPassword
	} else if pwdLen < passwordMinLength {
		return errPasswordMinLengthMSG
	} else if pwdLen > passwordMaxLength {
		return errPasswordMaxLengthMSG
	}
	return nil
}

//ValidateEmail :
func ValidateEmail(email string) error {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	matched := re.MatchString(email)
	if matched {
		return nil
	}
	return errInvalidEmail
}
