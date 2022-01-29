package app

import (
	// "fmt"
	// "net/http"
	"os"
	"path/filepath"

	"gopkg.in/macaron.v1"

	// "github.com/go-macaron/binding"
	"github.com/go-macaron/cache"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/gzip"
	"github.com/go-macaron/session"

	"github.com/go-facegit/facegit-rpc/internal/app/context"
	"github.com/go-facegit/facegit-rpc/internal/app/router"
	"github.com/go-facegit/facegit-rpc/internal/conf"
)

func newMacaron() *macaron.Macaron {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(gzip.Gziper())
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())

	//template end

	localeNames, _ := conf.AssetDir("conf/locale")
	localeFiles := make(map[string][]byte)
	for _, name := range localeNames {
		localeFiles[name] = conf.MustAsset("conf/locale/" + name)
	}

	m.Use(cache.Cacher(cache.Options{
		Adapter:       conf.Cache.Adapter,
		AdapterConfig: conf.Cache.Host,
		Interval:      conf.Cache.Interval,
	}))

	return m
}

func setRouter(m *macaron.Macaron) *macaron.Macaron {

	//data
	data_path := filepath.Join(conf.WorkDir(), "data")
	os.MkdirAll(data_path, os.ModePerm)

	// reqSignIn := context.Toggle(&context.ToggleOptions{SignInRequired: true})
	// ignSignIn := context.Toggle(&context.ToggleOptions{SignInRequired: conf.Auth.RequireSigninView})
	// reqSignOut := context.Toggle(&context.ToggleOptions{SignOutRequired: true})

	// bindIgnErr := binding.BindIgnErr
	m.SetAutoHead(true)

	m.Group("", func() {
		m.Get("/", router.Home)
	}, session.Sessioner(session.Options{
		Provider:       conf.Session.Provider,
		ProviderConfig: conf.Session.ProviderConfig,
		CookieName:     conf.Session.CookieName,
		CookiePath:     conf.Web.Subpath,
		Gclifetime:     conf.Session.GCInterval,
		Maxlifetime:    conf.Session.MaxLifeTime,
		Secure:         conf.Session.CookieSecure,
	}), csrf.Csrfer(csrf.Options{
		Secret:         "mmmkk",
		Header:         "X-CSRF-Token",
		Cookie:         "facegit",
		CookieDomain:   "facegit",
		CookiePath:     "/",
		CookieHttpOnly: true,
		SetCookie:      true,
		Secure:         false,
	}), context.Contexter())
	return m
}

func Start(port int) {
	m := newMacaron()

	m = setRouter(m)
	m.Run(port)
}
