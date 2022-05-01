package leaderboard

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-http-utils/headers"
	"github.com/pkg/errors"
)

type Token struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
}

const (
	tokenCacheKey = "token"
)

func (c *Client) getToken(ctx context.Context) (*Token, error) {

	cacheData := c.cache.Get(tokenCacheKey)
	if cacheData != nil {
		if token, ok := cacheData.(*Token); ok {
			return token, nil
		}
	}

	body, err := json.Marshal(c.credentials)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode credentials for token request")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/p/auth/login", apiURL), bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request to leaderboard api")
	}

	req.Header.Set(headers.ContentType, "application/json")
	req.Header.Set(headers.Accept, "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute request to leaderboard api")
	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			fmt.Printf("failed to close response body: %s\n", err)
		}
	}()

	data, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		fmt.Println(string(data))
		return nil, errors.Errorf("expected 200, get %d", res.StatusCode)
	}

	var token = new(Token)
	err = json.Unmarshal(data, token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode response body")
	}

	c.cache.Set(tokenCacheKey, token, time.Now().Add(time.Minute*30))

	return token, nil

}
