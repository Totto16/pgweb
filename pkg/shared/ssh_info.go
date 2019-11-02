package shared

import (
	"fmt"
)

// SSHInfo contains ssh server configuration
type SSHInfo struct {
	Host        string `json:"host,omitempty"`
	Port        string `json:"port,omitempty"`
	User        string `json:"user,omitempty"`
	Password    string `json:"password,omitempty"`
	Key         string `json:"key,omitempty"`
	KeyPassword string `json:"keypassword,omitempty"`
}

func (info SSHInfo) String() string {
	return fmt.Sprintf("%s@%s:%s", info.User, info.Host, info.Port)
}
