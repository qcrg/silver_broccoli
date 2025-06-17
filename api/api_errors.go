package api

import "errors"

var (
	ErrInvalidToken           = errors.New("Invalid token")
	ErrNoUserNoRights         = errors.New("No user or no rights")
	ErrNoWalletOrRights       = errors.New("No wallet or no rights")
	ErrNoRights               = errors.New("No rights")
	ErrNoUserNoWalletNoRights = errors.New("No user, no wallet or no rights")
	ErrInternal               = errors.New("Internal error")
	ErrNoWalletType           = errors.New("No wallet type")
)
