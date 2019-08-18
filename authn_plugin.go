package main

import (
	"fmt"
	"github.com/cesanta/docker_auth/auth_server/authn"
)

type PluginAuthnConfig struct {
	PluginPath string `yaml:"plugin_path"`
	Authn      authn.Authenticator
}

type PluginAuthn struct {
	cfg *PluginAuthnConfig
}

func (c *PluginAuthn) Authenticate(user string, password authn.PasswordString) (bool, authn.Labels, error) {
	fmt.Println(user, string(password))
	return true, make(map[string][]string), nil
}

func (c *PluginAuthn) Stop() {
}

func (c *PluginAuthn) Name() string {
	return "plugin auth"
}

var Authn PluginAuthn
