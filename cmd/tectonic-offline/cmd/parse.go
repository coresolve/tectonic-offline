package cmd

import (
	"github.com/alekssaul/tectonic-offline/pkg/spec"
	"github.com/spf13/cobra"
	"log"
	"fmt"
	"path/filepath"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parses Tectonic-Installer config variable file",
	Run: func(cmd *cobra.Command, args []string) {
		ParseTerraformTFVARS(cfgFile, tectonicImagesVar)
	},
}

func init() {
	RootCmd.AddCommand(parseCmd)
	parseCmd.PersistentFlags().StringVarP(&cfgFile, "config", "", "config.tf", "Path to config.tf file")
	parseCmd.PersistentFlags().StringVarP(&tectonicImagesVar, "imagevar", "", "tectonic_container_images", "Tectonic Images Variable used in TFVars File")
}

func ParseTerraformTFVARS(cfgFile string, tectonicImagesVar string) {

	tfvarsfile, err := filepath.Abs(cfgFile)
	if err != nil {
		log.Fatal(err)
	}

	v, err := spec.TerraformConfig(tfvarsfile)
	if err != nil {
		log.Fatal(err)
	}

	images:= spec.TectonicImages(v, tectonicImagesVar)
	tectonicversion := spec.TectonicVersion(v)

	fmt.Println("Tectonic Version: " , tectonicversion)
	fmt.Println("Container Images: ")
	for i := range images {
		fmt.Println("\t", images[i])
	}
	
}
