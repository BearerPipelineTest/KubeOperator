// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package i18n generated by go-bindata.// sources:
// locales/en-US/home.yml
// locales/zh-CN/home.yml
package i18n

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

var _localesEnUsHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x58\xcd\x72\xa3\x3a\x16\xde\xf3\x14\x0a\xae\xde\xf5\x74\xdd\x75\x76\x34\x26\x09\xd3\x36\x50\x80\x73\x27\xb3\xa1\x64\x38\xb6\x35\xc1\x12\x25\x89\x64\x7c\x77\xf3\x5e\xf3\x4e\xf3\x0a\x53\x47\x12\x18\x6c\xe7\x76\xba\xee\xaa\x93\xb4\xce\x8f\xce\xcf\xf7\x7d\x62\x51\x8b\xe3\x51\x70\x2f\x09\xd6\x51\x15\xfd\x23\x2e\xca\xe2\x9e\xf8\x09\x3d\x02\xa1\xad\x04\xda\x9c\x08\xfc\x9b\x29\xad\x7c\x2f\xce\xaa\x24\x2d\xcf\x87\xb2\x16\xa8\x02\xb2\x63\x6d\x4b\x18\x27\xfa\x00\x44\xc2\x9e\x29\x2d\x4f\x24\xce\x88\xb0\x7f\x52\x27\xa5\xe1\x48\x14\x68\xcd\xf8\x9e\x74\x74\x0f\xbe\xe7\x79\x8b\xba\xed\x95\x06\xe9\x85\xab\x4d\x51\x46\x79\xb5\x8c\x56\x51\x19\x55\x0f\x41\xbc\x8a\x96\xf7\xc4\xaf\x29\x27\x5c\x68\xd2\x40\x0b\x1a\x88\x3b\x8e\x81\xea\x5e\x4a\xe0\x9a\x28\x4d\x35\xf8\xa3\x83\xb8\x30\xe9\xe5\x9b\x24\x89\x93\xc7\x7b\xe2\x97\x87\x89\x99\x32\xce\x64\xcf\x39\xe3\xfb\x2b\xa3\x55\x1a\x06\xab\x7b\xe2\xc7\xc7\x4e\x48\x3d\x5a\xd5\x94\xa3\xd5\x16\x48\xdf\xed\x25\x6d\xa0\xf1\x31\x73\x09\x0d\x70\xcd\x68\xeb\xcd\x92\xae\xf2\xa8\x48\x37\x79\x18\xdd\x13\xff\x81\xb2\x16\x1a\xa2\x85\xcb\xff\x8e\x94\x07\x90\x80\x79\x50\x4e\xa8\x52\xa2\x66\x54\x43\x43\x0e\x42\x69\xd2\xf3\x06\x24\xd1\x07\xa6\xc8\x2b\x9c\xfc\x0f\xdc\x56\xff\x4c\x93\x5f\xf2\xfd\x87\xe0\x70\xc3\xf7\x43\xb0\x59\x95\x55\x98\x47\xcb\x28\x29\xe3\x60\x55\x85\x41\x62\xaa\x60\xc3\xde\x13\x7f\x09\x3b\xda\xb7\x9a\x9c\x6f\x3a\x29\x85\x0d\xda\xf8\x76\x64\xc2\xa7\x28\xfc\x71\xee\x9a\xa9\xf9\xd9\x8a\xe3\x1c\x9d\x4d\xcd\x3c\x98\xd1\x52\xe6\xe7\x5e\x81\x34\x67\x7c\x2f\x0b\x8a\xe2\xf7\x34\x5f\x8e\xc9\x24\x9b\x15\x76\xa4\xa3\x4a\xbd\x0b\xd9\x90\x61\x1e\xb6\x40\xb6\x2d\xe5\xaf\xff\xfb\xef\x7f\x7c\x2f\xcb\xe3\xe7\xa0\x8c\xaa\x1f\xd1\xcb\xa5\x21\x66\xd2\x49\xf6\x46\x35\xe0\xc5\x27\x59\x9c\xcd\xbd\x05\x96\xdf\x7b\x4a\x8b\xb2\x0a\x56\x79\x14\x2c\x5f\xce\xe3\xfd\x84\x9d\xb9\xdc\x01\xd7\x19\x63\x31\x5e\xfa\x66\x43\x6c\x67\xb1\x27\xce\xc5\xa4\x31\xef\x4c\x1f\x4c\x01\xdc\xa0\xdd\xf2\x5b\x7d\x7f\xa9\xb2\x3c\xfd\x7b\x14\x96\x7f\x29\x44\x27\xc5\xbf\xa0\xd6\xbe\x57\xbc\x14\x65\xb4\xae\xdc\x16\x3f\xa4\x9b\x64\x79\x7b\x89\x5b\x51\xd3\x16\x37\x78\xc7\xa4\xd2\xa6\x50\x49\x8a\x76\xc1\x73\x10\xaf\x82\xef\x2b\x1c\x91\x44\x90\xb8\x23\xf4\x8d\xb2\x96\x6e\x5b\xf0\xbd\xb8\xb0\x5b\x64\xee\x30\xd9\x5f\x66\x57\xca\x3a\xc5\x84\x7d\x5b\xef\x78\x9d\xa5\x79\x59\x45\x79\x9e\xe6\x43\xcf\x12\x41\x1a\xaa\x29\x5e\xd3\x99\xbd\x53\x45\x76\xa2\xe7\xcd\x1d\x71\x99\xd6\x07\xa8\x5f\x4d\x9e\xee\xc8\x8e\xb5\x70\x37\x77\x8a\xee\xaa\xe7\x60\xb5\xc1\x4c\xa3\x63\xa7\x4f\xd6\xaf\xe0\xa4\x65\x1c\xc8\x17\x35\x3f\xff\x7b\x9e\x26\x8f\xd5\x43\x9a\xaf\x03\x4c\x3d\xe6\xb5\x90\x12\x6a\x4d\x6c\x00\x21\x8f\x54\x7f\x68\x3c\x59\xa4\x69\x61\xc3\xc9\x16\x08\x6d\x2f\xf1\xa1\x0f\xd7\xe8\x79\x67\x6c\xe3\x3e\x61\xed\x06\x26\xd9\xac\xef\x89\x1f\x10\x2d\x34\x6d\x89\xd8\x91\x2f\x8a\x48\xf1\xae\xf0\x47\x73\x7d\x2a\x81\xd0\x2d\xc7\xeb\xb4\x5f\x89\x7a\x65\x9d\xf3\x33\x22\xcc\x30\x7a\x71\x32\x9f\xeb\x2d\xe3\x0e\xaa\x24\x28\xd1\xcb\x1a\x93\xf8\x4a\x24\x50\x25\xf8\xfd\x07\xf9\x14\xc1\xf3\x1c\xae\x14\x7d\x73\x33\x7b\x61\xec\x79\x0b\x84\x82\x33\x08\x60\x1d\xd6\x41\x19\x3e\x0d\x8b\x3c\xa0\x00\x53\x84\x0d\xdd\xf1\xbd\x34\x8f\x1f\xe3\xc4\x15\x7e\x7a\x5e\x48\xb6\x67\x9c\xb6\x1f\x19\x6e\x8a\x33\xf6\x07\x61\x19\x9b\x44\xcb\x01\x91\x1c\x59\x00\xc7\xc9\x9e\x4c\x9e\xe0\x9a\xd6\xda\xcc\x1e\x6d\x8e\x8c\x23\xd5\x51\x2d\xe4\x9d\x73\x38\xed\x5e\x22\x88\xea\xeb\x83\x71\x68\x56\x28\x58\xae\xe3\xe4\x1a\x6a\x31\x68\xe3\xe0\xd6\x38\xb5\x29\x5c\xc1\xed\xdd\x3c\xe9\x3c\x5a\x05\x65\xb4\x9c\x20\xc4\x06\xcd\x0e\x14\x53\x9f\xe2\x80\x5b\x7f\x93\xc2\x6a\x19\x64\x63\x06\x9b\x6c\x19\x8c\x19\xb4\x0d\xed\x2e\x03\x43\xc3\x6c\xdc\xe7\x28\x8f\x1f\x5e\xaa\x30\x5d\x4e\xe8\xf9\x19\x24\xdb\xb1\x9a\x6a\x26\x38\xa9\x45\x03\x04\xa4\x14\xd2\xf7\xa2\x75\x10\xaf\xaa\x65\x5c\x38\xa0\x58\x53\xd6\x0e\xec\xaf\xcc\x08\x36\x4c\x7d\xb2\xb0\x83\xb7\x69\x7b\xa3\x23\x3a\x3c\x52\x5d\x1f\xc8\xce\x8c\x96\x45\x28\x24\xa3\x71\x7e\x0a\xfc\x6d\xcc\x15\x4b\xf3\x27\x4c\x34\xcc\xc8\xa5\x13\x03\x4d\xf7\xc4\x7f\x97\x82\xef\xcf\x5c\x45\x84\x9c\x98\xd8\x04\x0d\x69\x8c\xc9\x5d\x92\x86\xb7\x40\x5d\x24\xf8\x80\xf2\x79\xf4\x18\xa7\xc9\x67\x55\x03\xb1\xc6\x3f\xc3\x79\x24\x7b\x0c\x85\xff\x0e\x81\x50\x30\x7c\x3a\x8c\x51\x0b\x3f\x23\x93\x96\xf2\xb9\x78\xb2\xc0\x1d\xda\xc2\xee\x41\x4f\x69\xed\x06\x66\xd7\x82\xef\xd8\xbe\x97\x66\x6e\x4c\xe3\xe2\x75\xf0\x18\x7d\xec\x8a\x1d\xe9\x1e\x3e\xe7\x28\xab\x8a\xa7\x34\xb7\x00\xae\xfa\xdd\x8e\xd5\x0c\x65\x62\xdc\x61\x59\x44\x07\x5c\x69\x5a\xbf\x7a\x8f\x51\x39\x74\x60\xe8\x70\x22\x86\x22\x1b\xa0\xc5\xf3\x6e\x6f\xd6\x70\xdc\x82\x1c\x57\x2f\x58\xe2\x3c\x7d\x51\x64\xdc\xb6\x2d\x00\x27\xb4\x31\xd2\x70\xba\xa0\x03\x0e\x7c\x51\x33\x4c\x31\xfe\x9d\xf6\x70\x21\x46\x45\x36\x90\xc0\x87\x72\xcc\x19\xdc\xd2\x62\x83\xed\x53\x50\x54\xae\x3d\x13\x0a\x99\xb4\x72\x6c\x4d\x78\x03\x61\xbc\x05\x17\x0d\x78\x09\x6e\xfa\xa0\x87\x9c\x9e\xae\xca\xa0\xf8\x81\xf4\xd2\x34\x04\x0f\xe1\x16\x38\x69\x6e\x7e\x1d\xa6\xc6\x29\xec\xaf\x9d\x6d\xd8\x3b\x65\x9a\x30\x4d\x1a\xc1\xe1\x1b\x06\xd8\xd2\xfa\xb5\xef\x82\xba\x16\x3d\xd7\x5e\x16\xe4\xc1\xba\x8a\xd6\x59\xf9\x72\xd9\xb6\x8e\x4a\x7a\x04\x0d\x52\xa1\xfc\x28\xab\x62\x93\x65\xb6\xbb\x1b\xae\xfa\x0e\x99\x19\x67\xf8\xd4\xe1\x13\x60\x2e\x42\x67\xd8\x64\x31\x62\x54\x58\xdf\x83\xf0\xc7\x26\xab\x82\x30\x4c\x37\xc9\xaf\x68\xad\x59\xe2\x9f\x16\x5d\xde\x02\x57\xe6\x42\xd0\x7f\x22\x1a\x5a\xfd\x42\x10\xd7\xd5\xef\x26\x47\xcf\xdd\xf1\x21\x5e\x45\xc5\x54\x12\xbb\x25\x72\x6d\xd3\xe3\xa5\x8c\x84\x22\x5b\xd8\x09\x09\x44\xbd\x33\x5d\x1f\xf0\xa9\x36\x39\x40\xed\xb5\x67\xab\x6f\xa3\x5c\xbf\xb3\xb6\x80\xc6\x68\x08\x0d\xe9\x3b\x33\xec\x13\xb3\x3c\x2a\xca\x34\x8f\xae\xed\x24\x28\x2d\x24\xe3\xfb\xf9\x7a\xe4\x4e\x70\x5c\xd7\x70\x72\xcd\x9f\x5e\xee\x2c\x65\x6f\x2b\xed\xf3\xce\x8c\xba\x7a\x28\xfd\x16\x5a\x81\xd4\xa5\xc5\x1c\xdb\x4a\x7c\x52\x89\x0e\xa4\xe3\xc0\x71\x9f\x3a\x90\x28\x19\xcd\x46\x59\x59\x74\x05\x09\x4f\x4e\xb8\x8f\x90\xe0\x7b\xa3\x06\xb3\x30\x13\xb9\x93\xc3\xf5\xcd\x30\x58\x98\xc9\x56\x41\x72\xc3\x67\xe6\x46\x66\xe2\xf3\x62\xd8\xaf\x6d\xbe\x5f\x0e\xf5\xc4\xd8\x5b\xa0\x2a\xb0\xa2\x61\x44\x36\x3b\x48\xc5\x89\xd7\x07\x29\x38\xfb\x03\x5b\xac\x40\x5a\x62\xff\xcd\x49\x8c\x55\xfa\x18\x27\x97\x36\x9b\xa9\xb2\x42\x62\xbc\x73\xa7\xcf\x52\xa1\x3c\x7f\x2e\xe8\xa4\x38\xb0\x2d\xd3\x8a\xe0\x19\x17\x63\x27\xc5\x91\xb4\x62\xbf\xc7\x01\x63\xfc\xdb\x67\x84\x99\xb7\xa8\x99\xf2\xc2\xb8\x30\x00\x76\x89\x6a\xf8\x4e\x61\x8a\x68\xaa\x5e\x2f\x11\x0c\x4d\xdf\x8e\xa1\xa1\x19\xef\x79\x5d\x85\x69\xf2\x10\x3f\x9e\x9f\x87\xe1\x94\x80\xae\x28\xff\x6c\x70\xf9\x5d\xa3\xbc\x24\xaf\x8f\x06\xad\x81\xae\x15\xa7\xa3\xc1\xc2\x96\xf2\x4f\x0e\x9c\xb7\x60\x1d\x52\xe1\x98\x27\xc6\x03\xae\x41\x42\x83\x6f\x3a\x05\x7b\xe3\x12\x53\x68\x59\xad\xd5\x19\x4f\x4c\xee\x58\xda\xf3\xb1\xaf\xa4\x9b\xbf\x0f\xe9\x9e\x32\x3e\x7e\x0a\x9a\x3e\x08\xe3\x0c\xdf\x1b\x98\x61\x5d\xf7\x1d\x83\x86\x50\xde\x4c\x92\x94\x60\x3c\x35\xc6\x38\x4e\x9e\x83\x55\x8c\xe5\x88\x3b\x2b\xcd\xdf\x68\xcb\x9a\x81\xc7\x27\x4f\x78\xde\x23\x01\xe3\x13\x66\x0f\x1c\xaf\x6e\xaf\x41\x9b\x46\x82\x52\x60\x22\xfe\xf6\xed\x5a\x1f\x28\x4d\xa5\xb9\x8c\x3b\x69\xb2\x51\xfd\x96\x03\x0e\x9e\x29\xd3\xdf\x3a\x21\x5a\x0c\x97\xa5\xe9\xea\x66\x9f\xe2\x8c\xe0\x99\x09\xd1\xdf\x80\xe1\xf1\x0d\x6c\x15\xd4\xfc\xd6\x23\x41\x5b\x0d\xa8\xb4\x3c\x79\x28\x3e\x8a\x32\x7f\xb9\xfe\xe4\x50\x4e\x3f\xa1\x89\x9d\xfd\x6a\x43\x65\x7d\x60\x1a\x6a\xdd\x4b\x30\x7c\x77\x43\x5f\x3a\x50\x1a\xb1\x72\xc0\xdb\x2c\x4f\x9f\xe3\x65\x94\x8f\x4a\x67\x8a\xb9\xb5\x04\x73\x13\x6c\x6b\xaf\xc5\x91\x6a\x56\x93\x23\x92\xb9\xcb\xff\x48\x79\x4f\xdb\xf6\x84\x7f\x64\xbb\xd3\xfc\x01\xa8\x26\x90\x55\xbe\x64\xd1\x2c\x84\x41\x2a\x9b\xac\x5b\xf9\x91\xb2\xef\x48\xca\xdb\xd3\xf0\xbb\x22\x88\x90\x5f\xc9\x1c\xa9\x7e\xce\xd6\x67\xbc\xde\xfd\x09\x5b\x13\x83\x03\xb0\xbf\xc5\x9f\xc8\x54\x33\x72\x3b\x73\x17\xad\xb5\xf9\x6e\x64\xf7\x49\x81\x52\x28\xde\x8b\xa8\x28\x50\x34\xfe\x88\x5e\x66\x30\xea\xfe\xdf\x7c\x64\x72\x73\xec\xa4\x60\x96\xa7\x48\x4e\x93\x69\x1f\xce\xda\xa7\x96\xe0\x6f\x20\xd5\x54\xa4\x18\x33\xd4\x70\x49\x3a\x15\xec\xfd\xe4\x6d\x37\xd4\x1f\x87\xf8\xff\x01\x00\x00\xff\xff\x66\x17\x1c\x29\xbd\x15\x00\x00")

func localesEnUsHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesEnUsHomeYml,
		"locales/en-US/home.yml",
	)
}

func localesEnUsHomeYml() (*asset, error) {
	bytes, err := localesEnUsHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/en-US/home.yml", size: 5565, mode: os.FileMode(420), modTime: time.Unix(1620292802, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesZhCnHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x58\x5d\x57\x1a\x49\x1a\xbe\xef\x5f\xc1\x81\x33\x77\xb3\x7b\xf6\x3a\x77\x1d\x68\x4d\x6f\x9a\xee\x3e\xdd\x8d\xb3\xee\x4d\x1f\xc7\xb0\xb3\x6e\x14\x3c\xa2\x7b\xce\xec\x55\x88\x1f\x80\x01\x61\x12\x75\xa3\x92\x28\x46\x23\xa3\x01\xfc\x98\x08\xa1\x41\xff\x4c\x57\x75\x73\xe5\x5f\x98\x53\x55\xdd\x45\x01\x31\xf1\x52\xfa\x7d\xde\x7a\xeb\xfd\x78\xde\xa7\x0c\x4d\x27\xe7\xe6\x92\x09\x4e\xe6\xa3\x82\x29\xfc\x43\xd4\x0d\xfd\x51\x20\x08\x4a\x05\xe7\xe4\x1c\x34\x2f\x41\xed\x2d\x28\x57\x83\x9c\xa8\x9a\xb2\x62\xf4\x0d\xdc\x46\x13\x94\xab\xce\x95\xe5\x58\xfb\x6e\xfd\xc6\xe9\xd6\x7b\x95\xcf\xbd\x77\x87\xa0\x72\x06\xd6\x76\x6c\xeb\x0d\x68\xbf\x11\xd5\x20\xc7\x85\xa6\x67\x97\x52\x8b\xf1\x05\x2e\x2c\xc5\x74\x43\xd0\xcc\x88\x20\x09\x86\x60\x8e\xf1\xa2\x24\x44\x1e\x05\x82\xf0\xff\x07\xf0\x6a\x0b\x64\x0f\x7a\x3b\x47\xa0\xfb\x06\xe4\x0a\xce\xfa\x35\x7c\x91\x76\x76\x57\x7a\x7b\x6b\xce\xcd\x51\x90\x42\x45\x1d\x07\xa1\xc5\x64\x59\x94\xc7\x1f\x05\x82\xc4\xc0\x6e\x15\x40\xb9\xea\xde\x96\xdc\x4a\xde\x6e\xd5\xee\x3a\xe9\x11\x88\xa4\x84\x79\x09\xdd\xab\xd1\x01\xab\xc7\x04\xe6\x1d\x5c\xc8\x38\xed\x13\x1c\xe8\x42\xfc\x59\x3c\xb1\x38\x33\x35\xcb\x0d\xc4\x68\x6a\x82\xae\xc4\xb4\xb0\x80\xf0\x24\xcc\xa3\x0b\xf7\x8f\xe3\xbb\x4e\xda\x6d\x1c\x3b\x27\x6f\x7b\xaf\x8f\xed\xd6\x2b\x58\xce\x81\xd5\x2b\x37\xbd\x69\xb7\x2c\x58\x6e\x07\xef\x71\x62\xfe\x53\x91\x1f\xea\x09\x14\x1b\xce\x66\x15\xe4\xb1\xb3\x31\x3e\x26\x19\x66\x58\x13\x22\x82\x6c\x88\xbc\x64\x86\x79\x19\xdf\x8d\x9c\x83\xb2\x61\xbd\x75\xeb\x47\x20\x53\x83\x85\xba\xdd\x2a\xb8\xcb\x5d\x72\x08\x4e\x08\xae\x6f\xf8\x89\x10\x7e\xda\x4f\x3d\x39\x91\xd4\x9a\x00\xec\xd6\x86\xb3\x59\x85\xd9\x26\xfa\x71\xaf\x05\x4a\xf9\x20\xa7\xf2\xba\xfe\x93\xa2\x45\xe8\x81\x72\x4c\x22\xb9\x5c\x73\x0e\xd2\x3e\xae\xed\xfc\xde\xc6\x07\xa9\x9a\x38\xc1\x1b\x82\xf9\x54\x98\x1c\x46\xf8\x37\x1c\x42\x70\x5c\xe8\xdf\xc9\xd4\x22\xf7\x44\xd1\x0d\x93\x97\x34\x81\x8f\x4c\xf6\x3b\x8d\xa4\x93\x69\x45\x2f\xaf\xd8\x9a\x5e\x65\x34\x9d\x14\xe7\x58\x45\x92\x4e\xbf\x9d\x46\x1d\x98\x8f\x27\x4d\x55\x53\xfe\x2e\x84\x8d\x87\xfa\xaa\x7c\x71\xf6\xea\x38\x7c\x7d\x52\x37\x84\xa8\xe9\x4d\xc8\x98\x12\x93\x23\xde\x80\xac\x66\xc9\x38\xc0\xf2\x27\x58\x6e\x8b\x2a\x29\x84\x82\x4c\xf9\x09\x5e\x94\xf8\xc7\x12\xaa\x9b\xa8\x06\xdc\xcf\x2b\xb0\x5d\x42\x99\xb9\xbe\x0a\x72\xa2\x4e\x1a\x16\x87\x88\x7d\x79\x11\xd8\xad\x0d\x32\x5a\x01\x51\x0d\x80\xb5\x4b\xe7\x34\x7d\xd7\xc9\xc3\xf2\xa9\xbb\xdc\x85\xd9\x12\x58\xdf\x87\x4d\x0b\xac\x1f\x04\x49\x2e\xc5\xa8\xaa\x68\x86\x29\x68\x9a\xa2\xf9\x35\x80\xe5\x53\x98\xbb\x01\xd9\x73\x50\x6c\x90\x69\x70\x76\x57\xe0\xd6\x39\x2c\xd4\xf1\x5d\x9b\xf0\xc3\x0b\xb8\x7f\x4c\x3e\xc1\xed\x8c\x6d\x5d\xe3\xb0\x59\x87\xc8\x95\x39\xc1\x4b\x31\x14\xfd\x0f\xa9\x80\x5b\xc9\xc3\x72\xce\xf9\xbd\x4d\xfc\x0c\x1a\xff\xa4\x29\xf2\xb8\x39\xa6\x68\x51\xde\xa0\xe6\xce\x59\x03\x14\x3f\xc0\x83\x0e\xe8\x14\xed\x56\x01\xd6\x3e\x38\x95\x21\x1c\xd3\xe9\x6c\x5e\xbd\xe3\x72\x37\x68\xea\xb3\xe7\xa0\xb1\xd6\x7b\x7d\x3c\x88\xf4\x2a\xf9\x2d\x18\x29\xdf\x20\xcc\x6b\x05\x39\x16\x45\x3d\xb0\x7a\x11\xf0\x30\xf8\x4e\x24\x48\xd0\x6a\xdd\x75\xf2\x6e\xf3\xca\xbd\xcd\x78\x60\x3a\xd7\x7e\x27\x89\xf8\x3c\x52\x2f\x52\x56\xe4\xc8\xbe\x7d\x87\xfa\xd7\xeb\xa9\x3c\xd8\xd8\x07\x7b\x07\x77\x9d\xdd\x1f\x52\x5f\x0d\x42\xe7\x27\x04\xea\xe5\x7b\x78\x8e\x0b\x2d\xa5\xe2\x0b\xfd\x41\x45\x17\x8f\xf2\x46\xf8\x09\x9d\xd2\xde\xe6\x8e\xdb\x68\x04\x39\x45\x13\xc7\x45\xd9\x4b\x29\x35\xd9\xd8\x1f\xb4\x8a\xe9\x7d\xe6\xe4\xc3\x86\x88\x63\x21\xbc\x00\xcb\xa7\xa0\x84\x58\x89\x34\x8b\x9b\xde\x44\x3b\xa0\x5e\x71\x4a\x6b\xe0\xb7\xb7\xb8\x53\x30\x9a\xcd\x3d\x22\xda\xda\x11\xc1\x63\x0b\x3e\x12\x15\xe5\xfb\xf8\x2b\x30\xf5\x6c\x6e\x26\x11\x20\xe6\x84\x2b\xdc\xc3\x33\x86\xc9\xd8\xe8\x34\x41\xe2\x0d\x21\xc2\x0c\xaf\x17\xe6\x65\x85\xb2\xa8\x5f\x6b\x29\xc2\xab\xf4\xd0\x98\x1a\xe1\xf1\xa1\xe8\xd7\x81\xc3\xec\xdb\x3a\xdc\xfc\x82\x4f\x9a\x10\x34\x71\x6c\xd2\x0c\x2b\x11\x66\x5f\xf5\x4e\xf3\x6e\x23\xcd\x64\x4b\x88\xf2\xa2\x64\x46\x44\xdd\x1b\xe7\xde\xcb\xba\x6d\x5d\x93\xa5\xe8\x1e\x9e\x39\x1f\xd3\xf7\xa5\xcb\xc7\xb2\xc5\x20\x68\x90\xff\xd2\x5b\x2d\x50\x0a\xf2\xe8\x9b\x16\x58\x47\x7f\xf5\x69\xdc\x67\x6c\xca\xe1\xa4\x9c\x3e\x81\x0f\x62\x31\x1b\xb0\x28\x98\xdd\x1e\x2c\x3f\x09\x0a\x13\x30\x09\xc8\xa9\x5f\x30\xfc\xcb\x71\xa1\x85\xf8\x2f\x33\xc9\x84\x4f\xa4\x9a\x30\x2e\x2a\xf2\x83\xb6\x25\xc8\xb7\xc1\xfe\x3e\x4b\xa4\xcc\x8e\xe3\x42\xff\x4b\x26\xe2\xbe\x57\xb4\x27\x1f\xe6\xd3\xf7\x30\xc0\xcf\xcb\x55\xa7\x7b\xe9\xd6\x2b\x20\xfb\x7a\x50\x0d\x10\x1a\x74\x37\x9a\xa0\xb8\xed\x91\x01\xde\x0b\x2c\xfb\xf5\x56\x0b\x4e\x97\xb0\xbb\x18\xe5\xc7\x85\xfb\x80\x5b\x65\xb0\x5c\xbc\x0f\xa8\x9a\xfa\x13\x45\x43\x29\x14\xe7\x03\x3e\xb5\x73\x5c\x28\x39\x1f\x4f\xa4\x16\xa7\xa6\x9f\x73\xe3\x82\xe1\x27\xcf\xaf\x4a\x9f\xd8\x70\xa6\x50\x52\xe6\x17\x92\xff\x89\x4f\x2f\x46\xe3\x73\x3f\xc7\x17\x68\xf7\xf3\x11\x8f\xd6\xbc\x3a\xe2\xbb\xfb\xec\xcf\x8e\x08\xc3\x80\x74\x84\xc9\x06\x20\xcb\xd7\xf7\x4f\x25\x86\x4f\x9f\xf7\xcc\x27\x19\xa7\x11\x7d\xe1\xa3\x9e\xf0\xba\xe9\xa5\x1b\x41\xb0\x31\xbb\x7e\xef\x3a\xe9\x11\x2c\x17\x4a\x24\x9f\xc5\x39\x19\xcd\x99\xbf\xff\x3d\x91\x67\x1a\xbc\xfe\x14\xf3\xf1\xb5\x6d\x6d\xbb\xeb\x2f\x9d\x97\x5f\xe0\xd6\x79\x2f\x53\x84\x6f\x0a\x76\xb7\x8c\x48\xb9\x5c\x85\xb9\x13\xb7\x92\xff\x31\xe0\x36\x9a\x4e\x2d\x07\x6e\x56\x41\x7d\xd9\xb6\x3e\x91\x9f\x41\x3d\x0f\x1b\x5b\x7f\x45\xc7\xfc\x3c\x35\xfd\x7c\x69\x9e\x9f\x9e\x4e\x2e\x25\x16\x39\x95\xd7\xf8\xa8\x29\x44\x55\x63\x12\x9d\x50\x7c\x09\xb7\xce\xfd\x3a\xa1\x8b\xeb\x31\x55\x25\x05\x44\xec\xbf\xd9\x80\x79\x24\x4a\x9d\x0b\x0b\xbc\x7f\x15\xe4\x86\xb4\x14\x3c\xa8\xf4\x4e\xf3\xcc\xc4\x7a\x6d\xfc\x98\x0f\x3f\x8d\xa9\x26\x1f\x0e\x2b\x31\xf9\xa1\x7a\x03\x1c\x65\x6c\xab\xeb\xfe\xf1\x11\x14\x9b\xf7\xa8\x0e\x2e\x34\x3f\x3b\x95\x18\x12\x99\xdf\x71\xcb\x0e\xc5\xa8\x5b\x46\xac\x3f\xc6\x79\xe2\xbc\xd8\xc7\x44\x49\xd0\x59\x1d\xe7\x69\x1b\xcf\x3f\x0a\x95\x08\x05\xb0\x56\x00\xd9\x0c\x2c\x1c\xb2\xf1\x0f\x4c\x1f\xf1\x48\x95\x3b\xa9\x1e\xb1\xfe\x8a\x72\xd7\x04\xdd\x50\x34\x61\xc8\x1c\xa6\x0f\xc1\x51\xc1\x37\xa7\xfd\xab\xc5\x53\xc9\xa5\x85\xe9\xf8\x68\x4a\x98\x6b\x7c\x23\x78\xb6\x68\x43\xd2\xb0\xdf\xcf\x03\x42\xf0\xe2\xbd\xdd\xde\x18\xea\x6a\xf7\x76\x0f\x69\x86\xda\x11\x69\x4f\x46\x3a\x8d\x4c\xa3\x2f\xe9\x0a\x3e\xb1\x52\x29\x41\xc6\x5b\xf0\x75\x0b\x16\x11\xa0\x79\xe9\x4f\xb7\x2a\xf1\xf2\x57\xfc\xb1\xd5\x65\xbc\x0e\x35\xe0\x28\x8e\x2d\x16\x83\xe3\x42\xb3\xcf\xa6\xe6\xc9\xd6\xa4\x7c\xe2\x29\xff\x52\x1e\xd6\x8e\x41\xf6\x1c\x8d\x83\xb7\x3a\xdb\x7f\xf3\x36\xac\xa4\x8c\x8b\xf2\x30\x82\x2e\x58\xe2\x1d\xa7\x05\x5b\xf7\xd7\x26\x79\x4b\x3a\x1f\xd3\xb0\x76\x88\x3e\x11\x88\xb3\x63\xf5\x76\xd6\x02\xf7\x48\x0d\x2e\x34\x3d\x93\xe2\xc2\xa2\x8e\x69\x62\x98\x3b\x50\xa4\xfe\xda\x82\xb9\x33\x58\x2c\xda\xad\x9a\xb3\xbb\x62\x5b\x16\x58\xaf\x20\xf4\x7f\xe7\xc2\xc9\xc4\xbf\x66\x7e\xe1\x26\xa2\x66\x58\x91\xc7\xc4\xf1\xfe\xbb\x83\x50\x39\xb3\xf7\xfa\x36\xc3\x4f\x58\x6a\xda\xef\x0a\xa6\x16\xdf\xec\x0d\x2e\x34\x33\x8f\x56\x45\xff\xe5\x8d\x5f\x0c\xce\xee\x8a\xa8\xc2\xfa\x67\xb4\xcb\x9b\x97\xb0\x9c\xc3\x7f\x51\xc5\xef\x36\x9a\xbd\x4c\x01\x6e\x9f\x13\x6b\xfa\x3c\x1f\x7a\x55\xa0\xe6\x39\x3c\x03\x85\x03\x2c\x40\xf2\x54\x4c\xf5\x32\xeb\x70\xf3\x06\xa3\x44\x79\x82\x97\xc4\x08\xb6\x07\xe5\x73\xf0\xee\x05\x95\xe3\xfe\x16\xa3\x25\xdc\x87\xd9\x12\x0e\x8c\x18\x12\x1a\x46\x95\x0f\xd0\xfd\xe7\x7e\x6e\x82\x93\x57\xe4\x3b\x78\x9d\x07\xb5\x92\xd3\xfd\x8d\x5e\xf4\x2f\xf3\xc9\xe4\x2c\x72\xa9\x2a\x8a\x34\x92\x45\x71\x3e\x00\x2f\x0e\xbe\xaa\x0d\xd0\x23\x87\xf9\x37\x41\xd0\x93\x20\xa9\xc5\x85\x5f\x39\xb4\x40\x75\x43\x9b\x1c\x7d\x39\xba\x8d\x63\xf8\xfe\x1a\xbe\xf7\x28\x1b\x57\x1e\x3d\x9e\xbc\x23\x68\x9f\x7b\xa4\x47\x09\xc4\x67\x20\x55\x53\x26\xc4\x88\xa0\xd1\xe5\xec\x66\x4e\xc1\x7a\x15\x56\x2b\xa0\x53\x04\xd9\x3d\x60\xb5\xe9\x3f\x29\x48\x6e\x61\xee\x15\x58\xaf\x12\x05\xc9\xbe\x00\x98\xd9\x36\x26\x55\x81\x3a\xa4\x7b\x85\x8c\x36\xdd\x2e\xa8\x65\xac\x55\xf2\x29\x80\xe8\xe3\xc7\xc1\x29\xfe\xfe\x76\xe9\x53\xd6\x77\xb6\x0b\x4e\xb4\xb3\xbb\x42\x7e\x74\x6a\xdb\xce\xd6\x31\xe2\x64\x2c\xf2\xfb\x24\x9b\x8a\xa7\x52\x48\xf2\xe9\x82\xae\x23\xbd\x82\x5e\xf7\x2c\x8f\x78\xdf\x03\xcf\xe3\xbf\x06\x50\xa9\xb6\xb2\x9e\x0a\x51\x35\x05\x11\x2f\xd3\x68\xbe\x29\x19\x70\xb7\xfb\x09\xe4\xb7\x49\x78\x1e\x04\x49\x08\x59\x61\x45\x1f\x2b\xea\xfd\x8c\x72\xdc\x9f\x01\x00\x00\xff\xff\xee\xb4\xe4\xec\xbb\x12\x00\x00")

func localesZhCnHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesZhCnHomeYml,
		"locales/zh-CN/home.yml",
	)
}

func localesZhCnHomeYml() (*asset, error) {
	bytes, err := localesZhCnHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/zh-CN/home.yml", size: 4795, mode: os.FileMode(420), modTime: time.Unix(1620354659, 0)}
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
	"locales/en-US/home.yml": localesEnUsHomeYml,
	"locales/zh-CN/home.yml": localesZhCnHomeYml,
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
	"locales": &bintree{nil, map[string]*bintree{
		"en-US": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesEnUsHomeYml, map[string]*bintree{}},
		}},
		"zh-CN": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesZhCnHomeYml, map[string]*bintree{}},
		}},
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
