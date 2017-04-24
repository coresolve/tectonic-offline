package cmd

import (
	"github.com/alekssaul/tectonic-offline/pkg/spec"
	"github.com/alekssaul/tectonic-offline/pkg/zip"
	"github.com/alekssaul/tectonic-offline/pkg/dockerclient"
	"github.com/spf13/cobra"
	"log"
	"io/ioutil"
	"path/filepath"
	"os/user"
	"os"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches Containers in use by Tectonic-Installer",
	Run: func(cmd *cobra.Command, args []string) {
		FetchDockerImages(cfgFile, tectonicImagesVar, coreospullsecret)
	},
}

func init() {
	usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }
    pullsecret := filepath.Join(filepath.Join(usr.HomeDir, ".docker"),"config.json")

	RootCmd.AddCommand(fetchCmd)
	fetchCmd.PersistentFlags().StringVarP(&cfgFile, "config", "", "config.tf", "Path to config.tf file")
	fetchCmd.PersistentFlags().StringVarP(&tectonicImagesVar, "imagevar", "", "tectonic_container_images", "Tectonic Images Variable used in TFVars File")
	fetchCmd.PersistentFlags().StringVarP(&coreospullsecret, 
		"coreospullsecret", 
		"", 
		pullsecret, 
		"CoreOS Pull Secret file. ")
}

func FetchDockerImages(cfgFile string, tectonicImagesVar string, coreospullsecret string) {
	
	tfvarsfile, err := filepath.Abs(cfgFile)
	check(err)

	v, err := spec.TerraformConfig(tfvarsfile)
	check(err)

	images := spec.TectonicImages(v, tectonicImagesVar)
	tectonicversion := spec.TectonicVersion(v)

	pathtmp, err := ioutil.TempDir("", "tmp")
	check (err)

	pathtectonic := filepath.Join(pathtmp, "tectonic-offline")

	auth, err := dockerclient.ParseQuayConfig(coreospullsecret, "quay.io")

	log.Println("Tectonic Version: " , tectonicversion)
	for i := range images {
		log.Printf("Fetching : " + images[i])
		dockerclient.DownloadImage(images[i], pathtectonic, auth)
	}

	log.Println("Download completed.. Compressing Container Images.. ")
	errZip := zip.Compress (pathtectonic, "tectonic-offline-" + tectonicversion + ".zip" )
	check(errZip)
	log.Println("Removing cached files")
	defer os.RemoveAll(pathtmp) 
	log.Printf("Done!")
	
}
