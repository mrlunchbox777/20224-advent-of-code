package common

import (
	"embed"
	"fmt"
	"io/fs"
	"log/slog"

	"github.com/spf13/viper"
)

//go:embed resources/*
var resources embed.FS

// Resources is a struct that contains the resources
type Resources struct {
	FS        embed.FS
	FileNames []string
	Files     []fs.DirEntry
}

// GetFile returns a file from the resources by name, nil if not found
func (r *Resources) GetFile(name string) fs.DirEntry {
	for _, f := range r.Files {
		if f.Name() == name {
			return f
		}
	}
	return nil
}

// NewResources creates a new Resources struct
func NewResources(l *slog.Logger, v *viper.Viper) (*Resources, error) {
	r := &Resources{
		FS: resources,
	}

	files, err := r.FS.ReadDir("resources")
	if err != nil {
		return nil, err
	}
	r.Files = files

	l.Debug("resources:")
	for _, f := range files {
		n := f.Name()
		r.FileNames = append(r.FileNames, n)
		l.Debug(fmt.Sprintf("file: %s", n))
	}

	return r, nil
}
