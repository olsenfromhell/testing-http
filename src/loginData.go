package src

import "github.com/mattn/go-nulltype"

type loginData struct {
	Login         string              `json:"login"`
	Password      string              `json:"password"`
	TwoFactorAuth nulltype.NullString `json:"2fa"`
}

var Body = loginData {

	Login:    "test",
	Password: "test",
}
