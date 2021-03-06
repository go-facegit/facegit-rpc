package router

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/ini.v1"

	"github.com/go-facegit/facegit-rpc/internal/conf"
	"github.com/go-facegit/facegit-rpc/internal/log"
	"github.com/go-facegit/facegit-rpc/internal/tools"
	"github.com/go-facegit/facegit-rpc/internal/tools/debug"
)

func autoMakeCustomConf(customConf string) error {

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

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

	cfg.Section("rpc").Key("http_port").SetValue("8080")
	cfg.Section("session").Key("provider").SetValue("file")

	repoPath := fmt.Sprintf("%s/data/repo", dir)
	cfg.Section("repo").Key("root_path").SetValue(repoPath)

	os.MkdirAll(repoPath, os.ModePerm)
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

	if strings.EqualFold(conf.App.RunMode, "dev") {
		go debug.Pprof()
	}

	return nil
}
