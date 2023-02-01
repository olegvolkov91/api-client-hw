package apiclient

import (
	"github.com/olegvolkov91/api-client-hw/package/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Start(config *config.Config) *clientPack {
	cl := &http.Client{}

	logger := logrus.New()
	newClient := newClientPack(cl, logger, config)
	newClient.logger.Info("new api client has been created")
	return newClient
}
