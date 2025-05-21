package ova

import "os"

type OVA struct {
	Path string  // Path to the OVA file
	Dir string // Directory where the OVA file is extracted
	OVF string // Path to the OVF file
	VMDK string // Path to the VMDK file
}

// Extract unpacks the .ova to a temporary directory
func (o *OVA) Extract() error {
	// TODO:  Implement the extraction logic
	return nil
}

// CleanUp removes the temporary directory and its contents
func (o *OVA) CleanUp() error {
	return os.RemoveAll(o.Dir)
}
