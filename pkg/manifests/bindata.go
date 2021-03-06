// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/dns/cluster-role-binding.yaml (223B)
// assets/dns/cluster-role.yaml (397B)
// assets/dns/daemonset.yaml (6.316kB)
// assets/dns/metrics/cluster-role-binding.yaml (279B)
// assets/dns/metrics/cluster-role.yaml (246B)
// assets/dns/metrics/role-binding.yaml (293B)
// assets/dns/metrics/role.yaml (284B)
// assets/dns/namespace.yaml (369B)
// assets/dns/service-account.yaml (85B)
// assets/dns/service.yaml (520B)

package manifests

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

var _assetsDnsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\x31\x8e\x83\x40\x0c\x05\xd0\x7e\x4e\xe1\x0b\xc0\x6a\xbb\xd5\x74\x9b\xdc\x80\x48\xe9\xcd\x8c\x09\x0e\x60\xa3\xb1\x87\x22\xa7\x8f\x10\x4a\x45\x3a\x17\xfe\xff\xfd\x89\x25\x47\xb8\xce\xd5\x9c\x4a\xa7\x33\x5d\x58\x32\xcb\x23\xe0\xca\x77\x2a\xc6\x2a\x11\x4a\x8f\xa9\xc5\xea\xa3\x16\x7e\xa1\xb3\x4a\x3b\xfd\x59\xcb\xfa\xb3\xfd\x86\x85\x1c\x33\x3a\xc6\x00\x00\x20\xb8\x50\x04\x5d\x49\x6c\xe4\xc1\x9b\x2c\x16\xac\xf6\x4f\x4a\x6e\x31\x34\x70\x78\x37\x2a\x1b\x27\xfa\x4f\x49\xab\x78\xf8\xc4\xf6\xe7\xe3\xb6\x15\xd3\xa9\xa7\xe8\x4c\x1d\x0d\x3b\x74\x9a\x1d\xbe\xd3\xef\x00\x00\x00\xff\xff\xfa\x62\xe7\x50\xdf\x00\x00\x00")

func assetsDnsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleBindingYaml,
		"assets/dns/cluster-role-binding.yaml",
	)
}

func assetsDnsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role-binding.yaml", size: 223, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd9, 0xf6, 0x2a, 0x3b, 0x84, 0xd7, 0x3e, 0xc4, 0xe1, 0x70, 0x66, 0x31, 0xda, 0xc4, 0x2f, 0x53, 0x27, 0x29, 0x13, 0xfe, 0x80, 0x36, 0xc5, 0xa1, 0x70, 0xdc, 0x2d, 0xef, 0xcf, 0xe0, 0xc4, 0xeb}}
	return a, nil
}

var _assetsDnsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xb1\x4e\xc4\x30\x10\x44\x7b\x7f\x85\x75\xfd\x05\xd1\xa1\xb4\x14\xf4\x14\xf4\x1b\x67\x50\x96\xe4\x76\xad\xdd\x75\x4e\xe2\xeb\x51\x2e\x57\xa0\x8b\xa0\x9b\x19\xd9\xf3\x3c\x9e\x59\xc6\x3e\xbf\x2e\xcd\x03\xf6\xae\x0b\x12\x55\xfe\x80\x39\xab\xf4\xd9\x06\x2a\x1d\xb5\x98\xd4\xf8\x9b\x82\x55\xba\xf9\xc5\x3b\xd6\xa7\xf5\x39\x5d\x10\x34\x52\x50\x9f\x72\x16\xba\xa0\xcf\x5a\x21\x3e\xf1\x67\x9c\x47\xf1\x64\x6d\x81\xf7\xe9\x9c\xa9\xf2\x9b\x69\xab\xbe\x9d\x3c\xe7\xd3\x29\xe5\x6c\x70\x6d\x56\x70\xcf\x20\x63\x55\x96\xf0\x9b\x73\xd8\xca\x05\xbb\xa9\x3a\xee\x62\x63\x78\xa5\x3d\x5f\x61\xc3\xfd\xee\xc2\x1e\x37\x71\xa5\x28\x53\x3a\x02\xb7\x01\x90\xe0\xf2\x7b\xc1\xf1\x0d\xa1\x33\xc4\xb0\x32\xae\x0f\x84\x62\xa0\xc0\x1f\xcd\x8f\x5f\x73\x2c\xf6\x36\x7c\xa1\x04\x95\x02\xf7\xff\x00\x3f\x01\x00\x00\xff\xff\x76\x1b\x55\x2e\x8d\x01\x00\x00")

func assetsDnsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleYaml,
		"assets/dns/cluster-role.yaml",
	)
}

func assetsDnsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role.yaml", size: 397, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x84, 0xae, 0xd1, 0xba, 0xfa, 0x6b, 0xf8, 0x6e, 0x8d, 0x28, 0xc2, 0xa7, 0xaf, 0xc9, 0x3b, 0xc7, 0xcd, 0x80, 0xbe, 0xec, 0x98, 0xb4, 0x61, 0xa0, 0x9, 0xae, 0xa, 0xd8, 0xb2, 0x2e, 0x16, 0xf2}}
	return a, nil
}

var _assetsDnsDaemonsetYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x5f\x73\xdb\x36\x12\x7f\xf7\xa7\xd8\xd0\xbe\x38\x69\x4d\xff\x49\xe2\x34\xc7\xc4\xbd\xaa\xb6\x5c\x7b\x1a\xdb\x1a\x4b\x69\x1e\x3c\x1e\x0f\x04\xae\x44\x9c\x40\x00\x05\x40\xca\x1c\x5b\xdf\xfd\x06\xa0\x28\x91\xa2\xec\x36\xbd\xe9\x4c\xf3\xe0\xd8\xd8\xdd\x1f\x76\x17\xfb\x97\x13\x26\xe2\x08\x4e\x08\xa6\x52\xf4\xd1\x6e\x10\xc5\x7e\x43\x6d\x98\x14\x11\x10\xa5\xcc\x5e\x7e\xb0\xb1\x09\x82\xa4\xb8\xe3\x7f\x1a\x45\x28\x02\x11\x31\x70\x32\x44\x6e\x80\x68\x04\x83\x16\x88\x05\x9d\x09\xcb\x52\xdc\x30\x0a\x69\xb4\x01\x60\x31\x55\x9c\x58\x74\xbf\x03\x54\xa7\xfe\x77\xd4\x39\xa3\xd8\xa1\x54\x66\xc2\x5e\x92\x14\x23\x88\x85\x99\x53\x95\x66\x52\x33\x5b\x1c\x73\x62\x4c\x49\x34\x85\xb1\x98\x86\x42\xc6\x18\x52\xcd\x2c\xa3\x84\xcf\xb9\xa9\x14\x96\x30\x81\xda\x54\xe8\xa1\xd7\xb4\x8e\x08\xb0\x09\x2c\x25\x63\x04\x66\x56\xb5\xad\x38\x3c\xbd\x97\x71\xde\x93\x9c\xd1\x22\x82\xf3\xd1\xa5\xb4\x3d\x8d\x06\x85\x5d\x70\x59\xd4\x29\x13\xc4\x32\x29\x2e\xd0\x18\x27\x32\x67\x3f\x25\x9c\x0f\x09\x9d\x0c\xe4\x67\x39\x36\x57\xa2\xab\xb5\xd4\x0b\x39\x2a\xd3\x94\x38\x57\xdf\x40\x40\xa5\xc6\x58\x98\x00\x6e\x17\x64\xa2\xc7\xc6\xd3\x42\x2a\xc5\x28\xd8\x81\x60\x0f\x2d\xdd\x9b\x73\xee\x1d\x4b\x8d\x23\xc6\xb1\x2e\x92\x4b\x9e\xa5\x78\xe1\x1c\xb8\xb0\x7c\x69\xbb\x83\x61\xe3\xb0\x64\x5a\x50\x01\x52\xc7\xdf\x23\x36\x89\xa0\x7e\x43\x8d\x43\x23\x89\xaf\x04\x2f\x22\xb0\x3a\x5b\x8a\x2a\xa9\x9b\xf7\x2c\xfc\xde\x93\xda\x46\x70\xf8\xf6\xf0\x6d\x0d\xa5\xfd\x02\xee\x5d\xa5\x95\x54\xf2\x08\xbe\x9c\xf4\xbe\x1d\x29\xb4\x54\xad\x45\x1b\x1c\x2f\xd1\x9c\xf6\x4c\xa0\x31\x3d\x2d\x87\x18\xd5\xf8\x13\x6b\xd5\x2f\x68\xeb\x47\x00\xaa\xf4\x44\x82\x84\xdb\xa4\x49\xf1\xba\x7c\xd8\xff\xb0\xdf\x38\x36\x34\x41\xa7\xcf\xd9\x60\xd0\xab\x11\x98\x60\x96\x11\x7e\x82\x9c\x14\x7d\xa4\x52\xc4\x26\x82\x83\xba\xa8\x42\xcd\x64\xbc\x9e\x66\x32\x4a\xd1\x98\x41\xa2\xd1\x24\x92\xc7\x11\x1c\xd4\xa8\x23\xc2\x78\xa6\xb1\x46\xad\xbb\xc7\xc5\xb0\xcc\xec\x3a\x60\xce\x72\xfc\x87\xb8\xe2\xfd\xfe\x33\x2a\x1f\xfe\x1f\xae\x38\xac\xbd\xbc\x91\x99\xa6\x68\xa2\x46\x30\xff\x9e\xa1\xb1\xa6\x69\x2a\x55\x59\x04\x87\xfb\x69\xe3\x30\xc5\x54\xea\x22\x82\x1f\xf6\x2f\xd8\x4a\x21\x99\x64\x43\x0c\xf5\x90\xd0\x50\x69\x79\x5f\x7c\x43\x51\xf1\x79\x5d\x0b\xf5\x30\xe4\x72\x6c\xa5\xb1\x31\x6a\xdd\x38\x37\x48\x33\x8d\x21\x67\xc6\xa2\x08\x49\x1c\x6b\x34\xe6\x28\xfa\xf7\xc1\xe1\xbb\x06\x9f\xe5\x26\xa4\x4c\x25\xa8\x43\x93\x31\x8b\xe6\x68\xf0\xb9\x7f\xd7\x3d\x3e\x39\xeb\xde\x5d\xf7\x3b\x77\x5f\xcf\x07\x67\x77\x9d\x6e\xff\xee\xe0\xcd\x87\xbb\x5f\x8e\x2f\xee\xfa\x67\x9d\x37\x87\xef\x77\x96\x5c\xdd\xe3\x93\x3f\xe0\x6b\xe1\x1c\xff\x7c\xfc\xa7\x70\xd6\xf2\x3d\x83\xd6\xb0\x2c\x53\xc6\x6a\x24\xe9\x91\x0b\xcf\x68\x6f\xef\xe0\xcd\x0f\xbb\xfb\xbb\xfb\xbb\x07\xce\x09\x6f\xf7\xda\x5e\x40\x6d\x43\x57\x15\x8f\x7c\x25\xb3\xdc\xec\x29\xcd\x72\x62\xd1\xfd\xbe\x4b\xb5\x6d\x89\xcc\xe9\xe1\x04\x8b\x67\x24\x27\x58\xfc\xe9\xb2\xd7\x78\x9f\xaa\x58\xa5\x68\x35\xa3\xe6\x2f\x87\xe6\xc1\x13\xa1\xf9\x6e\x19\x9a\x4f\xd7\xff\xd5\x0a\x5f\xb3\xee\x29\x45\x9d\x6f\xfe\xa8\x03\xd4\x9a\x6a\xd9\x86\x9d\x51\x3c\x47\xfd\x8f\x69\xb1\x3e\x83\xdc\xd8\x20\x85\xc5\xfb\x46\x75\x73\xf6\x33\x8e\x63\x8c\x57\xba\xda\xf3\x4d\x34\x91\xc6\x1a\x1f\x28\xcf\x74\x50\xcf\x54\x73\x02\x8a\x1c\x2e\x3b\x17\xdd\x7e\xf7\xfa\xb7\xee\xb5\x1f\x95\x8e\x3f\x7f\xe9\x0f\xba\xd7\x77\x27\x57\x17\x9d\xf3\xcb\x75\x23\x53\x25\x8e\x22\x6f\xab\xe1\x90\xce\x8f\xbb\xfd\x9a\x12\x9b\x70\xec\x06\x0a\x90\x1a\xca\x89\xcc\xa0\x22\x9a\x58\x8c\xc1\x55\x10\x90\xa3\x6a\xc6\x32\x0d\xa9\xcb\xab\x41\x37\x82\x53\xa9\x41\xc8\xe9\x0e\xa0\x30\x99\x46\xb0\x09\x1a\xf4\x6a\x69\xe4\xc4\xb2\x1c\xcb\x59\xef\x23\x8c\xa4\x06\x24\x34\x69\x12\x76\x1a\x98\x44\x00\xe1\x8c\x18\x98\x32\x9b\x38\xac\x55\x7b\x4d\x36\x1a\xb1\x7b\x98\x32\xce\x81\x70\x23\x61\x88\x40\xe2\x18\xe3\xdd\x1a\x4e\x4e\x78\x86\x11\x04\x3e\x46\x42\x8d\x63\x66\xac\x2e\x76\xa5\x42\x61\x12\x36\xb2\xe1\x0a\xc1\xe4\x34\x68\x4d\x57\x35\xd7\xed\x0d\x99\xd8\x1b\x12\x93\xd4\x8b\x00\xad\xfd\xf1\x58\x37\xe2\x45\x9b\x1d\xfc\x1b\x85\x99\x04\xc5\x14\xba\xce\xb3\x51\xef\x61\x9a\x28\xd8\xfe\xaf\x1c\x1a\x08\x15\x3c\xc2\xbd\xab\xf4\x30\x71\x26\x3e\x3e\xfa\x18\xfb\x08\x53\xc2\xec\x47\xc0\x7b\x66\x61\x7f\x1b\x06\xdd\xeb\x8b\x3a\xc2\x55\xaf\x7b\xd9\x3f\x3b\x3f\x1d\xdc\x5d\x74\xae\x7f\xed\x5e\x1f\x05\x4b\x5b\xc7\x28\xd0\xbf\x66\x33\xd5\x82\x9a\xf8\xd9\x55\x7f\xd0\xbf\x3b\x3d\xff\xdc\x3d\x0a\x96\x71\x58\xe7\x18\x74\x2f\x7a\x2d\x86\x5d\x9b\xaa\xa0\xae\xc6\xf9\x69\xff\x68\x7b\x07\xb6\x7d\xd6\x43\xa8\x21\x24\x8b\xd0\x81\x4f\x9f\x3e\x41\xb0\xf5\x50\x05\xe0\xac\x21\xb9\x09\x17\x64\x82\x40\xfc\x9c\x2f\x35\xd1\x05\xb8\x54\x59\x86\x81\xe4\x71\x99\x42\xfe\x7c\xdb\x00\xb1\x56\xb3\x61\x66\xd1\xd4\x5f\x9e\x2a\x08\x47\x10\x86\x4b\x6a\x28\x05\x2f\xdc\xc5\x4b\x23\x67\x81\xfb\x7b\x61\x52\x53\x93\x69\xe2\xee\x2d\x9d\x1e\xcb\x46\xe9\x8c\x91\x72\x17\xd8\x61\x07\x4c\x4e\xef\x98\x32\x0d\xb2\x8b\x6f\x93\x53\x60\xc2\xc1\x57\x76\xdf\xfc\x74\x3b\x0b\x5a\x50\xce\xe2\x53\xb4\x34\xa9\xfc\x03\xe7\x3d\x18\x69\x99\x02\xe5\x99\xb1\xa8\x5d\x6d\x04\x36\x02\x55\x16\xb4\x5d\xf8\x8a\x90\x3a\x17\x19\xcc\x51\x13\x0e\x56\x33\x34\x2d\x4c\x2b\x21\x96\xc0\x6c\x04\xe7\xbd\xfc\xdd\x8e\xfb\xf9\xde\xff\x7c\x07\x32\x47\xed\xc6\x5b\x5f\x45\xdc\xf9\xe2\x64\x17\x06\x09\x82\x9d\x4a\xe0\xc4\xe5\xbb\x58\x03\xec\xec\x76\x06\xc6\xa8\xb8\x2c\x52\x14\x76\x9e\xa3\xbf\x66\xba\xd0\x20\x85\x7b\x21\xd4\x70\xa5\x50\xf4\x2d\xa1\x13\x78\x75\xd5\xef\x1d\xbc\x7d\x0d\x21\xd8\x44\x1a\x74\x7a\x09\x69\x5b\xc0\x26\x53\xae\x2f\xba\x31\x1e\xb8\x24\xf1\x90\x70\x22\x28\x6a\xe3\xf5\x74\x8d\x8d\xf9\x5a\x42\x68\xc2\xc4\x18\x4e\x2e\xfb\x60\x13\x2d\xb3\x71\xe2\x55\x5f\xc1\xa3\x69\x6c\x8e\x5e\x6d\xc7\x6c\x0c\xa1\x85\x0e\xfc\x14\x6c\x3d\x2c\x0b\xe8\x2c\x80\xef\x4d\xe2\x6e\x73\x0f\x94\xd3\xd9\xee\xd6\x43\xb3\xbe\xcc\x82\xed\x15\xc4\xf2\xdf\x02\xb1\xd3\xf9\x1b\x40\xe1\x7b\x4b\xd5\xdf\xa3\xeb\x5f\x45\x7e\xbd\x02\xed\xde\x9e\xb9\xd0\xde\x7a\x78\xe1\x9c\x7c\xf3\xdd\xed\x6c\x85\xa5\x15\xe2\x00\x4c\x99\xa3\x57\x5b\xaf\x30\x27\xdc\x5d\xe6\x05\xd9\xed\x2c\x78\xbd\x0a\x0f\x2e\xd6\x6f\x6e\x20\xd8\xfa\x4f\x00\x21\xfe\x0e\xfb\xf0\xf2\xa5\x13\xd9\x64\xaa\x4c\x21\x08\x05\xc2\x3e\xdc\xde\x7e\x74\xf5\x40\xac\xb1\x7c\x9e\x93\x37\x73\xab\x82\xdb\xa3\x60\xeb\xa1\x12\x5f\xc3\x3f\xd4\x48\x26\xad\xf3\x11\x6b\x99\x25\x70\xa3\x75\xd0\x38\xd9\x84\x2f\x2a\x26\x16\x6b\x4d\x1c\x7c\xd9\x61\x23\x98\x22\x8c\xd1\xba\x96\xc4\xe2\x5a\xb2\x9b\x15\x80\xaf\x58\xf6\x34\x21\x2d\x64\x2d\xb0\x69\x82\xc2\x99\xad\xfd\x44\x34\xdf\xb3\x17\x68\x32\xb3\x6e\x56\x92\x1a\x88\x62\x90\x09\x92\x13\xc6\xc9\x90\x71\x66\x8b\x95\x6b\xfa\x96\x70\x04\x14\xbe\x7a\x00\x95\x19\x8f\x5d\x53\x31\xd6\x3d\x6d\xed\x42\x36\xf2\x55\xb7\xba\x81\x19\x88\x91\xa3\xc5\x78\xa3\xfd\x66\xa1\x98\x07\x92\xf7\xfe\x77\xb7\xe1\x2c\x78\xea\x99\x36\xe1\xe7\x8c\xf1\x18\x08\x08\x9c\xd6\xea\x79\x59\xfa\xea\x06\xbb\xd2\x22\x33\x0d\x34\x33\x56\xa6\x0b\x8d\x47\x8c\x5b\xd4\x18\x3b\x9b\x57\xb0\xc7\x1a\x15\x84\x39\x04\x9b\xb0\xf5\xb0\xda\x10\xcb\x92\xdf\x68\x01\x3f\x3e\xd3\x04\x4a\x5d\x3b\x4a\xa1\xaf\x41\x65\xc7\x5c\x2a\xe1\x0a\x7d\x7b\x22\x82\x56\x0f\x78\x51\x39\xe5\x89\x1e\x30\x4f\x2b\x55\xe6\x55\xc5\x5c\x86\xef\xed\x6c\xad\x00\x00\xd2\x44\x82\x8f\xec\x59\x29\x54\xfd\xd7\x4e\x63\x78\xc2\x15\x3f\xb6\x6c\x5f\xbd\xa4\x15\xf4\xeb\xc2\xde\xf9\x68\x70\x75\x72\x15\xad\x09\x7f\x62\x65\xca\x28\xe1\xbc\x70\x3d\x89\xe4\x92\xc5\x40\x44\x01\x4c\x50\x29\x8c\x5f\x4c\x2d\x0c\x31\x21\x39\xab\x8d\xdd\x15\xea\x35\x2a\xee\x26\xd1\x75\x11\x91\xca\x98\x8d\x18\xc6\x90\x97\xdf\x16\x5d\x14\x0a\xc4\x78\x25\x36\x5d\x2f\x50\x2b\x66\xb6\x62\xe0\xf1\x71\x3e\x31\x3c\xcf\xd7\xb6\xba\xe2\x75\x99\xe1\x52\x56\x63\x2a\x73\x8c\x97\xb6\xfa\xa8\xa6\x1a\xdd\x1e\x58\xa6\x8e\xef\x67\xcb\xb9\x04\xa8\x54\x05\xd0\x24\xd3\xcd\x24\x59\xa9\x3f\x86\x23\x2a\x78\xbf\x0f\x2f\xfd\x08\xd8\xa0\x65\xc2\x4d\x95\xed\x51\xa4\xf1\x78\xdf\xfc\x29\xa3\x5a\x17\x63\x61\xaa\x5d\xe9\x04\x47\x24\xe3\xd5\xe5\x6e\x8c\xec\x23\x47\x6a\xa5\x5e\x02\x4c\xb2\x21\x6a\x81\x6e\x1e\x63\x72\x4f\x9a\x08\x38\x13\xd9\x7d\x49\x9c\x73\x95\x85\xad\x6f\xdd\x38\x3a\x2e\x96\xb2\xb6\x50\x18\xc1\xb5\xe4\x9c\x89\x71\x19\x4b\x4b\xfd\xeb\xa7\x51\x73\x01\x91\x16\x23\x3f\xbb\xc4\xfe\x9b\xb3\xdf\xa7\x1d\x3f\x6a\xd0\x32\x13\xb1\xf1\x45\x4c\xa1\xa6\x28\x7c\x89\xcc\x54\x03\xe0\x55\x26\x38\x9b\xf8\x9d\xa5\x36\xd8\xd4\x60\x76\xdc\x38\xe8\x36\x96\x12\x2d\x96\x53\xf1\xba\x3e\x72\xa4\xe4\xfe\x4b\x55\x6e\x39\xba\x65\xfb\x5f\x73\x6a\xb9\x0c\xb6\x3e\x23\xaf\xff\x94\x5a\x9e\x5e\x10\x15\xd5\x76\xbf\x4b\x92\x3e\xb7\xff\x02\x30\x8b\x69\xe3\x05\x43\x98\x60\x11\x41\xf5\x81\x77\xcd\x17\xb9\x15\xd2\x33\xbb\xa9\x3b\xf2\x8b\xe9\xc6\x2a\xc6\x9a\x45\xb5\x7a\xc1\xd3\x36\xf4\xba\xaf\x02\x9b\x6e\xbd\xd6\x68\x9f\xb5\xd0\x4a\xee\xd6\x16\x26\xc5\xc2\xc6\x4d\x3f\xfd\xb9\x5c\x37\x2e\xd1\x74\x26\xc0\x0d\xc3\xc5\xd4\x35\xc6\x5d\x18\x94\x12\x08\x84\x73\xb0\x84\x89\x85\x86\x21\x48\xe5\x48\x52\x47\xd0\x75\xdd\xce\x6c\xfc\x2f\x00\x00\xff\xff\x07\x86\xa1\x36\xac\x18\x00\x00")

func assetsDnsDaemonsetYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsDaemonsetYaml,
		"assets/dns/daemonset.yaml",
	)
}

func assetsDnsDaemonsetYaml() (*asset, error) {
	bytes, err := assetsDnsDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/daemonset.yaml", size: 6316, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x22, 0xce, 0x92, 0x1, 0xb9, 0xc5, 0xd4, 0x2d, 0x24, 0x96, 0xc2, 0x40, 0xe9, 0x9, 0x7b, 0x12, 0x47, 0xce, 0x34, 0xbb, 0x26, 0x82, 0xbd, 0xd9, 0xc, 0x2d, 0x7c, 0x99, 0xe1, 0x12, 0xa9, 0x62}}
	return a, nil
}

var _assetsDnsMetricsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xb1\x4a\x04\x41\x0c\x86\xfb\x79\x8a\xbc\xc0\xae\xd8\x1d\xd3\xa9\x85\xfd\x09\xf6\xb9\x99\x9c\x1b\x77\x27\x19\x92\xcc\x16\x3e\xbd\x2c\x8a\x08\xe2\xb5\x81\x7c\xdf\xff\xad\x2c\x35\xc3\xd3\x36\x3c\xc8\xce\xba\xd1\x23\x4b\x65\x79\x4b\xd8\xf9\x95\xcc\x59\x25\x83\x5d\xb0\xcc\x38\x62\x51\xe3\x0f\x0c\x56\x99\xd7\x93\xcf\xac\x77\xfb\x7d\x6a\x14\x58\x31\x30\x27\x00\xc1\x46\x19\xaa\xf8\xd4\x54\x38\xd4\x0e\x92\x8f\xcb\x3b\x95\xf0\x9c\x26\xf8\xd2\xbd\x90\xed\x5c\xe8\xa1\x14\x1d\x12\x3f\x7f\xdd\xb4\x51\x2c\x34\x7c\x5a\x4f\xfe\x7d\xf6\x8e\x85\x32\x68\x27\xf1\x85\xaf\xf1\x9b\x6c\xba\xd1\x99\xae\x87\xf9\x4f\xc7\x7f\x6b\x00\xb0\xf3\xb3\xe9\xe8\x37\xba\xd2\x67\x00\x00\x00\xff\xff\x5b\x52\x00\xaa\x17\x01\x00\x00")

func assetsDnsMetricsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsClusterRoleBindingYaml,
		"assets/dns/metrics/cluster-role-binding.yaml",
	)
}

func assetsDnsMetricsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/cluster-role-binding.yaml", size: 279, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x79, 0x95, 0x6f, 0xa4, 0xd5, 0xed, 0x48, 0x27, 0x41, 0x56, 0x5c, 0xea, 0x5c, 0x89, 0xdc, 0xc1, 0x44, 0x91, 0xd4, 0xb, 0x18, 0x85, 0x79, 0x75, 0xaa, 0x6e, 0xb5, 0x98, 0xbe, 0xc6, 0x33, 0x43}}
	return a, nil
}

var _assetsDnsMetricsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcd\x31\x4b\x34\x41\x0c\x87\xf1\x7e\x3e\x45\xe0\xad\x77\x5f\xec\x64\x5a\x05\x3b\x0b\x05\xfb\xec\xce\xdf\xdb\x70\x3b\xc9\x90\x64\x0e\xf4\xd3\x8b\x70\xb6\x0f\x3f\x78\xfe\xd1\xd3\x39\x23\xe1\xe4\x76\x22\x48\x81\x86\x46\xdb\x17\x0d\xb7\x8e\x3c\x30\x83\xd2\x28\x76\xe7\x01\x7a\x7e\x7d\xa7\x8e\x74\xd9\x83\xa0\x6d\x98\x68\x16\x1e\xf2\x01\x0f\x31\xad\xe4\x1b\xef\x2b\xcf\x3c\xcc\xe5\x9b\x53\x4c\xd7\xeb\x63\xac\x62\xff\x6f\x0f\xe5\x2a\xda\xea\xdf\xf0\xcd\x4e\x94\x8e\xe4\xc6\xc9\xb5\x10\x29\x77\x54\x6a\x1a\x4b\x37\x95\x34\x17\xbd\x14\x9f\x27\xa2\x96\x85\x78\xc8\x8b\xdb\x1c\xf1\x4b\x17\xb2\x01\xe7\x34\x5f\x6d\x40\xe3\x90\xcf\x5c\xc5\x0a\x91\x23\x6c\xfa\x8e\x3b\x6b\x1a\x88\x42\x74\x83\x6f\xf7\x74\x41\x96\x9f\x00\x00\x00\xff\xff\x9f\xa8\x4d\x6c\xf6\x00\x00\x00")

func assetsDnsMetricsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsClusterRoleYaml,
		"assets/dns/metrics/cluster-role.yaml",
	)
}

func assetsDnsMetricsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/cluster-role.yaml", size: 246, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x64, 0xdb, 0xe0, 0x95, 0x65, 0xae, 0x53, 0x96, 0x3a, 0x5f, 0x5e, 0x8b, 0x69, 0xe2, 0x7d, 0x5, 0xbf, 0x1f, 0x3a, 0xf, 0xff, 0xd0, 0x6b, 0x23, 0x4f, 0xfd, 0x11, 0x7f, 0x57, 0xd4, 0x4a, 0x8b}}
	return a, nil
}

var _assetsDnsMetricsRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\xb1\x4e\xc4\x40\x0c\x04\xd0\x7e\xbf\xc2\x3f\x90\x20\xba\xd3\x76\xd0\xd0\x1f\x12\xbd\x6f\xd7\x97\x98\x64\xed\x95\xed\x4d\xc1\xd7\x23\xa4\x48\x54\x20\x5d\x3b\x9a\xd1\x1b\xec\xfc\x41\xe6\xac\x92\xc1\x6e\x58\x66\x1c\xb1\xaa\xf1\x17\x06\xab\xcc\xdb\xc5\x67\xd6\xa7\xe3\x39\x6d\x2c\x35\xc3\x55\x77\x7a\x65\xa9\x2c\x4b\x6a\x14\x58\x31\x30\x27\x00\xc1\x46\x19\xba\x69\xa3\x58\x69\xf8\xb4\x5d\xfc\x8c\xbd\x63\xa1\x0c\xda\x49\x7c\xe5\x7b\x4c\x55\x3c\x99\xee\x74\xa5\xfb\xcf\x14\x3b\xbf\x99\x8e\xfe\x8f\x9f\x00\x7e\xf9\xbf\x34\x1f\xb7\x4f\x2a\xe1\x39\x4d\x67\xfb\x9d\xec\xe0\x42\x2f\xa5\xe8\x90\x78\xf0\x65\x53\xe1\x50\x63\x59\x20\x7d\x07\x00\x00\xff\xff\xb9\xd9\xab\x8d\x25\x01\x00\x00")

func assetsDnsMetricsRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsRoleBindingYaml,
		"assets/dns/metrics/role-binding.yaml",
	)
}

func assetsDnsMetricsRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/role-binding.yaml", size: 293, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc, 0x7d, 0xc7, 0x45, 0x33, 0xc4, 0xd8, 0xf, 0x8d, 0x89, 0x8d, 0x6, 0x47, 0xa7, 0xa, 0x6b, 0x17, 0xf5, 0x5f, 0x5a, 0x2f, 0xd8, 0xf9, 0x6, 0x71, 0xaa, 0x78, 0x8d, 0xb5, 0x7a, 0xf6, 0x99}}
	return a, nil
}

var _assetsDnsMetricsRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xb1\x4e\xec\x40\x0c\x45\xfb\xf9\x0a\x6b\x5f\x9d\x7d\xa2\x5b\x4d\x8d\x44\x47\x01\x12\xbd\x77\xe6\x42\xac\x24\xe3\x91\xed\x04\xc1\xd7\xa3\xcd\x06\x89\xca\xf7\x1e\x59\x3e\xfe\x47\x2f\x3a\xc3\xa9\x01\x15\x95\xae\x5f\xd4\x4d\x17\xc4\x88\xd5\x29\x94\xbc\x18\x77\xd0\xe3\xf3\x2b\x2d\x08\x93\xe2\x84\x56\xbb\x4a\x8b\xc4\x5d\xde\x60\x2e\xda\x32\xd9\x95\xcb\x99\xd7\x18\xd5\xe4\x9b\x43\xb4\x9d\xa7\x8b\x9f\x45\xff\x6f\x0f\x69\x92\x56\xf3\x2e\x4a\x0b\x82\x2b\x07\xe7\x44\xd4\x78\x41\xfe\xe3\x1b\xa6\x8b\x1f\xd8\x3b\x17\x64\xd2\x8e\xe6\xa3\xbc\xc7\x50\x9b\x27\x5b\x67\x78\x4e\x03\x71\x97\x27\xd3\xb5\xfb\xed\xca\x40\xa7\x53\x22\x32\xb8\xae\x56\x70\x30\x87\x6d\x52\xe0\x7b\xf9\xfd\xf8\xde\xba\xd6\x5b\xd8\x60\xd7\x63\xf9\x03\xb1\xcf\x59\xfc\x1e\x3e\x39\xca\x98\x7e\x02\x00\x00\xff\xff\x29\x39\xda\x05\x1c\x01\x00\x00")

func assetsDnsMetricsRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsRoleYaml,
		"assets/dns/metrics/role.yaml",
	)
}

func assetsDnsMetricsRoleYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/role.yaml", size: 284, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8c, 0xf2, 0x4e, 0x40, 0x91, 0xd8, 0x5e, 0x1c, 0x98, 0xb6, 0x2f, 0x11, 0x2a, 0x15, 0x8f, 0xe4, 0x7c, 0xfe, 0xc6, 0x31, 0xf3, 0xb2, 0xa0, 0x38, 0xb2, 0x3f, 0x15, 0x5a, 0x33, 0x12, 0xd2, 0x88}}
	return a, nil
}

var _assetsDnsNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xcd\x4e\xc4\x30\x0c\x84\xef\x79\x8a\x51\x38\x2f\x3f\xd7\xbc\x03\x5c\x90\xb8\xbb\x8d\x97\x35\x4d\xed\x2a\x76\xcb\xeb\xa3\xb2\x15\xac\xb4\xc7\x68\x46\xf3\x7d\xf1\x24\x5a\x0b\xde\x68\x66\x5f\x68\xe4\x44\x8b\x7c\x70\x77\x31\x2d\xd8\x5e\xd2\xcc\x41\x95\x82\x4a\x02\x48\xd5\x82\x42\x4c\x7d\x7f\x02\xb6\xb0\xfa\x45\xce\xf1\x28\xf6\xa4\x56\xf9\xe4\xdc\x78\x0c\xeb\x05\x39\x27\x40\x69\xe6\xf2\x5f\x3b\x55\xf5\x04\x34\x1a\xb8\x1d\x13\x0f\x70\x0e\x6c\xd4\x56\x46\x18\x68\x33\xa9\xa8\xbc\xb0\x56\xd1\x4f\x98\x62\x5a\x07\x06\xd5\x59\x7c\x97\x42\x5c\x28\x8e\x82\xef\xf1\xdf\x38\x68\x11\xbf\xd7\xea\xab\x9e\x1a\x6f\xdc\x0a\xf2\x73\x3e\x98\xd4\x9a\x7d\xdf\x78\xcd\xa6\x12\xd6\x77\x62\x18\x9a\xd9\x84\xb3\x75\xbc\x73\xdf\x64\xe4\xd7\x6b\x0a\x1b\xbe\x78\x0c\x87\xec\x16\xe2\xbf\xbf\xbb\x1e\xed\x8e\x3a\xb6\xd5\x83\xfb\xcd\x70\x41\x8e\xbe\x72\x4e\x3f\x01\x00\x00\xff\xff\x82\x6d\x29\x03\x71\x01\x00\x00")

func assetsDnsNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsNamespaceYaml,
		"assets/dns/namespace.yaml",
	)
}

func assetsDnsNamespaceYaml() (*asset, error) {
	bytes, err := assetsDnsNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/namespace.yaml", size: 369, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe, 0xab, 0x50, 0x84, 0x61, 0x5f, 0x41, 0xf4, 0x17, 0x3b, 0x6, 0x84, 0xc0, 0x5f, 0x4f, 0xbb, 0xd8, 0x1d, 0xae, 0x26, 0x3e, 0x1f, 0x29, 0x2c, 0x84, 0x6d, 0x5e, 0xc1, 0x87, 0x97, 0x5f, 0xc9}}
	return a, nil
}

var _assetsDnsServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc9\xb1\x09\xc4\x30\x0c\x05\xd0\xde\x53\x68\x81\x2b\xae\x55\x77\x33\x1c\xa4\x17\xf2\x0f\x11\xc1\xb2\xb1\x14\xcf\x1f\x02\xe9\x1e\xbc\xd3\xbc\x32\xfd\x31\x97\x29\x7e\xaa\xfd\xf2\x2c\x32\x6c\xc3\x0c\xeb\xce\xb4\xbe\xa5\x21\xa5\x4a\x0a\x17\x22\x97\x06\xa6\xea\xf1\x3a\x86\x28\x98\xfa\x80\xc7\x61\x7b\x7e\x9e\xba\x03\x00\x00\xff\xff\x8e\x2c\xf1\x2e\x55\x00\x00\x00")

func assetsDnsServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceAccountYaml,
		"assets/dns/service-account.yaml",
	)
}

func assetsDnsServiceAccountYaml() (*asset, error) {
	bytes, err := assetsDnsServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service-account.yaml", size: 85, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x57, 0x12, 0x50, 0x4d, 0x67, 0x2f, 0x1b, 0x74, 0xa0, 0xa4, 0xbb, 0xa7, 0x59, 0xe9, 0x5a, 0xc6, 0xc1, 0x1a, 0xf8, 0x5f, 0xff, 0x5, 0xdb, 0xc, 0x10, 0x8b, 0xc1, 0x0, 0xcc, 0xf, 0x9f, 0x3a}}
	return a, nil
}

var _assetsDnsServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\x31\x6f\xe2\x40\x10\x85\x7b\xff\x8a\x27\xe8\x4e\xc0\x09\xdd\x51\x9c\xdb\xa3\x89\x52\x80\x14\x48\x3f\x5e\x4f\xcc\x8a\xf5\x8c\xb5\x33\x06\xf1\xef\x23\x4c\x42\x80\x14\x69\x56\xda\x7d\x9f\x3e\x3d\xbd\xdd\x47\xa9\x4b\xbc\x70\x3e\xc4\xc0\x05\x75\xf1\x95\xb3\x45\x95\x12\x87\x79\x31\x86\x50\xcb\x93\xe1\xb4\x8e\x02\x4f\x12\x55\x9c\x0c\x24\x35\x48\x44\x9d\x3c\xaa\x18\x28\x33\x8c\x1d\xe4\xc8\xbd\x78\x6c\xb9\xb0\x8e\x43\x59\x00\x63\x84\xd4\x9b\x73\x7e\x5a\xe3\x18\x53\x42\xc5\xa0\xde\xb5\x25\x8f\x81\x52\x3a\xa1\x25\xa1\x86\xeb\xd9\x00\x1b\x27\x0e\xae\x19\xd1\x1e\x8d\x40\xa7\xd9\xed\x2c\x9d\x0e\x95\x4a\xd4\x62\x05\x70\x09\x4a\x2c\xfe\x0c\x17\xa7\xdc\xb0\xaf\x87\xa7\x2b\x90\xd5\x35\x68\x2a\xb1\x5d\xae\xef\x05\x53\x0f\xdd\x8f\x92\x2f\xe8\x2a\xda\xfc\xbf\x15\xb5\xec\x39\x86\xdb\x36\xff\xe6\x8b\xbf\xdf\x54\x77\xd8\x83\x6a\x8c\xcd\x6a\xb9\x2a\xb1\x95\xa0\x6d\xcb\xe2\x38\xee\x58\x60\x97\xbf\x81\x6b\xa7\x49\x9b\x13\xde\x98\xbc\xcf\x8c\x86\x9c\xcf\x33\xb1\x50\x95\x3e\xf6\xfb\x84\x9e\xf9\x64\x97\xf5\x31\xc5\x68\xdf\x57\x9c\x85\x9d\x6d\x16\xf5\xf7\x4e\xcd\xcf\xa5\x47\xd7\xfc\xd7\xa8\x78\x0f\x00\x00\xff\xff\x82\x42\x75\xa4\x08\x02\x00\x00")

func assetsDnsServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceYaml,
		"assets/dns/service.yaml",
	)
}

func assetsDnsServiceYaml() (*asset, error) {
	bytes, err := assetsDnsServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service.yaml", size: 520, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x18, 0x69, 0xc5, 0xf1, 0xe, 0xc, 0x77, 0xe5, 0x78, 0xce, 0xfc, 0xc2, 0x41, 0xf8, 0x21, 0x87, 0x8a, 0xb7, 0x67, 0xdd, 0x48, 0x94, 0x63, 0x79, 0x69, 0x4e, 0x38, 0x53, 0x3c, 0xdb, 0xc7, 0x13}}
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
	"assets/dns/cluster-role-binding.yaml": assetsDnsClusterRoleBindingYaml,

	"assets/dns/cluster-role.yaml": assetsDnsClusterRoleYaml,

	"assets/dns/daemonset.yaml": assetsDnsDaemonsetYaml,

	"assets/dns/metrics/cluster-role-binding.yaml": assetsDnsMetricsClusterRoleBindingYaml,

	"assets/dns/metrics/cluster-role.yaml": assetsDnsMetricsClusterRoleYaml,

	"assets/dns/metrics/role-binding.yaml": assetsDnsMetricsRoleBindingYaml,

	"assets/dns/metrics/role.yaml": assetsDnsMetricsRoleYaml,

	"assets/dns/namespace.yaml": assetsDnsNamespaceYaml,

	"assets/dns/service-account.yaml": assetsDnsServiceAccountYaml,

	"assets/dns/service.yaml": assetsDnsServiceYaml,
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
	"assets": {nil, map[string]*bintree{
		"dns": {nil, map[string]*bintree{
			"cluster-role-binding.yaml": {assetsDnsClusterRoleBindingYaml, map[string]*bintree{}},
			"cluster-role.yaml":         {assetsDnsClusterRoleYaml, map[string]*bintree{}},
			"daemonset.yaml":            {assetsDnsDaemonsetYaml, map[string]*bintree{}},
			"metrics": {nil, map[string]*bintree{
				"cluster-role-binding.yaml": {assetsDnsMetricsClusterRoleBindingYaml, map[string]*bintree{}},
				"cluster-role.yaml":         {assetsDnsMetricsClusterRoleYaml, map[string]*bintree{}},
				"role-binding.yaml":         {assetsDnsMetricsRoleBindingYaml, map[string]*bintree{}},
				"role.yaml":                 {assetsDnsMetricsRoleYaml, map[string]*bintree{}},
			}},
			"namespace.yaml":       {assetsDnsNamespaceYaml, map[string]*bintree{}},
			"service-account.yaml": {assetsDnsServiceAccountYaml, map[string]*bintree{}},
			"service.yaml":         {assetsDnsServiceYaml, map[string]*bintree{}},
		}},
	}},
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
