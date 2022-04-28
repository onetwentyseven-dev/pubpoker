package tournaments

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/onetwentyseven-dev/pubpoker"
	"github.com/onetwentyseven-dev/pubpoker/internal/db"
	"github.com/onetwentyseven-dev/pubpoker/internal/leaderboard"
)

type Service struct {
	leaderboard *leaderboard.Client
	db.TournamentRepositoryAPI
}

type API interface {
	CreateTournament(ctx context.Context, venueID uuid.UUID) (*pubpoker.Tournament, error)
}

func NewTournamentService(lbc *leaderboard.Client, tournamentRepo db.TournamentRepositoryAPI) API {
	return &Service{
		leaderboard:             lbc,
		TournamentRepositoryAPI: tournamentRepo,
	}
}

func (s *Service) CreateTournament(ctx context.Context, venueID uuid.UUID) (*pubpoker.Tournament, error) {

	// _, err := s.leaderboard.Venue(ctx, venueID)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to fetch venue by id")
	// }

	tournament := &pubpoker.Tournament{
		ID:      uuid.Must(uuid.NewV4()),
		VenueID: venueID,
	}

	err := s.TournamentRepositoryAPI.CreateTournament(ctx, tournament)

	return tournament, err

}
