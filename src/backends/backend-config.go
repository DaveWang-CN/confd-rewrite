package backends

import "confdReWrite/src/util"

type BackendsConfig struct {
	BackendNodes util.Nodes
	Backend      string
	Interval     int
	Username     string
	Password     string
	AccessKey    string
	SecretKey    string
	Namespace    string
	Group        string
	Endpoint     string
}
