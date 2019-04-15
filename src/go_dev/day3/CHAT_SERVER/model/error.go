package model

import "errors"

var (
	ErrUserNotExit = errors.New("user not exist")
	ErrInvalidPasswd = errors.New("Passwd or Username not right")
	ErrInvalidParams = errors.New("Invalid params")
	ErrUserExist     = errors.New("User exits")




)