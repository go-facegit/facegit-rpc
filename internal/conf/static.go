package conf

import (
	"net/url"
	"os"
)

var (
	// Application settings
	App struct {
		// ⚠️ WARNING: Should only be set by the main package (i.e. "facegit-web.go").
		Version string `ini:"-"`

		Name    string
		RunMode string
	}

	// log
	Log struct {
		Format   string
		RootPath string
	}

	// Cache settings
	Cache struct {
		Adapter  string
		Interval int
		Host     string
	}

	// web settings
	Rpc struct {
		HttpAddr                 string `ini:"http_addr"`
		HttpPort                 int    `ini:"http_port"`
		Domain                   string
		AppDataPath              string
		AccessControlAllowOrigin string

		ExternalURL          string `ini:"EXTERNAL_URL"`
		Protocol             string
		CertFile             string
		KeyFile              string
		TLSMinVersion        string `ini:"TLS_MIN_VERSION"`
		UnixSocketPermission string
		LocalRootURL         string `ini:"LOCAL_ROOT_URL"`

		OfflineMode      bool
		DisableRouterLog bool
		EnableGzip       bool

		LoadAssetsFromDisk bool

		LandingURL string `ini:"LANDING_URL"`

		// Derived from other static values
		URL            *url.URL    `ini:"-"` // Parsed URL object of ExternalURL.
		Subpath        string      `ini:"-"` // Subpath found the ExternalURL. Should be empty when not found.
		SubpathDepth   int         `ini:"-"` // The number of slashes found in the Subpath.
		UnixSocketMode os.FileMode `ini:"-"` // Parsed file mode of UnixSocketPermission.

		MailSaveMode string
	}

	// Session settings
	Session struct {
		Provider       string
		ProviderConfig string
		CookieName     string
		CookieSecure   bool
		GCInterval     int64 `ini:"gc_interval"`
		MaxLifeTime    int64
		CSRFCookieName string `ini:"csrf_cookie_name"`
	}

	// Other settings
	Other struct {
		ShowFooterBranding         bool
		ShowFooterTemplateLoadTime bool
	}
)

type DatabaseOpts struct {
	Type         string
	Host         string
	Name         string
	User         string
	Password     string
	SslMode      string `ini:"ssl_mode"`
	Path         string
	Charset      string
	MaxOpenConns int
	MaxIdleConns int
}

// Database settings
var Database DatabaseOpts

type i18nConf struct {
	Langs     []string          `delim:","`
	Names     []string          `delim:","`
	dateLangs map[string]string `ini:"-"`
}

// DateLang transforms standard language locale name to corresponding value in datetime plugin.
func (c *i18nConf) DateLang(lang string) string {
	name, ok := c.dateLangs[lang]
	if ok {
		return name
	}
	return "en"
}

// I18n settings
var I18n *i18nConf
