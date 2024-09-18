package common

import (
	"os"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

type HCLWritable interface {
	WriteHCL(file *hclwrite.File)
}

func OutputHCLToFile(fileName string, writable HCLWritable) error {
	file := hclwrite.NewEmptyFile()
	writable.WriteHCL(file)
	output, err := os.Create(fileName)
	if err != nil {
		return err
	}
	_, err = file.WriteTo(output)
	return err
}
