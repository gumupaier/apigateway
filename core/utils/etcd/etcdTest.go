package etcd

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/coreos/etcd/client"
)

// Conn :
// connect to etcd
func Conn() client.KeysAPI {
	cfg := client.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)
	return kapi
}
