package maldocker

// Sample Virtualbox create independent of Machine CLI.
import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/machine/drivers/virtualbox"
	"github.com/docker/machine/libmachine"
)

// MakeDockerMachine creates a new docker host via docker-machine
func MakeDockerMachine(host string) {
	// log.SetDebug(true)

	client := libmachine.NewClient("/tmp/automatic")

	hostName := host

	// Set some options on the provider...
	driver := virtualbox.NewDriver(hostName, "/tmp/automatic")
	driver.CPU = 2
	driver.Memory = 2048

	data, err := json.Marshal(driver)
	if err != nil {
		log.Fatal(err)
	}

	pluginDriver, err := client.NewPluginDriver("virtualbox", data)
	if err != nil {
		log.Fatal(err)
	}

	h, err := client.NewHost(pluginDriver)
	if err != nil {
		log.Fatal(err)
	}

	h.HostOptions.EngineOptions.StorageDriver = "overlay"

	if err := client.Create(h); err != nil {
		log.Fatal(err)
	}

	out, err := h.RunSSHCommand("df -h")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Results of your disk space query:\n%s\n", out)

	fmt.Println("Powering down machine now...")
	if err := h.Stop(); err != nil {
		log.Fatal(err)
	}
}