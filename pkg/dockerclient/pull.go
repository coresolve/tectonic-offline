package dockerclient

import (
    "github.com/heroku/docker-registry-client/registry"
    "encoding/json"
 	"strings"
 	"log"
 	"os"
 	"io"
 	"io/ioutil"
 	"path/filepath"

)

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func DownloadImage(url string, path string, pullsecret map[string]string) (err error) {
	imageurl := strings.Split(url, "/")
	username := ""
	password := ""
	registryserver := "https://" + imageurl[0]

	if imageurl[0] == "quay.io" {
		// Some Images in Quay.io may require pull access
		username = pullsecret["username"]
		password = pullsecret["password"]
	}

	regconnection, err := registry.New(registryserver, username, password)
	check(err)
	//registry.LogfCallback := registry.Quiet

	os.MkdirAll(path, os.ModePerm)
	
	if len(imageurl) == 3 { 
		containerimage := strings.Split(imageurl[2], ":")
		
		// get manifest
		manifest, err := regconnection.Manifest(imageurl[1] + "/" + containerimage[0], containerimage[1])
		check(err)
		pathapplication := filepath.Join(path, imageurl[1] + "-" + containerimage[0] + "-" + containerimage[1])
		os.MkdirAll(pathapplication, os.ModePerm)
		manifestjson, err := manifest.MarshalJSON()
		check(err)
		manifestpath := filepath.Join(pathapplication, "manifest")
	    errcheck := ioutil.WriteFile(manifestpath, manifestjson, os.ModePerm)
	    check(errcheck)

		manifestFSLayers := manifest.FSLayers
		

		for i := range manifestFSLayers {
			digest := manifestFSLayers[i].BlobSum
		    
		    digestjson, _ := json.Marshal(digest)
		    splitdigestjson := strings.Split(string(digestjson),":")
		    imagesha := strings.Split(string(splitdigestjson[1]) , "\"")
		    imagepath := filepath.Join(pathapplication, string(imagesha[0]))
		    
			fo, err := os.Create(imagepath)
		    check(err)
		    defer fo.Close()

			reader, err := regconnection.DownloadLayer(imageurl[1] + "/" + 	containerimage[0], digest)
			if reader != nil {
			    defer reader.Close()
			}
			check(err)

			if _, err := io.Copy(fo, reader); err != nil {
				log.Fatal(err)
			}	
		}
	}

	return nil
}

