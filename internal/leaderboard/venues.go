package leaderboard

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type GetVenueResponse struct {
	Venues     []*Venue            `json:"venues"`
	Pagination *ResponsePagination `json:"pagination"`
}

func (c *Client) Venues(ctx context.Context) ([]*Venue, error) {

	token, err := c.GetToken(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch token for search players request")
	}

	endpoint := fmt.Sprintf("%s/league/%s/season", apiURL, c.config.LeagueID)
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

	if res.StatusCode != http.StatusOK {
		var out = new(ErrorResponse)
		err := json.NewDecoder(res.Body).Decode(out)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to decode error response: got %d from venues endpoint", res.StatusCode)
		}

		return nil, out
	}

	var out = new(GetVenueResponse)
	err = json.NewDecoder(res.Body).Decode(out)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decode venues response: got %d from venues endpoint", res.StatusCode)
	}

	return out.Venues, nil

}
