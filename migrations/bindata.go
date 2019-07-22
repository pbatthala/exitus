// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 20190721131113_extensions.down.sql (32B)
// 20190721131113_extensions.up.sql (1.287kB)
// 20190722202741_projects.down.sql (21B)
// 20190722202741_projects.up.sql (334B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __20190721131113_extensionsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xc8\xcb\x57\xc8\x4b\x4d\x4d\x51\x28\xc9\x57\x28\xcd\xcb\xc9\x4f\x4c\x51\x48\xad\x28\x49\xcd\x2b\xce\xcc\xcf\x2b\xe6\x02\x04\x00\x00\xff\xff\x15\xf6\x0f\xb7\x20\x00\x00\x00")

func _20190721131113_extensionsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190721131113_extensionsDownSql,
		"20190721131113_extensions.down.sql",
	)
}

func _20190721131113_extensionsDownSql() (*asset, error) {
	bytes, err := _20190721131113_extensionsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190721131113_extensions.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9b, 0x6, 0x4a, 0xb6, 0xfd, 0xb7, 0xb5, 0xac, 0xd1, 0x74, 0x68, 0xcf, 0x7c, 0xc1, 0xe, 0xc3, 0x3d, 0x44, 0x9b, 0x8c, 0xe1, 0x8d, 0xe7, 0x77, 0x82, 0x3e, 0x7d, 0xa8, 0x4f, 0xf2, 0x74, 0xed}}
	return a, nil
}

var __20190721131113_extensionsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\x4f\x8f\x23\x45\x0c\xc5\xef\xf9\x14\x4f\x73\x02\x29\x7f\xe0\xca\x9e\x10\x04\x69\x04\x9a\x01\x26\x2b\xed\x0d\x79\xaa\x9d\x6e\x6b\xab\xab\x6a\x6d\x57\x7a\xfa\xdb\xa3\xaa\x84\x11\x08\x09\x85\x53\xa2\x96\xed\xf7\xde\xcf\xae\xdd\x0e\xa7\x89\x11\xc4\xf9\xcd\x31\xe7\xa1\x46\x46\xd1\x7c\x91\x81\x0d\x84\x40\xc6\x3b\x49\xc6\xc9\xc4\xe5\xc2\x08\x13\x29\x05\x67\x85\xb9\x4a\x1a\xe1\x6b\xe1\xed\x6d\xc0\x1e\x47\x33\x4e\x2e\x14\xe3\xba\x85\xf8\x66\xb7\x83\x24\x67\x4d\xed\x0b\x02\xc5\x68\x88\x79\x61\xc5\x32\x71\x42\xc8\x73\xa1\x3e\xe6\x42\xb1\xb2\xed\xf1\xec\x13\xeb\x22\xc6\xad\x1d\xaf\x3c\xd1\xa5\x19\x89\x73\x36\x07\xbf\x51\xf0\xb8\x22\xca\x67\x46\x17\x6c\x02\x93\x7b\xb1\xef\x0e\x87\x65\x59\xf6\x25\x9b\x8f\xca\xf6\x25\xee\xb3\x8e\x87\x21\x07\x3b\x7c\xfb\xcd\xe1\x66\x6f\xf2\x39\x6e\x7e\xf8\xfd\xf8\xfd\xe9\x88\xe3\xa7\xd3\xf1\xe9\xe5\xf1\xf9\x09\x8f\x3f\xe1\xe9\xf9\x84\xe3\xa7\xc7\x97\xd3\x0b\x1e\xae\xb5\x0f\x1f\x36\x9b\x0e\x47\xec\x2f\x2c\x32\x97\xc8\x33\x27\x37\xf8\xc4\x98\xcc\xb3\x32\x06\x72\xea\x10\x70\xce\x0d\x4a\xee\x71\x8c\xdd\x90\xcf\xf8\xcc\xeb\xa1\x47\x43\x21\x51\xc3\x22\x3e\x49\x02\xc1\x24\x8d\x91\x9b\xc4\xaf\x57\xcb\x2f\xbf\xfd\x72\x85\xb0\xbf\x8a\x06\x4a\x78\x65\x54\xe3\x73\x8d\x90\x84\x0b\xa9\xe4\x6a\xb0\xc0\xa9\xfd\xb5\x2d\xac\x86\x09\x64\xd0\xbc\x5c\x47\x63\xa6\xb4\x82\xdc\x55\x5e\xab\x73\x33\x4a\x7d\x09\xa4\x0c\x25\xe5\xb8\x36\x88\xb3\x24\x1e\xb6\x68\x7e\x79\x96\x9d\xb9\xd6\xe0\x55\x79\xe8\x69\xf6\xf8\x99\x57\x03\xa5\xe1\xb6\x95\xde\x6d\x2d\xfe\xda\xa9\xdf\x56\x6f\xf7\xd2\xbf\x92\xba\x8f\xfe\xb5\xf6\x9d\x3e\xa3\x8c\x7f\xb8\x8e\xf3\xbf\x6e\xf3\x5c\x53\x70\xc9\xe9\x6a\x34\x17\x56\xf2\xac\xd6\xb7\x30\xb0\xb3\xce\x92\xfa\x7d\x4e\xdd\xbb\x44\x52\xf1\xb5\xed\x84\x62\x99\x28\xd5\x99\x55\x42\x53\xe9\x91\x5e\xc9\x78\x40\x4e\x70\x95\x51\x69\xc6\x4c\x1e\x26\x49\xe3\xb6\x01\x5e\x38\xc6\xf6\x2b\x69\xe0\xb7\x77\x31\x84\x48\x66\x37\xca\xb0\x5a\x4a\x56\xc7\x99\xcc\x61\x4c\xda\xdb\x9b\x9f\x26\x72\xb3\xf0\x7f\xd1\x95\xb1\x85\xbf\x0f\xdd\x8d\xd4\xdf\xd8\xd5\x2a\xc3\x2e\x9b\x95\xff\xa0\xe7\x19\x23\xa7\x16\x88\x51\x93\x5c\x58\xad\xbf\xd5\x9a\xe4\x4b\x65\xc8\xd0\xde\xf3\x59\x58\x0d\x5f\x7d\xfc\xf8\xf8\xa3\x7d\x8d\xda\x8e\xb7\x69\xe4\xc4\x0d\xa8\xf1\x85\x95\x22\xcc\x29\x0d\xa4\x03\x28\x8e\x59\xc5\xa7\xd9\xda\x31\xb3\x72\xbf\x20\x8a\x96\xff\x29\x5c\x34\x0f\x35\x30\x02\xab\x93\x24\x58\xe1\x20\x14\xd1\x74\x10\x72\x6a\x03\xfd\x6e\x56\xef\x69\xef\xc3\xf5\x5e\xfe\xf0\x61\xf3\x67\x00\x00\x00\xff\xff\xaa\x7d\x21\x8c\x07\x05\x00\x00")

func _20190721131113_extensionsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190721131113_extensionsUpSql,
		"20190721131113_extensions.up.sql",
	)
}

func _20190721131113_extensionsUpSql() (*asset, error) {
	bytes, err := _20190721131113_extensionsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190721131113_extensions.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5, 0x65, 0x51, 0x58, 0x2c, 0x49, 0xf0, 0x93, 0xb2, 0xdf, 0x71, 0x97, 0x93, 0x49, 0xe7, 0xc2, 0xd2, 0x28, 0x24, 0x49, 0x23, 0xc7, 0x39, 0xf3, 0x5e, 0xea, 0x3, 0xb7, 0x7, 0xe8, 0x97, 0x30}}
	return a, nil
}

var __20190722202741_projectsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x28\xca\xcf\x4a\x4d\x2e\x29\xb6\xe6\x02\x04\x00\x00\xff\xff\xb0\x5e\x15\x2c\x15\x00\x00\x00")

func _20190722202741_projectsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190722202741_projectsDownSql,
		"20190722202741_projects.down.sql",
	)
}

func _20190722202741_projectsDownSql() (*asset, error) {
	bytes, err := _20190722202741_projectsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190722202741_projects.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x80, 0x1a, 0x54, 0xa4, 0x9c, 0xdd, 0xe7, 0x4e, 0x5d, 0x89, 0x95, 0xa9, 0x1b, 0xe1, 0xca, 0xd, 0x4a, 0xa0, 0x71, 0x53, 0x74, 0xa0, 0x88, 0xfe, 0x4d, 0x2f, 0xe0, 0x41, 0xd1, 0x9c, 0xd0, 0x9d}}
	return a, nil
}

var __20190722202741_projectsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\x41\x4b\xc3\x30\x1c\x05\xf0\x7b\x3f\xc5\x23\x97\xb5\xe0\xd1\xd3\x3c\x55\x8d\x30\xac\x55\x4b\x7b\x18\x22\x25\x36\x7f\x34\xb2\x26\x21\xf9\xd7\x8a\xe2\x77\x17\xd7\x6c\x9e\x77\x7c\x79\xef\x97\xe4\xaa\x91\x65\x2b\xd1\x96\x97\x95\x84\x0f\xee\x9d\x06\x8e\xc8\x33\x00\x10\x46\x0b\x4c\x93\xd1\xb8\x96\x37\x65\x57\xb5\xfb\xd0\xbf\x92\xa5\xa0\x98\xfa\x8f\xf3\xbc\xc0\x43\xb3\xb9\x2b\x9b\x2d\x6e\xe5\xf6\x6c\x51\x56\x8d\x24\x30\x18\xa6\x4f\x46\x7d\xdf\xa2\xee\xaa\x0a\x5d\xbd\x79\xec\x64\x9a\x68\x8a\x43\x30\x9e\x8d\xb3\x02\x7f\xbb\x74\xee\x66\x4b\xa1\x3f\x3e\x7b\xc0\xa9\xdd\xa9\x17\xda\xc5\x05\x3c\x3d\xff\x5f\x7d\xf8\xdd\xea\xfb\x67\xb5\x5e\x2f\x6d\x22\x43\x20\xc5\xa4\x7b\xc5\x02\x6c\x46\x8a\xac\x46\x8f\xd9\xf0\xdb\x3e\xe2\xcb\x59\x3a\x7a\xeb\xe6\xbc\x48\x70\xf2\xfa\x44\x98\x15\x17\xd9\x6f\x00\x00\x00\xff\xff\x5b\xfa\xf9\xd2\x4e\x01\x00\x00")

func _20190722202741_projectsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20190722202741_projectsUpSql,
		"20190722202741_projects.up.sql",
	)
}

func _20190722202741_projectsUpSql() (*asset, error) {
	bytes, err := _20190722202741_projectsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20190722202741_projects.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x77, 0x6f, 0xec, 0x84, 0xed, 0x19, 0x68, 0x35, 0x73, 0x82, 0xad, 0x2f, 0x5b, 0x85, 0xd0, 0xae, 0xf3, 0x5a, 0xe6, 0xff, 0x56, 0x9, 0x2, 0xea, 0x4e, 0x53, 0x78, 0x99, 0xff, 0xaf, 0xa5, 0xf2}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"20190721131113_extensions.down.sql": _20190721131113_extensionsDownSql,

	"20190721131113_extensions.up.sql": _20190721131113_extensionsUpSql,

	"20190722202741_projects.down.sql": _20190722202741_projectsDownSql,

	"20190722202741_projects.up.sql": _20190722202741_projectsUpSql,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"20190721131113_extensions.down.sql": {_20190721131113_extensionsDownSql, map[string]*bintree{}},
	"20190721131113_extensions.up.sql":   {_20190721131113_extensionsUpSql, map[string]*bintree{}},
	"20190722202741_projects.down.sql":   {_20190722202741_projectsDownSql, map[string]*bintree{}},
	"20190722202741_projects.up.sql":     {_20190722202741_projectsUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}