package dist

//go:generate go-bindata -o dist_gen.go -ignore 'map|go' -pkg dist ../ui/build/...

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
)

// DebugAssets serves assets via a specified directory
type DebugAssets struct {
	Dir     string // Dir is a directory location of asset files
	Default string // Default is the file to serve if file is not found.
}

// Handler is an http.FileServer for the Dir
func (d *DebugAssets) Handler() http.Handler {
	return http.FileServer(NewDir(d.Dir, d.Default))
}

// BindataAssets serves assets from go-bindata
type BindataAssets struct {
	Prefix  string // Prefix is prepended to the http file request
	Default string // Default is the file to serve if the file is not found
}

// Handler serves go-bindata using a go-bindata-assetfs façade
func (b *BindataAssets) Handler() http.Handler {
	// def wraps the assets to return the default file if the file doesn't exist
	def := func(name string) ([]byte, error) {
		// If the named asset exists, then return it directly.
		octets, err := Asset(name)
		if err != nil {
			// If this is at / then we just error out so we can return a Directory
			// This directory will then be redirected by go to the /index.html
			if name == b.Prefix {
				return nil, err
			}
			// If this is anything other than slash, we just return the default
			// asset.  This default asset will handle the routing.
			return Asset(b.Default)
		}
		return octets, nil
	}
	var dir http.FileSystem = &assetfs.AssetFS{
		Asset:     def,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    b.Prefix,
	}
	return http.FileServer(dir)
}
