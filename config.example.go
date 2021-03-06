package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Tick     int    `yaml:"tick"`      // tick (in seconds) between receiving data
	Timeout  int    `yaml:"timeout"`   // threshold (in seconds) after a node is considered/displayed as offline
	RouterIp string `yaml:"router_ip"` // ip of cluster-smi-router
	SSLCert string `yaml:"ssl_cert"` // certificate for https
	SSLKey string `yaml:"ssl_key"` // certificate for https
	Ports    struct {
		Nodes   string `yaml:"nodes"`   // port of cluster-smi-router, which nodes send to
		Clients string `yaml:"clients"` // port of cluster-smi-router, where clients subscribe to
		SSE string `yaml:"sse"` // port for SSE
	} `yaml:"ports"`
}

func LoadConfig() Config {

	c := Config{}

	c.RouterIp = "127.0.0.1"
	c.SSLCert = ""
	c.SSLKey = ""
	c.Tick = 3
	c.Timeout = 180
	c.Ports.Nodes = "9080"
	c.Ports.Clients = "9081"
	c.Ports.SSE = "3000"

	if os.Getenv("CLUSTER_SMI_CONFIG_PATH") != "" {
		fn := os.Getenv("CLUSTER_SMI_CONFIG_PATH")

		yamlFile, err := ioutil.ReadFile(fn)
		if err != nil {
			log.Fatalf("Error: %v ", err)
		}

		err = yaml.Unmarshal(yamlFile, &c)
		if err != nil {
			log.Fatalf("Error in %s %v", fn, err)
		}
	}

	return c
}

func (c Config) Print() {

	if os.Getenv("CLUSTER_SMI_CONFIG_PATH") != "" {
		log.Println("Read configuration from", os.Getenv("CLUSTER_SMI_CONFIG_PATH"))
	} else {
		log.Println("use default configuration")
	}

	log.Println("  Tick:", c.Tick)
	log.Println("  Timeout:", c.Timeout)
	log.Println("  RouterIp:", c.RouterIp)
	log.Println("  SSLCert:", c.SSLCert)
	log.Println("  SSLKey:", c.SSLKey)
	log.Println("  Ports:")
	log.Println("    Nodes:", c.Ports.Nodes)
	log.Println("    Clients:", c.Ports.Clients)
	log.Println("    SSE:", c.Ports.SSE)

}
