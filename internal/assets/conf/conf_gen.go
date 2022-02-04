// Code generated for package conf by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../conf/app.conf
package conf

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _confAppConf = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\xc1\x6e\xe4\x36\x0c\xbd\xeb\x2b\x08\x9f\xbd\x5e\x4f\xb6\xc8\xa6\x33\xd0\x29\xc0\xb6\xc7\xa2\x1b\xa0\x87\x60\xe1\x28\x12\xed\x51\x63\x8b\x02\x45\xcf\x66\xfa\xf5\x85\x64\x7b\x32\xdb\xe6\xd0\x4b\x4f\xb6\x28\xf2\x3d\xf2\x91\x94\x89\xb1\x0b\x66\x42\xd0\x50\xf5\xc6\xe2\xe0\xe5\x03\x47\x5b\x29\x9e\x43\x37\x27\xe4\x7c\xc1\x44\xb2\x58\x26\x72\xc5\xd5\xe1\xa9\x52\xea\x71\xa4\xe1\x9b\x3a\xc0\xc3\x11\x61\xa4\x01\x7a\xe2\xc9\x08\xa0\x97\x23\x32\x54\x7f\x26\x0a\x15\x10\x43\x25\xf8\x2a\x95\x5a\xaf\xf5\x76\xce\xb0\x5d\x34\x72\xcc\xa6\x91\x86\x94\x21\xad\xb1\x47\xdc\x40\xcb\x01\x8c\x33\x51\x90\xeb\x0b\xf0\x84\x13\xf1\xb9\xaa\xa1\x62\x74\x3e\x55\x75\x21\x99\x70\x2a\xfe\x55\xa3\xd6\x08\xd0\xb0\xb8\xaa\x03\x7c\xa1\xb7\x40\xa0\x30\x9e\x6b\xf8\xe5\x1e\x7c\x10\xe4\x93\x19\xc1\x07\x48\x68\x29\xb8\xd4\xa8\x8b\x51\xc3\xee\xa6\xdd\x62\x17\x2a\x30\xc1\x5d\x51\xd5\x60\x29\x04\xb4\xe2\x29\xc0\x91\x92\x80\x71\x8e\x31\xa5\xbd\x3a\xc0\x07\x28\x31\x7b\x08\x28\xdf\x89\x5f\xb4\xd8\x58\xe7\x7b\xbd\xbf\xfd\xf4\xf9\xe7\x3a\x9a\x94\xbe\x13\x3b\x3d\x19\x6b\x98\x42\xed\x9e\x75\x5b\x47\xa2\xb1\x4b\xfe\x2f\xd4\xbb\xb6\xad\xbd\x1b\xb1\x13\x3f\x21\xcd\xa2\x77\x77\x6d\x81\xdd\xe8\xf7\xf0\xb4\xbb\xf9\xdc\xb4\x4d\xdb\xec\xf6\xbb\xdd\xcd\x6e\xf7\xa4\x4a\x12\x5a\xa9\xc7\x84\x29\x79\x0a\x9b\x94\xeb\x11\x22\xd3\xc9\xbb\xf7\xd5\xec\xfd\x88\xab\x98\x4b\xb5\x8d\xda\xdc\xaf\xa5\x2c\x9d\xa1\xd0\xfb\x61\x66\x53\x2a\xef\x89\x81\x31\xc5\x2c\xc4\x09\x2f\x1c\xfb\x2d\x5b\xe2\xf3\x1e\x1c\x61\x82\x40\x02\x01\xd1\x81\x09\xe7\x15\x04\xce\x28\xc5\x31\xd3\xef\x2f\x89\xe6\x13\xe4\xe9\xa8\x01\x9b\xa1\x81\x27\x67\xc4\x7c\x5c\x6f\xd3\xd3\xff\xa5\xef\x96\x7b\xb7\x26\xa7\xe1\x07\xde\x4b\xf9\xf4\xe2\x11\xca\xe6\x08\x41\x12\x62\x04\xb9\x92\xd9\x3b\x0c\xe2\x7b\x8f\xdc\xa8\xc5\x79\x5b\x33\xdf\x8d\xfe\x05\x3b\x3f\x19\x3f\xaa\x03\xfc\x71\xc4\xd2\x86\x8c\x82\xb2\x01\xfb\x00\xbf\x3e\x3c\xfc\xf6\xb5\x4c\xea\x05\x21\xa1\x9d\x39\x63\xf4\x66\x4c\xb8\xa6\xf2\xfe\x14\x97\x96\x6c\xc9\xe4\x0a\x1a\x35\xd8\xee\x6a\xb2\x3f\xdd\xb6\xed\x8a\x30\x99\x57\x3f\xcd\x13\x8c\xbe\x47\xc8\x5a\xfc\x13\xc7\x6c\x48\x8d\x9a\xcc\x6b\x97\xfd\x8a\x66\xa0\xe1\xee\xf6\xa7\x0b\xce\xb5\x28\x39\xec\xfe\xeb\xef\x5f\x40\xe8\x05\x43\xa3\x6c\xe2\xbe\xfb\x51\x88\x2e\xdb\x94\x7a\xe4\x68\xb7\x21\x8d\x4c\x42\x96\x46\x90\xa3\x11\xf0\x09\xe6\x84\x6e\x91\x86\x4f\x08\xce\x33\x5a\x01\x61\xd3\xf7\xde\x66\x7b\x96\xdc\xc4\x38\x7a\x5b\x06\xb1\x51\x07\xb8\x9f\x99\x31\xc8\x78\x86\x34\xc7\x48\x2c\x09\xaa\xa3\x48\xcc\x03\x9e\xbf\xa9\x4c\xba\x1d\xfc\xba\xcb\x73\xf0\xaf\xcb\xa0\x2f\xd4\x1a\xb2\xd7\x9a\xd0\xba\xcd\x99\xea\x19\x61\xf4\x49\x30\xa0\x83\xe7\xf3\xbf\x99\x73\x54\x97\xfd\x41\x43\x5b\x76\xb2\x55\x5b\x59\xc4\x02\x61\x9e\x9e\x97\x3e\xff\x27\xa4\x12\xa3\xe1\xae\xbd\x7b\x83\xc9\x4f\x65\xe9\xab\x10\xfb\x30\x5c\x47\x41\x5e\x3e\x9f\x55\x59\x9a\x9d\xdf\xf5\xfc\xb7\xbd\xaf\xf9\x5f\x65\xb1\x31\xd2\xb7\xb7\x97\x57\x57\x1f\x2b\xa5\xfe\x0e\x00\x00\xff\xff\x31\x03\xa2\x0d\x05\x06\x00\x00"

func confAppConfBytes() ([]byte, error) {
	return bindataRead(
		_confAppConf,
		"conf/app.conf",
	)
}

func confAppConf() (*asset, error) {
	bytes, err := confAppConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "conf/app.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"conf/app.conf": confAppConf,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"conf": &bintree{nil, map[string]*bintree{
		"app.conf": &bintree{confAppConf, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
