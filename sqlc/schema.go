package sqlc

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func sqlc_tmpl_fields_tmpl() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xbc, 0x56,
		0x41, 0x6f, 0xa3, 0x3a, 0x10, 0x3e, 0xe3, 0x5f, 0x31, 0x8a, 0xaa, 0x27,
		0x88, 0x68, 0x72, 0x8f, 0xd4, 0x43, 0xfa, 0x4a, 0x5f, 0xf3, 0x94, 0xa6,
		0x55, 0x42, 0xb7, 0xda, 0xd3, 0x8a, 0x12, 0x93, 0xb5, 0x96, 0x1a, 0x0a,
		0x26, 0xdb, 0x0a, 0xf1, 0xdf, 0x77, 0x6c, 0x13, 0x02, 0x09, 0x61, 0x83,
		0x56, 0xdb, 0x1c, 0x82, 0x19, 0xdb, 0xdf, 0xf7, 0xcd, 0x8c, 0x67, 0xcc,
		0x78, 0x0c, 0xee, 0xdd, 0x6c, 0x05, 0xb7, 0xb3, 0xb9, 0x03, 0xcf, 0xd3,
		0x15, 0x4c, 0x9f, 0xdc, 0x87, 0xff, 0x9c, 0x85, 0xb3, 0x9c, 0xba, 0xce,
		0x0d, 0x5c, 0xc2, 0x74, 0xf1, 0x15, 0x9c, 0x9b, 0x99, 0xbb, 0x02, 0xf7,
		0x41, 0x2f, 0x7d, 0x9e, 0xcd, 0xe7, 0x70, 0xed, 0xc0, 0xfc, 0x61, 0xe5,
		0xc2, 0xf3, 0x9d, 0xb3, 0x80, 0x99, 0x0b, 0x68, 0x5f, 0x3a, 0xd5, 0x3e,
		0x42, 0xf2, 0x1c, 0x2e, 0xe2, 0x84, 0xae, 0x53, 0x98, 0x5c, 0xc1, 0x48,
		0x8e, 0x98, 0xef, 0x09, 0x9a, 0x42, 0x51, 0x10, 0x12, 0x7b, 0xfe, 0x0f,
		0x6f, 0x43, 0x21, 0x7d, 0x0b, 0x7d, 0x42, 0xd8, 0x6b, 0x1c, 0x25, 0x02,
		0x4c, 0x62, 0x0c, 0x04, 0x7b, 0xa5, 0x03, 0x62, 0x11, 0x22, 0x3e, 0x62,
		0x0a, 0x33, 0x9e, 0xd2, 0x44, 0xac, 0xa8, 0x58, 0x09, 0x1a, 0x03, 0xe3,
		0x82, 0x26, 0x81, 0xe7, 0x53, 0xc8, 0x89, 0x81, 0xf0, 0x89, 0xc7, 0x11,
		0xe2, 0xe2, 0x9b, 0x0d, 0x17, 0x42, 0x91, 0xc8, 0x3d, 0x0a, 0x1f, 0x0c,
		0xdc, 0x23, 0x05, 0x88, 0xd1, 0x63, 0x42, 0x03, 0xf6, 0x8e, 0x46, 0xf3,
		0xe0, 0xfd, 0x96, 0xd1, 0x70, 0x6d, 0x83, 0xb6, 0xce, 0x19, 0x42, 0x7b,
		0x21, 0x9a, 0xad, 0x3d, 0xe9, 0x7d, 0x94, 0x50, 0x49, 0x8c, 0x70, 0xb8,
		0x8a, 0xf2, 0xb5, 0x84, 0x2e, 0x4a, 0x69, 0x4f, 0xf1, 0x1a, 0xbd, 0xf9,
		0x64, 0x69, 0x15, 0xe9, 0x29, 0x69, 0x9d, 0xd4, 0x41, 0xc6, 0x7d, 0x30,
		0x19, 0x0c, 0x99, 0xf2, 0xd0, 0x82, 0x16, 0x25, 0x01, 0xb4, 0x6b, 0xd9,
		0x9e, 0x13, 0x28, 0xe9, 0x7c, 0x42, 0x45, 0x96, 0x70, 0x60, 0xa3, 0x94,
		0x0a, 0x33, 0xb0, 0xb7, 0x16, 0xca, 0xda, 0x4b, 0x3c, 0x47, 0x60, 0x06,
		0xc3, 0x4c, 0xf9, 0xf9, 0xc7, 0x02, 0x8f, 0xc2, 0x55, 0x13, 0x98, 0x9d,
		0x10, 0x38, 0x96, 0xbf, 0x32, 0xc7, 0x4b, 0x1a, 0x84, 0xd4, 0x17, 0xde,
		0x4b, 0x48, 0x1b, 0x19, 0xee, 0x74, 0xc2, 0x68, 0xd3, 0x67, 0x72, 0xef,
		0x15, 0x4f, 0xbb, 0x48, 0x18, 0xdf, 0x58, 0xad, 0x1e, 0x90, 0xbe, 0x79,
		0x4c, 0x61, 0x98, 0x52, 0x29, 0x8f, 0x45, 0xbc, 0x1d, 0xf2, 0xf7, 0xa4,
		0xb5, 0x78, 0xfc, 0x83, 0xf3, 0x22, 0x9a, 0x47, 0x3f, 0x69, 0x72, 0xbc,
		0x2e, 0x97, 0x48, 0x13, 0x90, 0xff, 0x52, 0x9e, 0x16, 0x20, 0x40, 0x45,
		0xe6, 0x13, 0xb9, 0x6d, 0xcd, 0x38, 0x01, 0x51, 0xb4, 0x26, 0xad, 0x33,
		0x6a, 0x3a, 0xa5, 0x5d, 0x4c, 0x52, 0x6e, 0xe6, 0x0b, 0x29, 0xac, 0xa6,
		0x9e, 0x18, 0xfa, 0x00, 0xac, 0xe8, 0xee, 0x2c, 0x10, 0x03, 0x03, 0x00,
		0xb7, 0x18, 0x04, 0x19, 0x7b, 0x62, 0x78, 0x21, 0xf3, 0xd2, 0xdd, 0xea,
		0x1a, 0xd1, 0x31, 0x41, 0xa3, 0x4d, 0xb8, 0x12, 0x4b, 0xe7, 0xde, 0x70,
		0xde, 0xcc, 0xad, 0x17, 0x66, 0xb4, 0xe5, 0x10, 0xff, 0x1b, 0xf1, 0x35,
		0xd3, 0x44, 0xb3, 0xf4, 0x60, 0x5d, 0x13, 0xde, 0x82, 0xff, 0x23, 0xc6,
		0xf7, 0xeb, 0x51, 0x4a, 0x99, 0x2a, 0x1f, 0x86, 0x5d, 0x8e, 0x5b, 0x95,
		0x33, 0xe6, 0x7e, 0x58, 0xcb, 0x8f, 0x3f, 0x42, 0x18, 0x09, 0x77, 0x26,
		0x5a, 0xe0, 0x63, 0x65, 0x55, 0x40, 0x08, 0xd9, 0x2b, 0xdf, 0xc4, 0x30,
		0x74, 0xca, 0x91, 0x57, 0x65, 0x1d, 0x0d, 0x65, 0xde, 0xfd, 0x91, 0x1a,
		0x48, 0x0b, 0x4a, 0xc1, 0x15, 0x10, 0xe0, 0xb8, 0xe8, 0x21, 0x6d, 0x9a,
		0x9a, 0xf5, 0x74, 0xf5, 0xd5, 0xa6, 0x85, 0x95, 0xba, 0xe0, 0x40, 0x15,
		0x28, 0x4d, 0x2a, 0x58, 0x36, 0x28, 0x96, 0x89, 0x7e, 0xf4, 0x12, 0x28,
		0x37, 0x60, 0x1a, 0xb4, 0xbe, 0x46, 0x12, 0x14, 0x56, 0x0f, 0xa8, 0x7b,
		0xef, 0xe3, 0x85, 0x1e, 0xe3, 0xb1, 0x60, 0x87, 0x05, 0x57, 0x57, 0x30,
		0x18, 0x48, 0xdb, 0x9e, 0x44, 0x7a, 0x86, 0x21, 0x05, 0x1a, 0xa6, 0xb4,
		0x39, 0xa3, 0xe9, 0x7b, 0x45, 0x7b, 0x81, 0x60, 0xed, 0xbe, 0x28, 0x9a,
		0xf3, 0x81, 0x1e, 0xbd, 0x84, 0x72, 0x61, 0x5a, 0xb5, 0x2a, 0x6c, 0xc0,
		0xe9, 0xba, 0x54, 0x9d, 0x00, 0x2e, 0x2f, 0x0f, 0x3b, 0x41, 0x2c, 0x3b,
		0x41, 0xf9, 0x31, 0x52, 0x9c, 0x4f, 0xaa, 0xbe, 0x60, 0x46, 0xea, 0xa5,
		0x2a, 0x0a, 0xbc, 0x81, 0x24, 0x4e, 0x57, 0x9d, 0xd6, 0x84, 0x55, 0xb6,
		0xfc, 0x9a, 0xe1, 0x88, 0x6f, 0x26, 0xfa, 0xbc, 0x95, 0x6f, 0xf9, 0x17,
		0x59, 0xcb, 0x13, 0x90, 0x88, 0xb6, 0x9e, 0xc1, 0xf3, 0x53, 0xf4, 0x89,
		0xb0, 0xd6, 0x28, 0xeb, 0xfe, 0x84, 0xc4, 0xce, 0x16, 0x51, 0x93, 0xda,
		0xb0, 0xe7, 0xf3, 0xef, 0x78, 0x76, 0x7d, 0x1b, 0x96, 0xf2, 0xa9, 0xe5,
		0x3d, 0xee, 0xbe, 0xe0, 0x26, 0x25, 0x69, 0x65, 0x40, 0xf0, 0xdd, 0xa5,
		0xb5, 0x6f, 0xc8, 0x2a, 0x0d, 0xca, 0x89, 0xc3, 0x1b, 0xfc, 0xb0, 0x9d,
		0xda, 0xf0, 0x37, 0xef, 0x0a, 0xf5, 0x5f, 0x90, 0xdd, 0x25, 0x31, 0x3e,
		0xbf, 0x10, 0xb7, 0x1b, 0xb3, 0x6a, 0x0f, 0x46, 0xbd, 0x17, 0x62, 0x77,
		0xc3, 0x49, 0xab, 0x57, 0x25, 0xbe, 0x9f, 0xc6, 0xc2, 0xc9, 0x5e, 0x58,
		0x8c, 0x9b, 0xc7, 0x6d, 0xab, 0x84, 0x62, 0xdc, 0x6a, 0x64, 0xe2, 0x57,
		0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xd0, 0x65, 0xf7, 0xdd, 0x0b, 0x00,
		0x00,
	},
		"sqlc/tmpl/fields.tmpl",
	)
}

func sqlc_tmpl_schema_tmpl() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x9c, 0x54,
		0x5f, 0x6f, 0xda, 0x3e, 0x14, 0x7d, 0x26, 0x9f, 0xe2, 0xca, 0xaa, 0x7e,
		0x4a, 0x7e, 0x62, 0xce, 0x3b, 0x12, 0x0f, 0x99, 0x9a, 0xae, 0x91, 0x32,
		0xa8, 0x1a, 0x57, 0x68, 0x9a, 0x26, 0x64, 0x52, 0x87, 0x45, 0x0b, 0x49,
		0x66, 0x9b, 0x6d, 0x08, 0xf1, 0xdd, 0xe7, 0x3f, 0xa4, 0xa4, 0xc5, 0x94,
		0x68, 0x08, 0x21, 0xe3, 0x7b, 0xef, 0xb9, 0xe7, 0x1c, 0x5f, 0x3b, 0x0c,
		0x81, 0xdc, 0x27, 0x19, 0xdc, 0x25, 0x69, 0x0c, 0x8b, 0x28, 0x83, 0xe8,
		0x89, 0xcc, 0x3f, 0xc5, 0xb3, 0xf8, 0x31, 0x22, 0xf1, 0x2d, 0x7c, 0x80,
		0x68, 0xf6, 0x05, 0xe2, 0xdb, 0x84, 0x64, 0x40, 0xe6, 0x36, 0x75, 0x91,
		0xa4, 0x29, 0x7c, 0x8c, 0x21, 0x9d, 0x67, 0x04, 0x16, 0xf7, 0xf1, 0x0c,
		0x12, 0x02, 0x6a, 0xff, 0x31, 0x7e, 0xa9, 0xf3, 0xbc, 0xfd, 0x1e, 0x6e,
		0xe4, 0xae, 0x65, 0x02, 0x26, 0x53, 0xc0, 0xc4, 0xac, 0x0e, 0x07, 0xcf,
		0x6b, 0x69, 0xfe, 0x83, 0xae, 0x19, 0xa8, 0x38, 0x7e, 0x38, 0xae, 0xf5,
		0x7e, 0xb9, 0x69, 0x1b, 0x2e, 0xc1, 0xf7, 0x46, 0x68, 0x5d, 0xca, 0xef,
		0xdb, 0x15, 0xce, 0x9b, 0x4d, 0xc8, 0x59, 0xd5, 0xb4, 0x22, 0x14, 0x3f,
		0xab, 0xdc, 0xfc, 0x20, 0x15, 0x16, 0x92, 0x97, 0xf5, 0x5a, 0x20, 0x2f,
		0x30, 0x5d, 0x38, 0xad, 0x15, 0xc4, 0xcd, 0x72, 0xac, 0xfa, 0xd9, 0x5e,
		0x74, 0x55, 0x1d, 0x9b, 0x69, 0x02, 0xba, 0x93, 0x6c, 0xd2, 0xe6, 0x37,
		0xe3, 0x2a, 0x03, 0xcf, 0xe8, 0x46, 0x37, 0x04, 0x85, 0xb2, 0xcd, 0x25,
		0xec, 0xbd, 0xd1, 0x6b, 0x8c, 0x42, 0x63, 0xa8, 0xbc, 0xbb, 0x92, 0x55,
		0xcf, 0x06, 0x65, 0x64, 0x00, 0x9e, 0xda, 0x56, 0x03, 0x14, 0x27, 0x00,
		0x45, 0x07, 0x6b, 0x95, 0x85, 0x51, 0xa7, 0xb6, 0x4c, 0x89, 0x49, 0x67,
		0xf5, 0xb3, 0xa9, 0xa4, 0x55, 0x49, 0x05, 0x58, 0xc2, 0x9e, 0xe2, 0x53,
		0x6c, 0xeb, 0x1c, 0x7c, 0x09, 0xff, 0x3b, 0x39, 0x05, 0x90, 0x88, 0x8c,
		0x55, 0x2c, 0x97, 0x5a, 0x81, 0x1f, 0xc0, 0x7e, 0x40, 0x89, 0x5e, 0xa8,
		0x54, 0xdb, 0x43, 0xcb, 0xe1, 0x4c, 0x6e, 0x79, 0x0d, 0xc8, 0x99, 0x8f,
		0x86, 0xb0, 0x88, 0x84, 0x4f, 0x8f, 0x80, 0x81, 0x95, 0x79, 0x62, 0xd5,
		0xeb, 0xf0, 0x9f, 0xb3, 0x5c, 0xc5, 0xaf, 0x3b, 0xea, 0xb6, 0x74, 0x02,
		0x12, 0x3b, 0x03, 0x63, 0x5b, 0xd1, 0xb9, 0x6a, 0x6d, 0x9d, 0x50, 0xb5,
		0x7d, 0x18, 0xa4, 0x47, 0xa7, 0xbb, 0x3c, 0x92, 0xd8, 0x20, 0x0d, 0xc1,
		0xf8, 0x4c, 0x77, 0x2b, 0x76, 0x0e, 0x54, 0x16, 0x1d, 0x08, 0x4c, 0xa7,
		0x80, 0x90, 0xde, 0xbb, 0x76, 0x02, 0xa3, 0x03, 0xb0, 0x4a, 0xb0, 0x7e,
		0x6a, 0x47, 0xc4, 0xea, 0x09, 0xf5, 0xe7, 0x6c, 0xb6, 0x77, 0xd6, 0xc6,
		0xee, 0x22, 0x5d, 0x23, 0x6c, 0xef, 0x1f, 0x7e, 0xe0, 0xac, 0x28, 0xff,
		0x74, 0xc3, 0xe9, 0xd7, 0x3a, 0xfc, 0xea, 0x68, 0x9d, 0x79, 0x3d, 0x8f,
		0x9c, 0x49, 0x7e, 0xff, 0x98, 0x5e, 0x9a, 0x8e, 0x8f, 0xc8, 0x02, 0x13,
		0x1b, 0x33, 0xed, 0x82, 0x40, 0x69, 0x3a, 0x9d, 0x5e, 0xa7, 0xee, 0x1a,
		0x7f, 0x3b, 0x2d, 0xca, 0xec, 0xaf, 0xdf, 0x0c, 0x85, 0xb7, 0xbc, 0xfa,
		0xdb, 0xc3, 0x66, 0xee, 0xfc, 0xc2, 0x5e, 0x92, 0x81, 0x9c, 0x53, 0x88,
		0x82, 0x37, 0x73, 0x68, 0x0e, 0xab, 0xa7, 0xec, 0xfd, 0xc7, 0xe8, 0x17,
		0xe5, 0xb0, 0x5c, 0xba, 0x1f, 0xa3, 0xe9, 0xa5, 0xcb, 0x64, 0xcb, 0x9c,
		0x34, 0x2f, 0x16, 0x69, 0x93, 0xbc, 0x7f, 0x7b, 0xd4, 0x26, 0x8e, 0x57,
		0xcd, 0xbf, 0x40, 0xfa, 0x3d, 0x9b, 0x7a, 0xa6, 0xa8, 0xef, 0xe9, 0xef,
		0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x25, 0xac, 0x72, 0x6d, 0x06,
		0x00, 0x00,
	},
		"sqlc/tmpl/schema.tmpl",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"sqlc/tmpl/fields.tmpl": sqlc_tmpl_fields_tmpl,
	"sqlc/tmpl/schema.tmpl": sqlc_tmpl_schema_tmpl,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"sqlc": &_bintree_t{nil, map[string]*_bintree_t{
		"tmpl": &_bintree_t{nil, map[string]*_bintree_t{
			"fields.tmpl": &_bintree_t{sqlc_tmpl_fields_tmpl, map[string]*_bintree_t{
			}},
			"schema.tmpl": &_bintree_t{sqlc_tmpl_schema_tmpl, map[string]*_bintree_t{
			}},
		}},
	}},
}}
