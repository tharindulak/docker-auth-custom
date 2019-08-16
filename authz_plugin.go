package main

import (
	"github.com/cesanta/docker_auth/auth_server/authz"
	"log"
)

type authorizer string

func (a authorizer) Authorize(ai *authz.AuthRequestInfo) ([]string, error) {
	log.Printf("Received auth request info: %v", ai)
	return ai.Actions, nil
}

var Authz authorizer
