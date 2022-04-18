package leaderboard

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

func (s *Client) CurrentSeason(ctx context.Context, seasons []*Season) (*Season, error) {

	if len(seasons) == 0 {
		return nil, errors.New("expected at least 1 season, got 0")
	}

	var now = time.Now().Unix()
	for _, s := range seasons {
		if s.StartDate.Unix() < now && s.EndDate.Unix() > now {
			return s, nil
		}
	}

	// If we got here, there is not an active season, loop over the seasons again and find the most recent
	var latest = new(Season)
	for _, s := range seasons {
		if s.StartDate.Unix() > latest.StartDate.Unix() {
			fmt.Printf("setting latest to season %s with start time of %s\n", s.ID, s.StartDate)
			latest = s
		}
	}

	return latest, nil

}

func (c *Client) Seasons(ctx context.Context) ([]*Season, error) {
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

	data, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		fmt.Println(string(data))
		return nil, errors.Errorf("expected 200, get %d", res.StatusCode)
	}

	var respData = new(GetSeasonsResponse)
	err = json.Unmarshal(data, respData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode response data to structure")
	}

	return respData.Seasons, nil

}
