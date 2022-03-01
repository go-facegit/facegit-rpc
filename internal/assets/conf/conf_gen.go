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

var _confAppConf = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x4d\x6f\xe4\x36\x0c\xbd\xeb\x57\x10\x3e\x7b\xbd\x76\xb6\xc8\xa6\x33\xf0\x29\xc0\xb6\xc7\xa2\x1b\xa0\x87\x60\xe1\x28\x12\x6d\xab\xb1\x45\x81\xa2\x67\x33\xfd\xf5\x85\x64\x7b\x32\x69\xf7\xd0\x4b\x4f\xb6\x28\xf2\x3d\xea\xf1\x43\x87\xd0\x79\x3d\x23\xb4\x50\xf4\xda\xe0\xe0\xe4\x03\x07\x53\x28\x5e\x7c\xb7\x44\xe4\x74\xc1\x44\xb2\x5a\x66\xb2\xd9\xd5\xe2\xa9\x50\xea\x71\xa2\xe1\x9b\x3a\xc2\xc3\x88\x30\xd1\x00\x3d\xf1\xac\x05\xd0\xc9\x88\x0c\xc5\x9f\x91\x7c\x01\xc4\x50\x08\xbe\x4a\xa1\xb6\xeb\x76\x3f\x27\xd8\x2e\x68\x19\x93\x69\xa2\x21\x26\x48\xa3\xcd\x88\x3b\x68\x3e\x80\xb6\x3a\x08\x72\x79\x01\x9e\x71\x26\x3e\x17\x25\x14\x8c\xd6\xc5\xa2\xcc\x24\x33\xce\xd9\xbf\xa8\xd4\x16\x01\x2d\xac\xae\xea\x08\x5f\xe8\x2d\x10\xc8\x4f\xe7\x12\x7e\xb9\x07\xe7\x05\xf9\xa4\x27\x70\x1e\x22\x1a\xf2\x36\x56\xea\x62\x6c\xa1\xb9\xa9\xf7\xd8\x95\x0a\xb4\xb7\x57\x54\x25\x18\xf2\x1e\x8d\x38\xf2\x30\x52\x14\xd0\xd6\x32\xc6\x78\x50\x47\xf8\x00\x39\xe6\x00\x1e\xe5\x3b\xf1\x4b\x2b\x26\x94\xe9\xbe\x3d\xdc\x7e\xfa\xfc\x73\x19\x74\x8c\xdf\x89\x6d\x3b\x6b\xa3\x99\x7c\x69\x9f\xdb\xba\x0c\x44\x53\x17\xdd\x5f\xd8\x36\x75\x5d\x3a\x3b\x61\x27\x6e\x46\x5a\xa4\x6d\xee\xea\x0c\xbb\xd3\x1f\xe0\xa9\xb9\xf9\x5c\xd5\x55\x5d\x35\x87\xa6\xb9\x69\x9a\x27\x95\x93\x48\x89\xbf\xb3\x2b\xf5\x18\x31\x46\x47\x7e\x97\x76\x3b\x42\x60\x3a\x39\xfb\x63\x75\x7b\x37\xe1\x26\xee\xfa\xfa\x4a\xed\xee\xd7\xd2\xe6\x4a\x91\xef\xdd\xb0\xb0\xce\x4a\xf4\xc4\xc0\x18\x43\x12\xe6\x84\x17\x8e\xc3\x9e\x3d\xf1\xf9\x00\x96\x30\x82\x27\x01\x8f\x68\x41\xfb\xf3\x06\x02\x67\x94\xec\x98\xe8\x0f\x97\x44\xd3\x09\x52\xb7\x94\x80\xd5\x50\xc1\x93\xd5\xa2\x3f\x6e\xb7\xf1\xe9\xff\xd2\x7b\xcf\xbd\xdb\x92\x6b\xe1\x1d\xef\xe5\xf9\xf4\xe2\x10\xf2\x24\x09\x41\x14\x62\x04\xb9\x92\xd9\x59\xf4\xe2\x7a\x87\x5c\xa9\xd5\x79\x1f\x3b\xd7\x4d\xee\x05\x3b\x37\x6b\x37\xa9\x23\xfc\x31\x62\x2e\x43\x42\x41\xd9\x81\x9d\x87\x5f\x1f\x1e\x7e\xfb\x9a\x3b\xf7\x82\x10\xd1\x2c\x9c\x30\x7a\x3d\x45\xdc\x52\xf9\x71\x57\xe7\x92\xec\xc9\xa4\x17\x54\x6a\x30\xdd\x55\xa7\x7f\xba\xad\xeb\x0d\x61\xd6\xaf\x6e\x5e\x66\x98\x5c\x8f\x90\xb4\xf8\x27\x8e\xde\x91\x2a\x35\xeb\xd7\x2e\xf9\x65\xcd\xa0\x85\xbb\xdb\x9f\x2e\x38\xd7\xa2\xa4\xb0\xfb\xaf\xbf\x7f\x01\xa1\x17\xf4\x95\x32\x91\xfb\xee\xbd\x10\x5d\xb2\x29\xf5\xc8\xc1\xec\x4d\x1a\x98\x84\x0c\x4d\x20\xa3\x16\x70\x11\x96\x88\x76\x95\x86\x4f\x08\xd6\x31\x1a\x01\x61\xdd\xf7\xce\x24\x7b\x92\x5c\x87\x30\x39\x93\x1b\xb1\x52\x47\xb8\x5f\x98\xd1\xcb\x74\x86\xb8\x84\x40\x2c\x11\x8a\x51\x24\xa4\x06\x4f\xdf\x98\x3b\xdd\x0c\x6e\x9b\xed\xc5\xbb\xd7\xb5\xd1\x57\xea\x16\x92\xd7\x96\xd0\x36\xdd\x89\xea\x19\x61\x72\x51\xd0\xa3\x85\xe7\xf3\xbf\x99\x53\x54\x97\xfc\xa1\x85\x3a\xcf\x62\xad\xf6\x67\x11\x0b\xf8\x65\x7e\x5e\xeb\xfc\x9f\x90\x72\x4c\x0b\x77\xf5\xdd\x1b\x4c\x5a\x9d\xb9\xae\x42\xec\xfc\x70\x1d\x05\x69\xf8\x5c\x52\x65\x2d\x76\xda\xf3\xe9\x6f\xdf\xb7\xe9\x5f\x25\xb1\x31\xd0\xb7\xb7\x4d\xdc\x16\x1f\x0b\xa5\xfe\x0e\x00\x00\xff\xff\x41\x33\xd9\x8a\x15\x06\x00\x00"

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
