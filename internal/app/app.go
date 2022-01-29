package app

import (
	// "fmt"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/macaron.v1"

	// "github.com/go-macaron/binding"
	"github.com/go-macaron/cache"
	"github.com/go-macaron/captcha"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/gzip"
	"github.com/go-macaron/i18n"
	"github.com/go-macaron/session"

	"github.com/go-facegit/facegit-rpc/internal/app/context"
	"github.com/go-facegit/facegit-rpc/internal/app/router"
	"github.com/go-facegit/facegit-rpc/internal/app/template"
	"github.com/go-facegit/facegit-rpc/internal/assets/public"
	"github.com/go-facegit/facegit-rpc/internal/assets/templates"
	"github.com/go-facegit/facegit-rpc/internal/conf"
)

func newMacaron() *macaron.Macaron {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(gzip.Gziper())
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())

	m.Use(macaron.Static(
		filepath.Join(conf.CustomDir(), "public"),
		macaron.StaticOptions{
			SkipLogging: true,
		},
	))

	var publicFs http.FileSystem
	if !conf.Web.LoadAssetsFromDisk {
		publicFs = public.NewFileSystem()
	}
	m.Use(macaron.Static(
		filepath.Join(conf.WorkDir(), "public"),
		macaron.StaticOptions{
			FileSystem:  publicFs,
			SkipLogging: false,
		},
	))

	//template start
	renderOpt := macaron.RenderOptions{
		Directory:         filepath.Join(conf.WorkDir(), "templates"),
		AppendDirectories: []string{filepath.Join(conf.CustomDir(), "templates")},
		Funcs:             template.FuncMap(),
		IndentJSON:        macaron.Env != macaron.PROD,
	}
	if !conf.Web.LoadAssetsFromDisk {
		renderOpt.TemplateFileSystem = templates.NewTemplateFileSystem("", renderOpt.AppendDirectories[0])
	}

	m.Use(macaron.Renderer(renderOpt))
	//template end

	localeNames, _ := conf.AssetDir("conf/locale")
	localeFiles := make(map[string][]byte)
	for _, name := range localeNames {
		localeFiles[name] = conf.MustAsset("conf/locale/" + name)
	}
	m.Use(i18n.I18n(i18n.Options{
		CustomDirectory: filepath.Join(conf.CustomDir(), "conf", "locale"),
		Files:           localeFiles,
		Langs:           conf.I18n.Langs,
		Names:           conf.I18n.Names,
		DefaultLang:     "en-US",
		Redirect:        true,
	}))

	m.Use(cache.Cacher(cache.Options{
		Adapter:       conf.Cache.Adapter,
		AdapterConfig: conf.Cache.Host,
		Interval:      conf.Cache.Interval,
	}))

	m.Use(captcha.Captchaer(captcha.Options{
		SubURL: conf.Web.Subpath,
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
