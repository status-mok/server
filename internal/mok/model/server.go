package model

import (
	"fmt"
)

type MOKServer struct {
	Name string `mapstructure:"name"`
	IP   string `mapstructure:"ip"`
	Port uint16 `mapstructure:"port"`
	Type string `mapstructure:"type"`
}

func (s MOKServer) Addr() string {
	return fmt.Sprintf("%s:%d", s.IP, s.Port)
}
