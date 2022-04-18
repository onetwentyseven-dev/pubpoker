package leaderboard

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func (c *Client) SearchPlayersByName(ctx context.Context, q string) ([]*Player, error) {
	token, err := c.GetToken(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch token for search players request")
	}

	endpoint := fmt.Sprintf("%s/league/%s/player?filters[Player.name]=%s&order_by[Player.name]=asc", apiURL, c.config.LeagueID, url.QueryEscape(q))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request to search players by name")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.Token))

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
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

	var respData = new(GetPlayersResponse)
	err = json.Unmarshal(data, respData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode response data to structure")
	}

	return respData.Players, nil
}
