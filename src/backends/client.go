package backends

import (
	"confdReWrite/src/backends/nacos"
	"errors"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
)

// The StoreClient interface is implemented by objects that can retrieve
// key/value pairs from a backend store.
type StoreClient interface {
	GetValues(keys []string) (map[string]string, error)
	WatchPrefix(prefix string, keys []string, waitIndex uint64, stopChan chan bool) (uint64, error)
}

// New is used to create a storage client based on our configuration.
func New(config BackendsConfig) (StoreClient, error) {

	switch config.Backend {
	case "nacos":
		return nacos.NewNacosClient(config.BackendNodes, config.Group, constant.ClientConfig{
			NamespaceId: config.Namespace,
			AccessKey:   config.AccessKey,
			SecretKey:   config.SecretKey,
			Endpoint:    config.Endpoint,
			Username:    config.Username,
			Password:    config.Password,
		})
	}
	return nil, errors.New("Invalid backend")
}
