//go:build !release

package app

import (
	"github.com/status-mok/server/internal/server/config"
)

func (a *app) SetConfig(conf *config.AppConfig) {
	a.conf = conf
}
