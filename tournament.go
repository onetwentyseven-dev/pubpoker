package pubpoker

import (
	"time"

	"github.com/gofrs/uuid"
)

type Tournament struct {
	ID uuid.UUID `db:"id" json:"id"`
	// This is the Tournament ID provided by the PokerLeaderboard API and will be null until this tournment is approved and POST to their API
	TournamentID     uuid.NullUUID       `db:"tournamentID" json:"tournamentID"`
	VenueID          uuid.UUID           `db:"venueID" json:"venueID"`
	Approved         bool                `db:"approved" json:"approved"`
	ApprovedByUserID uuid.UUID           `db:"approvedByUserID" json:"approvedByUserID"`
	Players          []*TournamentPlayer `json:"players"`
	CreatedAt        time.Time           `db:"createdAt" json:"createdAt"`
}

type TournamentPlayer struct {
	ID uuid.UUID `db:"id" json:"id"`
	// This is our TournamentID and relates to the Tournment these player initially registered with
	TournamentID  uuid.UUID `db:"tournamentID" json:"tournamentID"`
	PlayerID      uuid.UUID `db:"playerID" json:"playerID"`
	FinalPosition uint      `db:"finalPosition" json:"finalPosition"`
	Name          string    `db:"name" json:"name"`
	CreatedAt     time.Time `db:"createdAt" json:"createdAt"`
}
