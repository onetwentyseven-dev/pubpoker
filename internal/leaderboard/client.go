// package leaderboard is an http api lib that interfaces with the PokerLeaderboard API. Documentation is available
package leaderboard

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ulule/deepcopier"
)

type Client struct {
	baseURI     *url.URL
	client      *http.Client
	cache       inMemoryCache
	credentials Credentials
	config      Config
}

type Config struct {
	LeagueID string
}

type Credentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

const (
	apiURL = "https://api.pokerleaderboards.com"
)

func New(h *http.Client, credentials Credentials, config Config) *Client {
	base, err := url.Parse(apiURL)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize leaderboard client, failed to parse apiURL: %s", err))
	}
	return &Client{
		base,
		h,
		inMemoryCache{},
		credentials,
		config,
	}
}

func (c *Client) apiURL() *url.URL {
	var out = new(url.URL)
	err := deepcopier.Copy(c.baseURI).To(out)
	if err != nil {
		panic("failed to clone baseURL")
	}
	return out

}
