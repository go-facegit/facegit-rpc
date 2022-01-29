package router

import (
	// "fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/ini.v1"

	"github.com/go-facegit/facegit-rpc/internal/conf"
	"github.com/go-facegit/facegit-rpc/internal/db"
	"github.com/go-facegit/facegit-rpc/internal/log"
	"github.com/go-facegit/facegit-rpc/internal/tools"
	"github.com/go-facegit/facegit-rpc/internal/tools/debug"
)

func autoMakeCustomConf(customConf string) error {

	if tools.IsExist(customConf) {
		return nil
	}

	// Auto make custom conf file
	cfg := ini.Empty()
	if tools.IsFile(customConf) {
		// Keeps custom settings if there is already something.
		if err := cfg.Append(customConf); err != nil {
			return err
		}
	}

	cfg.Section("").Key("app_name").SetValue("facegit-rpc")
	cfg.Section("").Key("run_mode").SetValue("prod")

	cfg.Section("database").Key("host").SetValue("127.0.0.1:3306")
	cfg.Section("database").Key("name").SetValue("facegit")
	cfg.Section("database").Key("user").SetValue("root")
	cfg.Section("database").Key("password").SetValue("root")

	cfg.Section("web").Key("http_port").SetValue("8080")

	cfg.Section("session").Key("provider").SetValue("file")

	os.MkdirAll(filepath.Dir(customConf), os.ModePerm)
	if err := cfg.SaveTo(customConf); err != nil {
		return err
	}

	return nil
}

func Init(customConf string) error {
	var err error

	if strings.EqualFold(customConf, "") {
		customConf = filepath.Join(conf.CustomDir(), "conf", "app.conf")
	} else {
		customConf, err = filepath.Abs(customConf)
		if err != nil {
			return errors.Wrap(err, "custom conf file get absolute path")
		}
	}

	err = autoMakeCustomConf(customConf)
	if err != nil {
		return err
	}

	conf.Init(customConf)
	log.Init()
	db.Init()

	if strings.EqualFold(conf.App.RunMode, "dev") {
		go debug.Pprof()
	}

	return nil
}
