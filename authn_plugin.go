package main

import (
	"fmt"
	"github.com/cesanta/docker_auth/auth_server/authn"
)

type authenticator string

func (a authenticator) Authenticate(user, password string) (bool, authn.Labels, error) {
	fmt.Println(user, password)
	return true, make(map[string][]string), nil
}

var Authn authenticator
