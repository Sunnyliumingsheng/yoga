package nets

import (
	"cli/config"
)

type SudoAuthentication struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

var AuthenticationInfo SudoAuthentication

func Init() {
	AuthenticationInfo.Account = config.Config.MyInfo.Account
	AuthenticationInfo.Password = config.Config.MyInfo.Password
}
