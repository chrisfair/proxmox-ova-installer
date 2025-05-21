package ova

import (
	"os"
	"path/filepath"
	"github.com/chrisfair/proxmox-ova-installer/internal/filesystem"
	"archive/tar"
)

type OVA struct {
	Path string  // Path to the OVA file
	Dir string // Directory where the OVA file is extracted
	OVF string // Path to the OVF file
	VMDK string // Path to the VMDK file
	fs filesystem.FileSystem // File system interface for file operations
}

// Extract unpacks the .ova to a temporary directory
func (o *OVA) Extract(tarGzPath string) (string, error) {
	if err := o.fs.MkdirAll(o.Dir, os.ModePerm); err != nil {
		return "", err
	}

	file, err := o.fs.Open(tarGzPath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	gzipReader, err := o.fs.NewGzipReader(file)
	if err != nil {
		return "", err
	}

	defer gzipReader.Close()

	tarReader := o.fs.TarReader(gzipReader)
	for {
		head, err := tarReader.Next()
		if err != nil {
			break
		}

		outputPath := filepath.Join(o.Dir, head.Name)
		if head.Typeflag == tar.TypeDir {
			if err := o.fs.MkdirAll(outputPath, head.FileInfo().Mode()); err != nil {
				return "", err
			}
		} else {
			outputFile, err := o.fs.Create(outputPath)
			if err != nil {
				return "", err
			}
			defer outputFile.Close()

			if err := o.fs.Copy(outputFile, tarReader); err != nil {
				return "", err
			}
		}
	}

	return "", nil
}

// CleanUp removes the temporary directory and its contents
func (o *OVA) CleanUp() error {
	return os.RemoveAll(o.Dir)
}
