package armstrong

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL = Default localhost
const HostURL string = "http://localhost"

// AuthStruct - 
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse -
type AuthResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// Client - 
type Client struct {
	HostURL string
	HTTPClient *http.Client
	Token string
	Auth AuthStruct
}

// NewClient - 
func NewClient(host, username, password *string) (*Client, error) {
	
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}

	temp := ""
	if username == nil {
		username = &temp
	}
	if password == nil {
		password = &temp
	}

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second, Transport: transCfg},
		HostURL: HostURL,
		Auth: AuthStruct{
			Username: *username,
			Password: *password,
		},
	}

	if host != nil {
		c.HostURL = *host
	}

	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", c.Token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

