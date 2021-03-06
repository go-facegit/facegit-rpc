package conf

import (
	"github.com/pkg/errors"
	"gopkg.in/ini.v1"

	"github.com/go-facegit/facegit-rpc/internal/assets/conf"
)

// Asset is a wrapper for getting conf assets.
func Asset(name string) ([]byte, error) {
	return conf.Asset(name)
}

// AssetDir is a wrapper for getting conf assets.
func AssetDir(name string) ([]string, error) {
	return conf.AssetDir(name)
}

// MustAsset is a wrapper for getting conf assets.
func MustAsset(name string) []byte {
	return conf.MustAsset(name)
}

// File is the configuration object.
var File *ini.File

func Init(customConf string) error {

	File, err := ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: true,
	}, conf.MustAsset("conf/app.conf"))
	if err != nil {
		return errors.Wrap(err, "parse 'conf/app.conf'")
	}

	File.NameMapper = ini.TitleUnderscore

	if err = File.Append(customConf); err != nil {
		return errors.Wrapf(err, "append %q", customConf)
	}

	// ***************************
	// ----- Log settings -----
	// ***************************
	if err = File.Section("log").MapTo(&Log); err != nil {
		return errors.Wrap(err, "mapping [log] section")
	}

	// ****************************
	// ----- RPC settings -----
	// ****************************
	if err = File.Section("rpc").MapTo(&Rpc); err != nil {
		return errors.Wrap(err, "mapping [web] section")
	}

	// ****************************
	// ----- Repo settings -----
	// ****************************
	if err = File.Section("repo").MapTo(&Repo); err != nil {
		return errors.Wrap(err, "mapping [repo] section")
	}

	if err = File.Section("cache").MapTo(&Cache); err != nil {
		return errors.Wrap(err, "mapping [cache] section")
	} else if err = File.Section("other").MapTo(&Other); err != nil {
		return errors.Wrap(err, "mapping [other] section")
	}

	return nil
}
