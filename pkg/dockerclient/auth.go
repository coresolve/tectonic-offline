package dockerclient

import (
	"github.com/docker/docker/cliconfig"
	"log"
	"os"
)

type DockerRegistryAuth map[string]string

func ParseQuayConfig(path string, registry string) (d DockerRegistryAuth, err error) {
	
	f, err := os.Open(path)
	if err != nil {
        log.Fatal(err)
    }

	config, err := cliconfig.LoadFromReader(f)
	if err != nil {
        log.Fatal(err)
    }

    m := make(map[string]string)
	m["username"] = config.AuthConfigs[registry].Username
	m["password"] = config.AuthConfigs[registry].Password
	m["servername"] = config.AuthConfigs[registry].ServerAddress

	if m["username"] == "" {
		log.Fatal("Could not parse Docker Pull Secret Configuration")
	}

	return m, err
}
