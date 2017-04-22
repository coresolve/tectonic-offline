package cmd

import (
	"github.com/alekssaul/tectonic-offline/pkg/spec"
	//"github.com/alekssaul/tectonic-offline/pkg/docker"
	"github.com/spf13/cobra"
	"log"
	"fmt"
	"path/filepath"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches Containers in use by Tectonic-Installer",
	Run: func(cmd *cobra.Command, args []string) {
		FetchDockerImages(cfgFile, tectonicImagesVar)
	},
}

func init() {
	RootCmd.AddCommand(fetchCmd)
	fetchCmd.PersistentFlags().StringVarP(&cfgFile, "config", "", "config.tf", "Path to config.tf file")
	fetchCmd.PersistentFlags().StringVarP(&tectonicImagesVar, "imagevar", "", "tectonic_container_images", "Tectonic Images Variable used in TFVars File")
}

func FetchDockerImages(cfgFile string, tectonicImagesVar string) {
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
	for i := range images {
		fmt.Println("Fetching : ", images[i])
	}
	
}
