package main

import (
	"fmt"
	"github.com/cesanta/docker_auth/auth_server/authz"
)

type authorizer string

func (a authorizer) Authorize(ai *authz.AuthRequestInfo) ([]string, error) {
	fmt.Printf("Received auth request info: %v", ai)
	return ai.Actions, nil
}

var Authz authorizer
