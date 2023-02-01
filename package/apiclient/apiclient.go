package apiclient

import (
	"github.com/olegvolkov91/api-client-hw/package/config"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Start(config *config.Config) *clientPack {
	cl := &http.Client{
		Timeout: time.Second * 10,
	}

	logger := logrus.New()
	return newClientPack(cl, logger, config)
}
