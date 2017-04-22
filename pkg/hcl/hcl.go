package hcl

import (
 	"github.com/hashicorp/hcl/parser"

)

type Config struct {
	Region      string
	AccessKey   string
	SecretKey   string
	Bucket      string
	Directories []DirectoryConfig
}