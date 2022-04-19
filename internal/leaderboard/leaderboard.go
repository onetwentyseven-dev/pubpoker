package leaderboard

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

func (c *Client) SeasonLeaderboard(ctx context.Context, page, pageSize uint, seasonID, search string) ([]*LeaderboardRecords, error) {

	v := url.Values{}
	v.Set("season_id", seasonID)
	v.Set("page", strconv.Itoa(int(page)))
	v.Set("pageSize", strconv.Itoa(int(pageSize)))
	if search != "" {
		v.Set("search", search)
	}

	uri := c.apiURL()
	uri.Path = fmt.Sprintf("/p/league/%s/leaderboard", c.config.LeagueID)
	uri.RawQuery = v.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request to search players by name")
	}

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

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("failed to read response body")
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println(string(data))
		return nil, errors.Errorf("expected 200, get %d", res.StatusCode)
	}

	var respData = new(GetLeaderboardPlayers)
	err = json.Unmarshal(data, respData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode response data to structure")
	}

	return respData.Records, nil

}

const pubLeaderboardDefaultAvaterURI = "https://pokerleaderboards.com/assets/images/avatars/default_avatar.png"

func (c *Client) SeasonRecentWinners(ctx context.Context) ([]*RecentWinnerRecord, error) {

	uri := c.apiURL()
	uri.Path = fmt.Sprintf("/p/league/%s/recent-winners", c.config.LeagueID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request to search players by name")
	}

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

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("failed to read response body")
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println(string(data))
		return nil, errors.Errorf("expected 200, get %d", res.StatusCode)
	}

	var respData = new(GetRecentWinnersResponse)
	err = json.Unmarshal(data, respData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode response data to structure")
	}

	records := respData.Records
	for _, record := range records {
		pSettings := record.Player.PlayerSetting
		if pSettings.AvatarURL == "" {
			u := c.apiURL()
			u.Path = fmt.Sprintf("/p/player/%s/avatar", record.Player.ID)
			pSettings.AvatarURL = u.String()
		} else if pSettings.AvatarURL == "avatars/default_avatar.png" {
			pSettings.AvatarURL = pubLeaderboardDefaultAvaterURI
		}
	}

	return respData.Records, nil
}
