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
	Files     []*File
}

// File is a struct that contains the file and metadata
type File struct {
	Name     string
	DirEntry fs.DirEntry
	Contents []byte
}

// newFile creates a new File struct
func newFile(de fs.DirEntry, contents []byte) *File {
	return &File{
		Name:     de.Name(),
		DirEntry: de,
		Contents: contents,
	}
}

// GetFile returns a file from the resources by name, nil if not found
func (r *Resources) GetFile(h *Helpers, name string) *File {
	h.Logger.Debug(fmt.Sprintf("Getting file: %s", name))
	for _, f := range r.Files {
		if f.Name == name {
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
		l.Error(fmt.Sprintf("Error reading resources: %s", err))
		return nil, err
	}

	r.Files = make([]*File, 0, len(files))
	for _, f := range files {
		n := f.Name()
		l.Debug(fmt.Sprintf("file: %s", n))
		contents, err := r.FS.ReadFile(fmt.Sprintf("resources/%s", n))
		if err != nil {
			l.Error(fmt.Sprintf("Error reading file: %s", err))
			return nil, err
		}
		r.Files = append(r.Files, newFile(f, contents))
	}

	l.Debug("resources:")
	for _, f := range files {
		n := f.Name()
		r.FileNames = append(r.FileNames, n)
		l.Debug(fmt.Sprintf("file: %s", n))
	}

	return r, nil
}
