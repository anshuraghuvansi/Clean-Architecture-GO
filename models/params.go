package models

import "user/utils/validators"

//SigninBodyParams :
type SigninBodyParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Validate :
func (s *SigninBodyParams) Validate() error {

	if err := validators.ValidateEmail(s.Email); err != nil {
		return err
	}

	if err := validators.ValidatePassword(s.Password); err != nil {
		return err
	}
	return nil
}

//SignupBodyParams :
type SignupBodyParams struct {
	*SigninBodyParams
	Name string `json:"name"`
}

//Validate :
func (sp *SignupBodyParams) Validate() error {

	if err := validators.ValidateName(sp.Name); err != nil {
		return err
	}

	if err := validators.ValidateEmail(sp.Email); err != nil {
		return err
	}

	if err := validators.ValidatePassword(sp.Password); err != nil {
		return err
	}

	return nil
}

//ToEntity :
func (sp *SignupBodyParams) ToEntity() User {
	return User{Name: sp.Name, Password: sp.Password, Email: sp.Email}
}

//UpdateProfileBodyParams :
type UpdateProfileBodyParams struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

//Validate :
func (up *UpdateProfileBodyParams) Validate() error {

	if err := validators.ValidateName(up.Name); err != nil {
		return err
	}

	if err := validators.ValidatePassword(up.Password); err != nil {
		return err
	}

	return nil
}

//ProfileResponse :
type ProfileResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//AuthResponse :
type AuthResponse struct {
	Token string `json:"token"`
}
