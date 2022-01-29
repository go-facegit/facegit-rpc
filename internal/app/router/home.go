package router

import (
	"github.com/go-facegit/facegit-rpc/internal/app/context"
)

const (
	HOME = "home"
)

func Home(c *context.Context) {
	c.Data["PageIsHome"] = true
	c.Success(HOME)
}
