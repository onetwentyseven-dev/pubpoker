package leaderboard

import (
	"time"

	"github.com/volatiletech/null/v8"
)

type GetLeaderboardPlayers struct {
	Pagination *ResponsePagination   `json:"pagination"`
	Records    []*LeaderboardRecords `json:"records"`
}

type LeaderboardRecords struct {
	Wins           int         `json:"wins"`
	NumTournaments int         `json:"num_tournaments"`
	PointsAwarded  string      `json:"points_awarded"`
	AvgFinish      string      `json:"avg_finish"`
	AvgPoints      string      `json:"avg_points"`
	Ranking        int         `json:"ranking"`
	GeneratedAt    time.Time   `json:"generated_at"`
	PlayerID       string      `json:"player_id"`
	SeasonID       string      `json:"season_id"`
	VenueID        interface{} `json:"venue_id"`
	LeagueID       string      `json:"league_id"`
	RegionID       interface{} `json:"region_id"`
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	IsVerified     bool        `json:"is_verified"`
	AvatarURL      string      `json:"avatar_url"`
}

type GetPlayersResponse struct {
	Players    []*Player           `json:"players"`
	Pagination *ResponsePagination `json:"pagination"`
}

type ResponsePagination struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	RowCount  int `json:"rowCount"`
	PageCount int `json:"pageCount"`
}

type Player struct {
	ID                    string         `json:"id,omitempty"`
	LeagueID              string         `json:"league_id,omitempty"`
	UserID                string         `json:"user_id,omitempty"`
	Name                  string         `json:"name,omitempty"`
	Email                 string         `json:"email,omitempty"`
	IsActive              bool           `json:"is_active,omitempty"`
	Notes                 string         `json:"notes,omitempty"`
	MembershipType        string         `json:"membership_type,omitempty"`
	IsMembership          bool           `json:"is_membership,omitempty"`
	Is12MonthPackage      bool           `json:"is_12_month_package,omitempty"`
	IsFirst6MonthPackage  bool           `json:"is_first_6_month_package,omitempty"`
	IsSecond6MonthPackage bool           `json:"is_second_6_month_package,omitempty"`
	StartDate             null.Time      `json:"start_date,omitempty"`
	EndDate               null.Time      `json:"end_date,omitempty"`
	QualifiedDate         null.Time      `json:"qualified_date,omitempty"`
	IsInvitedMainEvent    bool           `json:"is_invited_main_event,omitempty"`
	NumTournaments        int            `json:"num_tournaments,omitempty"`
	IsVerified            bool           `json:"is_verified,omitempty"`
	CreatedAt             time.Time      `json:"created_at,omitempty"`
	UpdatedAt             time.Time      `json:"updated_at,omitempty"`
	PlayerSetting         *PlayerSetting `json:"PlayerSetting,omitempty"`
}

type GetSeasonsResponse struct {
	Seasons    []*Season           `json:"seasons"`
	Pagination *ResponsePagination `json:"pagination"`
}

type Season struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	LeagueID          string    `json:"league_id"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	IsActive          bool      `json:"is_active"`
	ExcludeFromPoints bool      `json:"exclude_from_points"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type GetRecentWinnersResponse struct {
	Records []*RecentWinnerRecord `json:"records"`
}

type RecentWinnerRecord struct {
	Player     Player     `json:"Player"`
	Tournament Tournament `json:"Tournament"`
}

type PlayerSetting struct {
	AvatarURL string `json:"avatar_url"`
}

type Venue struct {
	Name string `json:"name"`
}
type Tournament struct {
	SeasonID string `json:"season_id"`
	VenueID  string `json:"venue_id"`
	Venue    Venue  `json:"Venue"`
}
