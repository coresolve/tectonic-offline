package spec

import (
 	"github.com/hashicorp/hcl"
 	"io/ioutil"
 	"log"
	"fmt"
)

const (
	TectonicVersionVar         = "tectonic_versions"
	TectonicVersionKey        = "tectonic"
)


type Configuration map[string]interface{}


func TerraformConfig(file string) (s Configuration, err error) {
	input, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	err = hcl.Unmarshal(input, &s)
	if err != nil {
		return s, fmt.Errorf("unable to parse HCL: %s", err)
	}
	
	return s, nil
}

func TectonicImages(tfvars Configuration, TectonicImagesVar string) (images []string){
	tfvariables := tfvars["variable"].([]map[string]interface{})
	for i := range tfvariables{
		mymap := tfvariables[i]
		if  mymap[TectonicImagesVar] != nil {
			md :=  mymap[TectonicImagesVar].([]map[string]interface{})[0]
			imagemap := md["default"]
			m := imagemap.([]map[string]interface{})[0]
			for _, v := range m {
        		images = append(images, v.(string))
			}
   		}
	}

	return images
}

func TectonicVersion(tfvars Configuration) (tectonicversion string){
	tfvariables := tfvars["variable"].([]map[string]interface{})
	for i := range tfvariables{
		mymap := tfvariables[i]
		if  mymap[TectonicVersionVar] != nil {
			md :=  mymap[TectonicVersionVar].([]map[string]interface{})[0]
			imagemap := md["default"]
			m := imagemap.([]map[string]interface{})[0]
			for k, v := range m {
        		if k == TectonicVersionKey {tectonicversion = v.(string)}
			}
   		}
	}

	return tectonicversion
}

