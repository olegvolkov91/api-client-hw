package apiclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olegvolkov91/api-client-hw/package/config"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

var (
	errOnParseUsers  = errors.New("users parsing went wrong")
	errOnLoadingUser = errors.New("loading users went wrong")
)

type clientPack struct {
	client *http.Client
	logger *logrus.Logger
	config *config.Config
}

func newClientPack(client *http.Client, logger *logrus.Logger, config *config.Config) *clientPack {
	return &clientPack{client, logger, config}
}

func (cp *clientPack) GetUsers() (Users, error) {
	start := time.Now()
	cp.logger.Info("Users loading process has been started ...")
	defer cp.logger.Infof("Users are loaded in %v", time.Now().Sub(start))
	resp, err := cp.client.Get(fmt.Sprintf("%s/users", cp.config.ClientAddr))

	if err != nil {
		return nil, errOnLoadingUser
	}
	defer resp.Body.Close()

	var users []User

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &users); err != nil {
		return nil, errOnParseUsers
	}

	return users, nil
}
