package error

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrPasswordIncorrect = errors.New("password incorrect")
	ErrUsernameExist = errors.New("username already exist")
	ErrUserDoesNotMatch = errors.New("user does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordIncorrect,
	ErrUsernameExist,
	ErrUserDoesNotMatch,
}