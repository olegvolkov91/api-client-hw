package apiclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olegvolkov91/api-client-hw/package/config"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

var (
	errOnParseUsers         = errors.New("data parsing went wrong")
	errOnLoadingUser        = errors.New("loading users went wrong")
	errOnConvertUsersToJson = errors.New("converting users to JSON went wrong")
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
	resp, err := cp.client.Get(fmt.Sprintf("%s/users", cp.config.ClientAddr))
	if err != nil {
		return nil, errOnLoadingUser
	}
	defer resp.Body.Close()

	data, err := cp.parseBody(resp)

	return data, err
}

func (cp *clientPack) CreateUser(u User) error {
	user, err := json.Marshal(u)
	if err != nil {
		return errOnConvertUsersToJson
	}

	postData := bytes.NewBuffer(user)
	req, err := cp.makePrivateRequest(http.MethodPost, fmt.Sprintf("%s/users", cp.config.ClientAddr), postData)
	if err != nil {
		return err
	}

	created, err := cp.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println("created", created)
	return nil
}

func (cp *clientPack) parseBody(data *http.Response) (Users, error) {
	cp.logger.Info("Data is parsing now...")
	body, err := io.ReadAll(data.Body)
	if err != nil {
		return nil, err
	}

	var parsedBody Users
	if err := json.Unmarshal(body, &parsedBody); err != nil {
		return nil, errOnParseUsers
	}
	return parsedBody, nil
}

func (cp *clientPack) makePrivateRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	cp.logger.Infof("\n\nMethod: %s, URL: %s, body: %s\n\n", req.Method, req.URL, req.Body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cp.config.PrimaryToken))
	//Allow CORS here By * or specific origin
	req.Header.Set("Access-Control-Allow-Origin", "*")
	return req, nil
}
