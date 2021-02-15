// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package tmpl generated by go-bindata.// sources:
// templates/base.layout.gohtml
// templates/call.layout.gohtml
// templates/call.page.gohtml
// templates/debug-call.layout.gohtml
// templates/debug.page.gohtml
// templates/help.page.gohtml
// templates/home.page.gohtml
// templates/menu.layout.gohtml
// templates/search.page.gohtml
package tmpl

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

// ModTime return file modify time
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

var _baseLayoutGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x91\x41\x8e\x9d\x30\x0c\x86\xf7\x9c\xc2\xe3\xf5\x04\xd4\x5d\x17\x09\x9b\xde\xa0\x73\x82\xbc\xc4\x80\x35\xc1\xa1\x89\xdf\x9b\x22\xc4\xdd\x2b\x60\xda\x79\xea\x2a\xca\xf7\xfb\xb7\x7f\xd9\xdb\x16\x69\x60\x21\xc0\x9b\xaf\x84\xfb\xde\xd8\x97\x98\x83\xae\x0b\xc1\xa4\x73\xea\x1b\x7b\x3c\x90\xbc\x8c\x0e\x49\xf0\x00\xe4\x63\xdf\x00\x00\xd8\x17\x63\xe0\x27\xfd\xba\x73\xa1\x08\x33\xa9\x07\xf5\x63\x05\x63\x3e\xf5\x13\x85\xc9\x97\x4a\xea\xf0\xae\x83\xf9\x8e\xcf\x92\xf8\x99\x1c\x3e\x98\x3e\x96\x5c\x14\x21\x64\x51\x12\x75\xf8\xc1\x51\x27\x17\xe9\xc1\x81\xcc\xf9\x79\x05\x16\x56\xf6\xc9\xd4\xe0\x13\xb9\x6f\xaf\x50\xa7\xc2\xf2\x6e\x34\x9b\x81\xd5\x49\xc6\xbe\xf9\x8a\xf5\xe3\xed\x0d\xb2\xa4\xf5\x2b\x4c\x62\x79\x87\xa9\xd0\xe0\xb0\xab\xea\x95\x43\x77\xcb\x59\xab\x16\xbf\xb4\x33\x4b\x1b\x6a\x45\x28\x94\x1c\x56\x5d\x13\xd5\x89\xe8\xc8\x54\x72\xad\xb9\xf0\xc8\xe2\xd0\x4b\x96\x75\xce\xf7\xfa\x6f\xd6\xd9\xf5\x34\x71\xc8\x82\xff\x0d\x18\xfc\xe3\xc0\xed\x22\x23\xc2\xb1\x55\x87\x3c\xfb\x91\xba\xdf\xe6\x2c\xff\xdb\x45\x59\x13\xf5\xdb\xa6\x34\x2f\xc9\x2b\x01\x9e\x04\xa1\xdd\x77\xdb\x5d\x6a\x63\xbb\x6b\xf5\xf6\x96\xe3\xfa\xe9\x7c\xb6\x1c\xf8\x74\x34\x8d\xed\xae\x1a\xdb\x5d\x57\xdc\x36\x92\xb8\xef\x7f\x02\x00\x00\xff\xff\xfc\x01\xc3\xc2\xef\x01\x00\x00")

func baseLayoutGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_baseLayoutGohtml,
		"base.layout.gohtml",
	)
}

func baseLayoutGohtml() (*asset, error) {
	bytes, err := baseLayoutGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "base.layout.gohtml", size: 495, mode: os.FileMode(420), modTime: time.Unix(1613403086, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _callLayoutGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\xdd\x6e\xdb\x36\x14\xbe\xef\x53\x1c\xf0\xc6\x1b\x50\x1e\x91\x92\x28\x4b\x45\x94\x61\x6d\x2e\x32\xa0\xeb\xc5\xba\xf6\x5e\x91\x98\x88\xa8\x4c\x79\x92\x62\xb7\x33\x74\x91\x0c\xc3\x30\xa0\xd8\xde\x65\x48\x80\x2c\x59\xfc\x0c\xd4\x1b\x0d\x94\x9d\x58\xf9\x43\xb6\xae\xc0\x0c\x58\x20\x79\xce\xf9\xce\x47\x7e\x9f\x7e\x16\x8b\x4c\xee\x2b\x2d\x81\xa4\x49\x51\x90\xb6\x7d\xf2\x04\x00\x60\x2b\x53\x33\x48\x8b\xa4\xae\x63\x92\x96\xba\x49\x94\x96\x15\x4c\x1a\xea\x91\xed\x3e\xa1\x4f\xca\x83\xed\xc5\x02\x9f\x27\xb5\x7c\x91\x14\x05\xbe\x79\xf3\xcd\x0e\xbe\x6e\x2a\xa5\x0f\xda\x76\xcb\xc9\x83\xdb\xa9\x6a\x1f\xe4\x0f\xb0\x29\xd8\x51\x95\x4c\x1b\x55\x6a\x20\x4a\xa7\xa4\x6d\xcd\x1f\xdd\xcf\x66\x69\x4e\xba\xdf\xbb\x5f\xcd\x99\xf9\x73\xb1\x90\x45\x2d\xdb\xd6\x9c\x75\x47\x77\x22\x3a\x6b\x5b\x18\xf6\x7f\xdd\x24\x55\xf3\xba\x49\x26\xd3\xfb\xda\x9b\xf3\xee\xd8\x2c\x9f\xdd\xa8\x78\x53\xcb\x4a\x27\x13\xd9\xb6\x60\xce\xcd\xd2\xfc\xd5\xfd\x74\x33\x61\x47\xd6\x8d\xd2\x89\xe5\xf8\xea\x70\xb2\x27\xab\xfb\x90\xbb\x23\xb3\x34\xa7\xe6\xc4\x9c\x99\x4b\x73\x6a\x2e\xcd\x99\x39\xed\x61\x5e\x94\x5a\xcb\xb4\xf9\x5e\xf5\x1d\xba\x23\x73\x6a\xce\x11\xcc\x89\xb9\x30\x67\xdd\xb1\x39\x35\x17\xdd\x47\x73\x69\x96\xdd\x51\x77\xdc\x7d\xbc\xd9\xf9\xb9\x2a\x8a\x5a\xa6\x9b\xba\x7b\x76\xb4\xec\x8e\xcd\xb9\xb9\xe8\x7e\xeb\x7e\xb9\xd1\x77\x47\xd5\xe9\xaa\xf5\x4d\xba\xbd\x02\x5a\x02\x7e\x27\xd3\xb2\xca\x5e\x25\x13\x09\xc4\x6a\x0e\xd7\xbf\xad\xe9\xf6\x60\xd6\xaf\x24\x90\x57\x72\x3f\x26\xe4\xca\x10\x7b\x8d\x86\xbd\x46\xd3\x69\xa5\x26\x49\xf5\x81\x40\xa9\xd3\x42\xa5\xef\x62\x32\x57\x3a\x2b\xe7\x58\x4e\xa5\xfe\x62\xe4\x24\x53\xe5\xcc\xb8\x53\xa8\xba\x91\xfa\xab\xc3\x43\x95\xc5\x0f\xda\x65\xf4\x74\x34\x99\x7a\xa3\xa7\xa3\xb9\xca\x9a\x3c\xf6\x18\x7b\x9a\x4b\x75\x90\x37\xb1\xcb\xd8\xe8\x4b\x72\x9b\x55\xcf\xac\x9e\x1d\xc0\xfb\x49\xa1\xeb\x98\xe4\x4d\x33\x7d\xe6\x38\xf3\xf9\x1c\xe7\x1e\x96\xd5\x81\xe3\x32\xc6\x9c\x7a\x76\x40\x60\x05\x49\x78\x40\x60\x8d\xd9\x8f\xf7\x55\x51\xc4\x24\x3d\xac\x2a\xa9\x9b\x17\x65\x51\x56\x9b\x1d\x2a\xd8\x53\x34\x4d\x2a\xd9\xd0\xca\x56\x50\x9b\x4c\x60\xa6\xe4\xfc\x79\xf9\x3e\x26\x0c\x18\xf0\x00\x78\x70\x2f\xb1\xd5\x41\x26\x4d\x0e\x59\x4c\xbe\xe5\x2e\x72\x1f\x42\x1c\x0b\xaf\xa0\x02\xfd\xd0\x05\x1f\xc7\x51\x90\x52\x0c\xfc\x00\x45\x10\x50\x8e\x81\x08\x91\xb3\xf5\x88\xda\xdc\xb7\x1e\xba\xcc\x4f\x38\x70\xe8\xbb\x81\x0d\x45\x7d\xa8\xb0\x28\x2b\x90\x41\x9c\x01\x47\xc1\x82\x1f\x89\x73\xef\x61\xd9\xb3\xb8\xa3\xad\x93\x3c\x28\xf7\xbf\x92\xef\x61\x6f\x64\xe5\x5c\x17\x65\x92\xc5\x64\xb1\x18\xf8\xae\x6d\xff\x17\x49\xaf\xd8\x7c\x82\x94\x16\x9b\x56\x87\x85\x8c\x89\x9c\x49\x5d\x66\x19\xe9\xe5\x45\x01\x11\x46\x09\x0a\x14\x6b\x21\xec\x70\xe6\xa2\xb8\xd6\xa6\x57\x2f\xe7\xee\x70\x81\xf2\x19\xb5\x39\x83\x3a\x0e\xac\x2f\x73\xc1\x5d\xe7\xb8\xe0\xee\xba\xc3\x39\x75\xef\x54\xa1\xa0\x28\x1e\x10\xfd\x71\xf6\x63\xeb\x41\xe0\x1c\x43\xe1\x6f\x50\x19\xe0\x98\x85\xc0\x0a\x8f\x7a\x83\x55\x6a\x57\xfb\xcb\xcb\x10\x05\x70\x86\x6e\xe4\xbd\xe5\x43\x3e\x8c\xda\x5d\x84\x38\x8e\xbc\x97\x02\x3d\x61\x8d\xcf\xfd\xe0\x2a\x81\xaf\x41\xec\xbf\xf0\xc0\xfb\x6c\x66\x9d\x27\x1f\x78\x18\x3a\xf6\x25\xf6\xa9\x56\xfd\x0c\x7e\x74\xd9\xc6\x8f\x76\xfc\xa8\x1f\x0b\xa5\xdf\xdd\xf6\xa2\x00\x2e\xfe\xc1\x63\x25\xe8\x4f\x57\xa0\xd8\xf5\x13\x0f\xbc\xb5\x6e\x0c\x82\xdc\x1b\xcc\x5d\x0c\x3d\xea\xef\x46\x29\x45\x16\x06\xf6\xf4\xf9\x18\x19\xa7\xe8\x0a\x64\x1e\xff\xfa\xda\x5b\x30\xb6\x7a\x5a\xb0\xd5\x52\x6f\x47\xea\xe7\x1c\x85\x27\x52\x74\x79\x48\xd1\x1b\x07\xe8\x47\x82\xe2\x98\xfb\x18\xba\x94\x3f\x6a\x3b\x4b\x34\xb2\x24\x37\x94\xa8\xa5\x04\x16\x98\x45\xe1\xa0\x7f\x04\x01\x8a\xdc\x1b\xb6\x07\x3f\xa7\x7d\xff\xc4\x47\x66\x9f\x9a\xec\xea\x46\xc0\xd0\x05\xbe\xcb\xdd\x35\x6c\x7f\x57\xd1\x60\x37\xfa\x0f\x7e\xda\x72\xa6\x9b\xf7\x64\xff\x71\xb1\xfa\x12\x72\x32\x35\x5b\x05\xb6\xf2\x6a\xfb\xc9\x3a\xf6\x77\x00\x00\x00\xff\xff\x76\xff\x58\xc6\x38\x09\x00\x00")

func callLayoutGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_callLayoutGohtml,
		"call.layout.gohtml",
	)
}

func callLayoutGohtml() (*asset, error) {
	bytes, err := callLayoutGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "call.layout.gohtml", size: 2360, mode: os.FileMode(420), modTime: time.Unix(1613492284, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _callPageGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x50\xcd\x4a\x03\x31\x10\xbe\xf7\x29\x86\xb9\xdb\x1e\x82\xb7\xba\x17\x9f\x24\xdd\x4c\x49\x60\x9a\x4a\x13\x0a\x32\xe4\x9d\x44\x5c\x44\x61\xf7\x19\xe6\x91\x24\xae\xab\x86\xa5\xdf\x21\x87\x2f\xdf\x5f\x22\x02\x99\x4e\x4f\x6c\x33\x01\x1e\x6c\x22\x84\x6d\x29\x9b\x8d\x88\xa3\x63\x88\x04\x98\x43\x66\xc2\x52\x1e\x2d\xb3\x08\x45\xd7\x5e\x1f\xce\xee\x19\x2b\x05\x00\xb0\x77\xe1\x0a\x3d\xdb\x94\x1e\xb0\x3f\xc7\x6c\x43\xa4\x0b\x9c\xf2\x9d\xc1\x6e\x56\x54\x88\x84\x23\x6c\x6b\xde\xe2\x5b\x70\xd3\x7f\x8f\x5d\x23\xfc\x16\xfb\xcb\x9a\x14\xf9\x7b\x4e\x6f\x99\xf1\xb7\xa8\xb1\xee\x5c\xb8\x36\x8b\x88\x13\xad\xd6\x78\xb3\x8c\xf1\x06\x3b\x7d\xd7\x57\x9d\x74\xd4\x49\x3f\x41\x47\x1d\xea\xf1\xa2\x1f\xfa\xa6\x83\x8e\xfb\x9d\x37\x6d\xe2\xfc\x53\xff\xeb\x7e\xc8\xaf\x00\x00\x00\xff\xff\x2a\xf8\x69\xaf\x74\x01\x00\x00")

func callPageGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_callPageGohtml,
		"call.page.gohtml",
	)
}

func callPageGohtml() (*asset, error) {
	bytes, err := callPageGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "call.page.gohtml", size: 372, mode: os.FileMode(420), modTime: time.Unix(1613416708, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _debugCallLayoutGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\xdd\x4e\xe3\x46\x14\xbe\xe7\x29\x8e\xe6\x26\xad\xc4\x1c\xcf\xf8\x37\x5e\x61\xaa\xee\x72\xd1\x4a\xec\x5e\x74\x0b\xf7\xc6\x1e\xf0\x68\x27\x76\xea\x38\x09\x25\xf2\x05\x54\x55\x55\x69\xd5\xbe\x4b\x05\x12\x85\x92\x67\x18\xbf\x51\x35\x4e\xc8\x0f\x84\x56\xa2\x95\xd8\x8b\x45\x22\x1a\x9f\x9f\xef\x9c\x99\xef\xcb\x19\x67\x32\x49\xc5\xb1\xcc\x05\x90\x54\x1c\x0d\x4f\x68\x12\x2b\x45\xea\x7a\x6b\x0b\x00\x60\x27\x95\x23\x48\x54\x3c\x18\x44\x24\x29\xf2\x2a\x96\xb9\x28\xa1\x57\x51\x87\xec\xb6\x01\x6d\x50\xe6\xef\x4e\x26\x78\x70\xf0\xed\x5e\x5d\xef\x58\x99\xff\xd0\x25\x8f\x41\xfc\x00\xf8\x26\x56\x0a\xf7\x64\x29\x92\x4a\x16\x39\x10\x99\x27\xa4\xae\xf5\x1f\xcd\xcf\x7a\xaa\x2f\x9b\xdf\x9b\x5f\xf5\xb5\xfe\x73\x32\x11\x6a\x20\xea\x5a\x5f\x37\xe7\x8f\x3c\x79\x5a\xd7\x30\x99\xcc\xa0\xde\x57\x71\x59\xbd\xaf\xe2\x5e\x7f\x53\x59\x7d\xd3\x5c\xe8\xe9\xab\x45\xf4\xc1\x40\x94\x79\xdc\x13\x75\x0d\xfa\x46\x4f\xf5\x5f\xcd\x4f\x4b\xe7\x9e\x18\x54\x32\x8f\x4d\x5f\xef\x86\xbd\x23\x51\x6e\x42\x6c\xce\xf5\x54\x5f\xe9\x4b\x7d\xad\xef\xf4\x95\xbe\xd3\xd7\xfa\x6a\x06\x51\xe4\xf9\x6c\x53\xdf\xcb\xb6\x40\x73\xae\xaf\xf4\x0d\x82\xbe\xd4\xb7\xfa\xba\xb9\xd0\x57\xfa\xb6\xf9\xa8\xef\xf4\xb4\x39\x6f\x2e\x9a\x8f\xcb\xc2\xaf\xa5\x52\x03\x91\x2c\x73\x36\x6c\x64\xda\x5c\xe8\x1b\x7d\xdb\xfc\xd6\xfc\xb2\x56\x76\x4f\x0e\x92\x59\xe5\xf5\x6e\xdb\x03\xcf\x05\xe0\x77\x22\x29\xca\xf4\x5d\xdc\x13\x40\x0c\xa5\xb0\xf8\xdb\xe9\xef\xae\x3c\xb5\x96\x18\xb2\x52\x1c\x47\x84\xdc\xf3\x7d\x54\xe5\x70\x54\xe5\xb4\x5f\xca\x5e\x5c\xfe\x48\xa0\xc8\x13\x25\x93\x0f\x11\x19\xcb\x3c\x2d\xc6\x58\xf4\x45\xfe\x45\xc7\x8a\xfb\xd2\x1a\x71\x4b\xc9\x41\x25\xf2\xaf\x86\x43\x99\x46\x0b\x35\x74\xb6\x3b\xbd\xbe\xd3\xd9\xee\x8c\x65\x5a\x65\x91\xc3\xd8\x76\x26\xe4\x49\x56\x45\x36\x63\x9d\x2f\xc9\xc3\x2e\xda\x4e\x06\xa3\x13\x38\xed\xa9\x7c\x10\x91\xac\xaa\xfa\xaf\x2c\x6b\x3c\x1e\xe3\xd8\xc1\xa2\x3c\xb1\x6c\xc6\x98\x35\x18\x9d\x10\x98\x41\x12\xee\x13\x98\x63\xb6\xeb\x63\xa9\x54\x44\x92\x61\x59\x8a\xbc\x7a\x53\xa8\xa2\x5c\xee\x48\xc2\x91\xa4\x49\x5c\x8a\x8a\x96\x26\x83\x9a\x60\x02\x23\x29\xc6\xaf\x8b\xd3\x88\x30\x60\xc0\x7d\xe0\xfe\xc6\xc6\x66\x07\x17\x57\x19\xa4\x11\x79\xcb\x6d\xe4\x2e\x74\x31\xf0\x1c\x45\x3d\x74\xbb\x36\xb8\x18\x84\x7e\x42\xd1\x77\x7d\xf4\x7c\x9f\x72\xf4\xbd\x2e\x72\x36\x5f\x51\x13\x7b\xe8\xa0\xcd\xdc\x98\x03\x87\xb6\x1a\x18\x57\xd8\xba\x94\x41\x99\x81\xac\xf8\x19\x70\xf4\x98\x7f\x46\xac\x8d\x87\x65\xce\xe2\x11\x97\x56\xfc\x24\xbd\xff\x48\xd7\xd3\xdc\xa7\xc5\x38\x57\x45\x9c\x46\x64\x32\x59\xd1\x55\x5d\xbf\x08\x85\xf7\xdd\x3c\x83\x3a\x83\x4d\xcb\xa1\x12\x11\x11\x23\x91\x17\x69\x4a\x5a\x3a\xd1\x83\x10\xc3\x18\x3d\xf4\xe6\x07\x6f\x96\x23\x1b\xbd\x05\x17\x2d\x5b\x19\xb7\x57\x0d\x94\x8f\xa8\x89\x59\xc9\xe3\xc0\xda\x34\x1b\xec\x79\x8c\x0d\xf6\x37\xf6\xea\x33\xb5\x1f\x65\xa1\x47\xd1\x7b\x82\xe4\x7f\xef\x3e\x30\x9a\x03\xce\xb1\xeb\xb9\x4b\x54\x06\x18\xb0\x2e\x30\xe5\x50\x67\xc5\x4a\x8d\xb5\xfd\xd8\xef\xa2\x07\x9c\xa1\x1d\x3a\x87\x7c\xb5\x1f\x46\xcd\x2e\xba\x18\x84\xce\xbe\x87\x8e\x67\x84\xce\x5d\xff\x3e\x80\xcf\x41\xcc\xbf\x72\xc0\xf9\xbf\xc4\xf9\xdc\xd9\x93\xa4\xe5\x69\x4f\x3d\x9e\x3d\x49\x5a\x2e\x66\x8f\xbf\x9c\x3d\xee\x8b\xcd\x9e\x63\xa9\x04\x15\x71\xd9\x8b\xcb\x0f\x74\x5c\x94\xe9\x73\x27\xd0\x53\x42\xb0\x61\x45\x68\x60\x53\x3b\xf3\x0c\xbb\x5f\xaf\xcc\x1b\x06\xc6\xb2\xcf\x1d\x0c\x58\x00\xab\xa3\xc8\xd8\x8d\xf1\x90\xbb\x0f\xd4\xeb\xae\xab\xf7\xd0\x3e\xeb\x05\x46\x3a\xe8\x8d\xa8\x6d\x14\x90\xd1\xe5\xd7\x82\x53\x4e\xf9\xd9\x5b\x33\xcd\x3c\xf0\xb1\x1b\x84\x6b\xba\x09\x03\xb4\x5d\x5b\x71\xf4\xc0\x5f\xd3\x6a\xe8\x07\xc8\xf8\x7e\x17\x42\x74\x99\x09\x60\xbc\x0b\x0e\x06\xce\xc3\x28\x8a\x8c\x9b\x7c\xea\xaf\xa9\x3a\x0c\x68\x8b\x4c\x39\x32\xc7\x07\x17\xb9\xeb\x52\x0c\xc3\x80\x3a\xe8\x7b\xde\x7a\xac\xef\x02\x53\xad\x17\x5a\xef\xfe\xa2\xdd\xee\x7f\x50\xf3\x8e\xd5\x5f\xde\xc2\xb3\xf7\x98\xe5\x5d\xfe\xe0\xce\xfd\x54\x54\xff\xc9\x6b\xfe\xb3\xe2\x5f\x44\xf1\x1b\xf4\xbe\xa6\xf6\x75\xad\x9b\x37\xf3\xd9\x4f\x06\x2b\x95\xa3\x99\x63\x27\x2b\x77\xb7\xe6\xbe\xbf\x03\x00\x00\xff\xff\x2f\x7f\x62\xa3\x67\x0c\x00\x00")

func debugCallLayoutGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_debugCallLayoutGohtml,
		"debug-call.layout.gohtml",
	)
}

func debugCallLayoutGohtml() (*asset, error) {
	bytes, err := debugCallLayoutGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "debug-call.layout.gohtml", size: 3175, mode: os.FileMode(420), modTime: time.Unix(1610873696, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _debugPageGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x4d\x6e\xdb\x3e\x10\xc5\xf7\x3a\xc5\x64\xd6\xa1\x85\x7f\x8c\x3f\xd0\x85\xe8\x4d\xb3\xe9\xb6\x41\x0e\x40\x4b\x63\x6b\x10\x7e\xa8\xe4\xc8\xa9\x21\x78\xd1\x93\xf4\x0a\x5d\x34\x40\x5b\x20\x3d\x83\x7a\xa3\x82\xfe\x88\x9d\x3a\x29\x50\x2f\x2c\x70\xde\xcc\xe3\x6f\x1e\xab\x8b\x26\xd4\xb2\xee\x08\x5a\x71\x76\x56\x54\xf9\x03\xd6\xf8\xa5\x46\xf2\x98\x0b\x64\x9a\x59\x01\x00\x50\x5d\x28\x05\xef\xe9\x43\xcf\x91\x1a\x70\x24\x06\xc4\x2c\x13\x28\xb5\xd7\xb7\xa5\xba\x35\x31\x91\x68\xec\x65\xa1\xde\xe0\xa9\xe4\x8d\x23\x8d\x2b\xa6\xfb\x2e\x44\x41\xa8\x83\x17\xf2\xa2\xf1\x9e\x1b\x69\x75\x43\x2b\xae\x49\x6d\x0f\x97\xc0\x9e\x85\x8d\x55\xa9\x36\x96\xf4\x7f\x97\x90\xda\xc8\xfe\x4e\x49\x50\x0b\x16\xed\x03\xce\x8a\x23\xd6\xdb\x9b\x1b\x08\xde\xae\x8f\x30\x96\xfd\x1d\xb4\x91\x16\x1a\xcb\x24\x46\xb8\x2e\xe7\x21\x48\x92\x68\xba\x89\x63\x3f\xa9\x53\x42\x88\x64\x35\x26\x59\x5b\x4a\x2d\x51\x66\x8a\x21\xa5\x10\x79\xc9\x5e\xa3\xf1\xc1\xaf\x5d\xe8\xd3\xd3\x5d\xc2\x62\x69\x76\x4d\xf3\x7e\x59\x95\xbb\x43\x51\x95\xbb\x8c\xaa\x79\x68\xd6\x87\xc6\x86\x57\x50\x5b\x93\x92\xc6\xbc\xa6\x61\x4f\x11\x9c\xa8\xe9\x21\x91\x45\x88\x0e\x4c\x2d\x1c\xbc\xc6\x7d\x71\x2b\xb0\xef\x7a\x01\x6e\x34\x26\x32\xb1\x6e\x29\xe2\xc1\x29\xcf\xa8\x6c\x17\x83\x05\x17\x55\x72\xea\x0a\x21\x3f\x9f\x46\xa1\x8f\x82\xfb\x88\x77\x83\xe7\xa6\xbb\xce\xd4\xcf\x1d\xcb\x93\xe9\x5c\x3c\xcc\xc5\xab\x2e\xb2\x33\x71\x9d\x21\xaf\x10\x56\xc6\xf6\xa4\x71\xfc\x3c\xfe\x1c\xbf\xfd\xfa\x34\xfe\x78\x8a\xa0\xcc\x14\xfb\x25\xca\x86\x57\xfb\x7a\xfe\x0d\x03\x2f\xc0\x13\x4c\x6e\xb6\x00\x80\xb8\xd9\x1c\xd5\x67\x1d\xb7\xb7\xef\xae\x5f\xd0\xff\x9a\xdd\xff\x27\x1b\x3d\x1b\x68\xe3\xcb\xc2\x30\x08\xb9\xce\x1a\x21\xc0\x26\xbf\x99\xaa\x8d\xb5\x08\x93\xcd\xe6\xfc\xd6\x3f\x76\xd9\xcd\x93\x4d\xf4\x4f\x8c\xd3\x57\x19\xa7\x87\xfe\x76\x8a\xb3\x61\xd8\x67\xb4\xd9\xc0\xf8\x38\x3e\xe4\xbf\x2f\xe3\xf7\xf1\xeb\xf8\x30\x3e\x56\x65\x3b\x3d\x77\x79\x05\xd0\x37\xa7\x7c\x2f\x13\x9f\xe6\xd0\x99\x25\xed\x12\x28\xce\x5c\x8a\x61\x68\x68\xc1\xfe\xd0\x96\x6b\x47\xb1\x28\x7e\x07\x00\x00\xff\xff\x94\xba\xd0\xfd\x2e\x04\x00\x00")

func debugPageGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_debugPageGohtml,
		"debug.page.gohtml",
	)
}

func debugPageGohtml() (*asset, error) {
	bytes, err := debugPageGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "debug.page.gohtml", size: 1070, mode: os.FileMode(420), modTime: time.Unix(1610873696, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _helpPageGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\xcf\x8a\xdb\x30\x10\xc6\xef\x79\x8a\x61\xee\xc9\xd2\xdd\x6b\x9c\x5b\xa1\xaf\x21\xdb\x93\x5d\x83\x2c\x19\x49\x49\x1b\x84\x61\xbb\xbd\x6c\xa1\xb0\xd0\x7b\xfb\x0c\x69\x68\xd8\x94\xed\x3a\xaf\x30\x7a\xa3\x22\x6f\xb2\x71\xfe\x14\xd2\xea\x60\x5b\x9a\x4f\xf2\xf7\x9b\x19\xdb\x7b\x70\x54\x56\x52\x38\x02\x4c\x85\x25\x84\x41\x5d\xf7\x7a\xde\xe7\x34\x2e\x14\x01\xba\xc2\x49\xc2\xba\x7e\x47\xb2\xf2\x9e\x54\xbe\x1f\x4e\x75\x3e\xc3\xb8\x04\x00\xe0\xfd\xee\xb0\x92\xd4\x64\x73\x58\x0c\x0d\xf3\x62\x0a\x99\x14\xd6\x26\x98\x69\xe5\x44\xa1\xc8\x40\xe9\xfa\x57\x38\x6a\x05\x87\x22\x21\xc9\x38\x68\xaf\x7d\x3b\xc9\x32\xb2\x16\xc1\x68\x49\x9b\x50\x67\x5b\x1c\xfc\x8d\xe7\xbc\xe6\x55\xb8\xe7\x55\xb8\xe3\x25\x84\x8f\xdc\x70\xc3\x3f\xc2\x67\x5e\xf2\x33\xaf\x78\x09\xbc\x02\x5e\xf0\x3c\xdc\x03\x3f\x73\xc3\xbf\x79\x19\x6e\x81\x7f\xf2\x53\x78\x88\xf2\x45\x78\xe0\x47\x5e\xed\xdc\x5c\xe4\xc5\x74\xd4\xdb\xcd\xc7\xda\x94\x20\x32\x57\x68\x95\x20\x42\x49\xee\x46\xe7\x09\x56\xda\x1e\xba\xe9\x82\x94\xe9\x1e\xe3\xab\xc4\xd1\x07\x27\x0c\x89\xad\x2e\x9e\xde\x8f\xa9\x31\x5a\x22\x28\x51\x52\x82\x25\x59\x2b\xae\x29\x82\xbf\xb7\x09\x5e\x21\x54\x52\x64\x74\xa3\x65\x4e\x26\x41\xfe\x1a\x69\x4e\xc3\x66\x1d\xc8\x78\x3f\x06\x1d\x0c\x06\x38\x1a\x5e\x6c\x7d\x1c\x10\xbc\xc0\xff\x0d\x2a\xd3\xb2\xff\xe6\xf2\x14\x56\xa1\xaa\x89\x03\x37\xab\x28\x41\x3b\x49\xcb\xc2\xe1\x76\x53\xea\x14\xa4\x4e\xf5\x2b\x53\x94\xc2\xcc\x62\xf5\x2f\x11\xa6\x42\x4e\x28\x41\xfe\x1e\xee\x78\x1d\x6e\x79\xce\x8b\x58\xc2\xf0\xe5\x30\xa7\xfb\x8e\x86\x17\x31\x61\x2f\xf3\x6e\xa5\xbc\x2f\xc6\xbb\xb6\x3b\xbb\xf5\x5e\xb7\x2a\xed\x60\xf0\xd6\x98\xba\x3e\x66\xfb\xff\xf6\xdc\x8e\x53\x5d\xd9\x74\xc0\x9f\xda\xe5\xe6\xf8\xd5\xc7\xe5\xf0\x9e\xa4\xa5\x7f\xb0\x99\x0b\x75\x4d\xe6\x1c\x97\xbc\xe6\xa6\xb5\x14\x9d\x7e\xe2\x5f\xed\x07\xc5\xcb\x68\x1a\x5a\xa3\x8f\x67\x1a\x6c\xff\x16\x07\x92\x5e\x37\xd6\xdb\x3c\xfc\x09\x00\x00\xff\xff\xc7\xb4\x97\x31\x88\x04\x00\x00")

func helpPageGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_helpPageGohtml,
		"help.page.gohtml",
	)
}

func helpPageGohtml() (*asset, error) {
	bytes, err := helpPageGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "help.page.gohtml", size: 1160, mode: os.FileMode(420), modTime: time.Unix(1613403086, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _homePageGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\x41\xaa\xc2\x30\x10\x86\xf7\x3d\xc5\x90\xfd\x7b\x5d\xbc\x6d\x5f\xd7\x5e\x23\x6d\x46\x08\x34\x69\xb1\xb1\x20\x43\x16\xad\x07\xf0\x2a\x15\x5b\x2c\x82\x9e\x61\x72\x23\x69\x14\xc4\xdd\x30\xff\xc7\xc7\x47\x04\x0e\x4d\x53\x49\x87\x20\x0a\xd9\xa2\x80\x5f\xef\x93\x84\x48\xe1\x56\x5b\x04\xe1\xb4\xab\x50\x78\xbf\xa9\x0d\x12\xa1\x55\xdf\x73\x51\xab\x83\x58\x5f\x00\x00\x44\x1f\x99\x41\xbb\x7f\xcb\xd6\x29\x53\xba\x83\xb2\x92\x6d\xfb\x2f\xca\xda\x3a\xa9\x2d\xee\xc0\xb8\x9f\x3f\x91\x47\x20\x42\x4d\xce\xe7\x70\xe4\x89\xe7\x30\x40\xe8\xc3\xc0\x63\x18\x78\x89\xd7\xc2\x37\x1e\x81\xaf\x3c\x42\xe8\x79\xe6\x0b\x3f\x78\xe2\x7b\x38\x65\x69\xf3\x52\x64\xa9\xd2\x5d\xbe\xd6\xc5\xcc\x67\x00\x00\x00\xff\xff\xf8\x6a\xe3\x5c\xdc\x00\x00\x00")

func homePageGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_homePageGohtml,
		"home.page.gohtml",
	)
}

func homePageGohtml() (*asset, error) {
	bytes, err := homePageGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "home.page.gohtml", size: 220, mode: os.FileMode(420), modTime: time.Unix(1613403104, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _menuLayoutGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x90\xcd\x4a\xc3\x40\x10\x80\xef\x7d\x8a\x61\xef\x6b\x40\xaf\x69\xde\x65\x92\x9d\x34\xab\x9b\x49\xd9\x9f\xa0\x84\x1c\x3c\x0a\x82\xaf\x52\xf4\x22\xfe\xd4\x57\x98\xbc\x91\x98\x56\xa8\xbd\x08\x85\x9e\x96\x1d\x3e\xbe\x6f\x98\x61\x30\x54\x5b\x26\x50\x2d\x71\x52\xe3\xb8\x00\x00\xc8\x8d\xed\xa1\x72\x18\xc2\x52\x55\x1d\x47\xb4\x4c\x1e\xda\xa8\x2f\x55\x31\x03\x33\x94\xdc\x2f\xc3\xd8\x03\x63\x5f\xa2\xd7\x06\xfd\x0d\x94\xab\xdd\xbb\x9f\xd1\xed\x1a\xd9\x68\xb7\x82\xeb\x14\xa2\xad\xef\xf4\x8f\x94\x38\xea\x8a\x38\x92\x3f\x90\xce\x62\x67\x0f\xc4\x25\x7a\x68\x83\xbe\x3a\x82\x66\x10\xff\x72\xba\xf4\xc8\x46\x41\xe3\xa9\x5e\xaa\x8b\xac\xe9\x5a\x52\x85\xbc\xc8\xbb\x6c\xe4\x59\x3e\x65\x33\x3d\xe5\x19\x1e\xd5\x32\x67\xcf\xd4\x0f\x84\xbe\x6a\x54\x21\x5f\xb2\x95\xd7\xe9\x5e\xde\x4e\xad\x63\x8a\xdd\x29\x07\x20\xb7\xde\xe7\x3f\x64\x3b\x3d\x4c\x8f\xff\x2c\x90\x67\xc9\xed\x7e\x79\x66\x6c\x5f\x2c\x86\x81\xd8\x8c\xe3\x77\x00\x00\x00\xff\xff\x69\x6b\xa2\x6d\x27\x02\x00\x00")

func menuLayoutGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_menuLayoutGohtml,
		"menu.layout.gohtml",
	)
}

func menuLayoutGohtml() (*asset, error) {
	bytes, err := menuLayoutGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "menu.layout.gohtml", size: 551, mode: os.FileMode(420), modTime: time.Unix(1613402978, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _searchPageGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x4d\x8e\xe3\x36\x13\xdd\xf7\x29\x0a\xdc\xf8\xfb\x80\x66\x89\x3f\x22\x25\x05\x56\x2f\x32\xeb\xac\x72\x02\x5a\xa2\x2d\x22\xfa\xf1\x48\xb4\xdd\x3d\x82\x17\x39\x49\xae\x90\x45\x06\x48\x02\x4c\xce\xa0\xdc\x28\xa0\xe4\xfe\x77\x07\x81\x61\xe3\xb1\xaa\xf8\xea\xb1\x4c\xbe\x71\x04\x6f\x9b\x7d\x6d\xbc\x05\xb2\x31\x83\x25\x80\xe7\xf3\xcd\xcd\x38\x96\x76\xeb\x5a\x0b\xc4\x3b\x5f\x5b\x72\x3e\xff\x68\x4d\x5f\x54\xe3\x68\xdb\xf2\x75\xc1\xa6\x2b\x1f\x48\x08\x01\x00\x8c\xe3\x33\x5d\x63\xdb\xc3\x85\x2e\xa4\xd6\xa5\x3b\x42\x51\x9b\x61\xc8\x49\xd1\xb5\xde\xb8\xd6\xf6\xd0\x78\x2a\xc9\xdd\x5c\x30\x17\x6d\xbb\xbe\x01\x53\x78\xd7\xb5\x39\x79\x91\x78\xcb\xe0\xda\xfd\xc1\xd3\x5d\xdf\x1d\xf6\xd0\x6c\x5e\x71\xcc\xa5\x73\x1e\x5c\x99\x93\x61\x16\x6e\x7b\xf2\xb8\x35\xb4\xa0\x41\x41\xdf\xd5\xd0\xf4\x74\x68\xa8\x20\xe0\x1f\xf6\x36\x27\xde\xde\x7b\x02\xad\x69\xec\xe3\xc6\x37\xc4\x33\xf9\xb0\x37\xed\x15\x21\x74\xd9\x1d\x9a\x6e\xcc\xe0\x0a\x6a\xca\xb2\x6b\x39\xb9\x5b\x1b\xa8\x7a\xbb\xcd\x09\x81\xae\x2d\x6a\x57\xfc\x94\x93\x93\x6b\xcb\xee\x84\xdd\xde\xb6\xff\x5b\x45\x83\x37\xde\x15\xd1\xd2\x12\x2b\x5b\xef\xb1\xf2\x4d\xbd\xba\x5d\x2d\x21\xd8\x9a\xcf\xab\xdb\xd5\xc9\x95\xbe\xca\xa5\x62\xb7\x95\x75\xbb\xca\xe7\x5a\xb1\xd5\xff\xaf\x48\x5c\x64\x1e\x77\x70\xdf\xd4\xed\x90\x93\xca\xfb\xfd\x77\x51\x74\x3a\x9d\xf0\x24\xb1\xeb\x77\x91\x60\x8c\x45\xc3\x71\x47\x60\x21\x25\x42\x10\xb8\xb0\xce\x78\xeb\xea\x3a\x27\xc5\xa1\xef\x6d\xeb\x3f\x75\x75\xf7\x3c\xc1\x8d\x83\x8d\xa3\x9f\x0f\x76\x08\xff\x13\x81\xa3\xb3\xa7\xef\xbb\xfb\x9c\x30\x60\xc0\x35\x70\xfd\x81\x24\x80\xf5\xde\xf8\x0a\xca\x9c\xfc\xa0\x50\x28\x05\x0a\x93\x54\x1b\x14\x32\x09\x5f\x60\xf3\x07\x45\xcc\x51\xc4\x49\x85\xa9\x50\x05\x72\x99\x2e\xc1\x94\x22\xe7\x12\x85\xd6\x14\x85\x42\x96\x51\xd4\x4a\xa3\x8a\x29\x47\x2e\x63\xe0\x28\x63\xb1\x60\xd4\xa9\x0e\x62\x50\xf2\x18\x65\x2c\x17\x04\x1c\xb9\x9e\xc9\xb4\x54\x14\x65\x12\x63\x26\x12\x8a\x99\x56\xa1\x20\xe1\x14\x75\x22\x31\x4e\x33\xca\x51\x30\x0d\x1c\x99\xa6\xcb\x26\x8e\x59\x9a\xd4\xc8\x98\x44\xc1\x13\x13\x04\x08\xf5\x24\x58\xa1\x88\x75\x85\x29\xe7\xef\x32\x41\xeb\x91\x22\x67\xaa\x60\x14\x13\x9e\xa2\x48\x24\x0d\x8d\x03\x3d\xa7\x1c\xe3\x54\xa3\x0e\xa7\x89\x75\x10\x2a\xe2\x98\x62\x96\x24\x17\x28\x90\x29\x0d\x8c\x72\x54\x3c\x54\x8b\x44\x53\x11\x46\x44\x45\x50\x7b\xc1\x1c\x85\x4e\x80\x85\x98\x52\xa8\x32\x2a\x30\x51\x20\x50\xa4\xfa\x4b\xc3\x51\xa9\x24\xcc\x5a\xcb\x82\x01\x2a\x29\x31\x16\xea\x49\x42\x00\x41\xc0\x3c\x30\x26\x52\x8a\x32\x8b\x1f\x61\x28\x62\x14\x95\x12\x14\x63\x41\x31\x0b\xd3\x66\x22\x9b\x11\xaa\x34\x9e\xa5\x31\x96\xa1\x4c\xd3\x0b\xca\xe2\x2f\x24\xfa\xe0\x56\x86\x4b\x77\x25\xb5\x8e\xcc\xdd\x3a\x0a\x2f\xeb\xcd\x4b\x8e\x4a\x77\xbc\xfa\xb8\x97\x17\x3b\x1c\x36\x8d\xf3\xcf\x57\xd3\xb7\xb0\xf1\x2d\xdd\xf7\xae\x31\xfd\x03\x81\xa3\xa9\x0f\x36\x27\xd3\x2f\xd3\x5f\xd3\xef\x7f\xff\x3c\xfd\xf9\xd2\x6d\xa2\xe0\x05\xcb\xfa\xd2\xe7\x19\x7f\xe8\x55\x17\xa3\x73\x5b\x68\x2d\xe0\xe2\x8b\x40\x9e\x2c\xf0\x29\x8b\x9f\x4c\x5d\x0f\x2f\xc3\xf0\x6f\x36\xa8\xae\x39\x4d\xd5\xbf\x0f\x8e\x63\x6f\xda\x9d\x7d\xe6\xbf\x36\xe6\x97\x46\x5c\x98\xba\x5e\x8c\xf8\x7d\xd5\x62\xe7\x57\x06\xfe\xe2\x2c\xb6\x1e\xec\x7f\x3e\xc6\x5b\x27\x5e\x8e\x21\x1f\x6b\x2b\x49\xee\xc6\xf1\x32\xb5\xf3\x19\xa6\x6f\xd3\xd7\xf0\xf3\xeb\xf4\xc7\xf4\xdb\xf4\x75\xfa\xb6\x8e\x2a\x79\xf5\x06\xdc\xbc\x15\x7d\xf3\x6a\xf1\x08\xfe\x09\x00\x00\xff\xff\x66\x38\x8b\x7a\xd2\x06\x00\x00")

func searchPageGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_searchPageGohtml,
		"search.page.gohtml",
	)
}

func searchPageGohtml() (*asset, error) {
	bytes, err := searchPageGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "search.page.gohtml", size: 1746, mode: os.FileMode(420), modTime: time.Unix(1613403104, 0)}
	a := &asset{bytes: bytes, info: info}
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
	"base.layout.gohtml":       baseLayoutGohtml,
	"call.layout.gohtml":       callLayoutGohtml,
	"call.page.gohtml":         callPageGohtml,
	"debug-call.layout.gohtml": debugCallLayoutGohtml,
	"debug.page.gohtml":        debugPageGohtml,
	"help.page.gohtml":         helpPageGohtml,
	"home.page.gohtml":         homePageGohtml,
	"menu.layout.gohtml":       menuLayoutGohtml,
	"search.page.gohtml":       searchPageGohtml,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
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
	"base.layout.gohtml":       &bintree{baseLayoutGohtml, map[string]*bintree{}},
	"call.layout.gohtml":       &bintree{callLayoutGohtml, map[string]*bintree{}},
	"call.page.gohtml":         &bintree{callPageGohtml, map[string]*bintree{}},
	"debug-call.layout.gohtml": &bintree{debugCallLayoutGohtml, map[string]*bintree{}},
	"debug.page.gohtml":        &bintree{debugPageGohtml, map[string]*bintree{}},
	"help.page.gohtml":         &bintree{helpPageGohtml, map[string]*bintree{}},
	"home.page.gohtml":         &bintree{homePageGohtml, map[string]*bintree{}},
	"menu.layout.gohtml":       &bintree{menuLayoutGohtml, map[string]*bintree{}},
	"search.page.gohtml":       &bintree{searchPageGohtml, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
