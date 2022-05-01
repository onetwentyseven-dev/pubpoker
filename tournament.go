package pubpoker

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/onetwentyseven-dev/pubpoker/internal/leaderboard"
)

type Tournament struct {
	ID uuid.UUID `db:"id" json:"id"`
	// This is the Tournament ID provided by the PokerLeaderboard API and will be null until this tournment is approved and POST'd to their API
	TournamentID         uuid.NullUUID `db:"tournamentID" json:"tournamentID"`
	VenueID              uuid.UUID     `db:"venueID" json:"venueID"`
	Approved             bool          `db:"approved" json:"approved"`
	ApprovedByUserID     uuid.NullUUID `db:"approvedByUserID" json:"approvedByUserID,omitempty"`
	IsRegistrationClosed bool          `db:"isRegistrationClosed" json:"isRegistrationClosed"`
	CreatedAtDate        time.Time     `db:"createdAtDate" json:"createdAtDate"`
	CreatedAt            time.Time     `db:"createdAt" json:"createdAt"`

	Players []*TournamentPlayer `json:"players"`
	Venue   *leaderboard.Venue  `json:"venue"`
}

type TournamentPatchFields struct {
	IsRegistrationClosed bool `json:"isRegistrationClosed"`
}

type TournamentPlayer struct {
	// This is our TournamentID and relates to the Tournment these player initially registered with
	TournamentID uuid.UUID `db:"tournamentID" json:"tournamentID"`
	// This is the PlayerID of the Player on Poker Leaderboard
	PlayerID      uuid.UUID `db:"playerID" json:"playerID"`
	Name          string    `db:"name" json:"name"`
	FinalPosition uint      `db:"finalPosition" json:"finalPosition"`
	CreatedAtDate time.Time `db:"createdAtDate" json:"createdAtDate"`
	CreatedAt     time.Time `db:"createdAt" json:"createdAt"`
}
