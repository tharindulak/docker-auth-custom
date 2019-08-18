package main

import (
	"fmt"
	"github.com/cesanta/docker_auth/auth_server/authz"
)

type PluginAuthz struct {
	cfg *authz.PluginAuthzConfig
}

func (c *PluginAuthz) Authorize(ai *authz.AuthRequestInfo) ([]string, error) {
	fmt.Printf("Received auth request info: %v\n", ai)
	return ai.Actions, nil
}

func (c *PluginAuthz) Stop() {
}

func (c *PluginAuthz) Name() string {
	return "plugin authz"
}

var Authz PluginAuthz
