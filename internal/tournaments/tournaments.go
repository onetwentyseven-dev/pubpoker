package tournaments

import (
	"context"
	"database/sql"

	"github.com/gofrs/uuid"
	"github.com/onetwentyseven-dev/pubpoker"
	"github.com/onetwentyseven-dev/pubpoker/internal/db"
	"github.com/onetwentyseven-dev/pubpoker/internal/leaderboard"
	"github.com/pkg/errors"
)

type Service struct {
	leaderboard *leaderboard.Client
	db.TournamentRepositoryAPI
}

type API interface {
	Tournament(ctx context.Context, tournamentID uuid.UUID) (*pubpoker.Tournament, error)
	TournamentPlayers(ctx context.Context, tournamentID uuid.UUID) ([]*pubpoker.TournamentPlayer, error)
	TournamentVenue(ctx context.Context, tournamentID uuid.UUID) (*leaderboard.Venue, error)
	Tournaments(ctx context.Context) ([]*pubpoker.Tournament, error)
	CreateTournament(ctx context.Context, venueID uuid.UUID) (*pubpoker.Tournament, error)
	CreateTournamentPlayer(ctx context.Context, tournamentID, playerID uuid.UUID) (*pubpoker.TournamentPlayer, error)
	UpdateTournament(ctx context.Context, tournamentID uuid.UUID, fields *pubpoker.TournamentPatchFields) (*pubpoker.Tournament, error)
	UpdateTournamentPlayerPosition(ctx context.Context, tournamentID, playerID uuid.UUID, position uint) (*pubpoker.TournamentPlayer, error)
}

func NewTournamentService(lbc *leaderboard.Client, tournamentRepo db.TournamentRepositoryAPI) API {
	return &Service{
		leaderboard:             lbc,
		TournamentRepositoryAPI: tournamentRepo,
	}
}

func (s *Service) Tournaments(ctx context.Context) ([]*pubpoker.Tournament, error) {
	return s.TournamentRepositoryAPI.SelectTournaments(ctx)
}

func (s *Service) UpdateTournament(ctx context.Context, tournamentID uuid.UUID, fields *pubpoker.TournamentPatchFields) (*pubpoker.Tournament, error) {

	tournament, err := s.TournamentRepositoryAPI.SelectTournament(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	tournament.IsRegistrationClosed = fields.IsRegistrationClosed

	err = s.TournamentRepositoryAPI.UpdateTournament(ctx, tournament)

	return tournament, err

}

func (s *Service) TournamentPlayers(ctx context.Context, tournamentID uuid.UUID) ([]*pubpoker.TournamentPlayer, error) {

	_, err := s.TournamentRepositoryAPI.SelectTournament(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	players, err := s.TournamentRepositoryAPI.SelectTournamentPlayers(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	return players, err
}

func (s *Service) TournamentVenue(ctx context.Context, tournamentID uuid.UUID) (*leaderboard.Venue, error) {

	tournament, err := s.TournamentRepositoryAPI.SelectTournament(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	venue, err := s.leaderboard.Venue(ctx, tournament.VenueID)
	if err != nil {
		return nil, err
	}

	return venue, err
}

func (s *Service) Tournament(ctx context.Context, tournamentID uuid.UUID) (*pubpoker.Tournament, error) {

	tournament, err := s.TournamentRepositoryAPI.SelectTournament(ctx, tournamentID)
	if err != nil {
		return nil, err
	}

	tournament.Venue, err = s.leaderboard.Venue(ctx, tournament.VenueID)
	if err != nil {
		return nil, err
	}

	return tournament, err
}

func (s *Service) CreateTournament(ctx context.Context, venueID uuid.UUID) (*pubpoker.Tournament, error) {

	_, err := s.leaderboard.Venue(ctx, venueID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch venue by id")
	}

	tournament := &pubpoker.Tournament{
		ID:      uuid.Must(uuid.NewV4()),
		VenueID: venueID,
	}

	err = s.TournamentRepositoryAPI.CreateTournament(ctx, tournament)

	return tournament, err

}

func (s *Service) CreateTournamentPlayer(ctx context.Context, tournamentID, playerID uuid.UUID) (*pubpoker.TournamentPlayer, error) {

	lbPlayer, err := s.leaderboard.Player(ctx, playerID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate player")
	}

	tournament, err := s.TournamentRepositoryAPI.SelectTournament(ctx, tournamentID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch tournment with provided id")
	}

	player := &pubpoker.TournamentPlayer{
		TournamentID: tournament.ID,
		PlayerID:     playerID,
		Name:         lbPlayer.Name,
	}

	return player, s.TournamentRepositoryAPI.CreateTournamentPlayer(ctx, player)

}

func (s *Service) UpdateTournamentPlayerPosition(ctx context.Context, tournamentID, playerID uuid.UUID, position uint) (*pubpoker.TournamentPlayer, error) {

	_, err := s.TournamentRepositoryAPI.SelectTournament(ctx, tournamentID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch tournment with provided id")
	}

	_, err = s.leaderboard.Player(ctx, playerID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate player")
	}

	player, err := s.TournamentRepositoryAPI.SelectTournamentPlayer(ctx, tournamentID, playerID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch tournament player from DB")
	}

	if position > 0 && position <= 10 {
		_, err := s.TournamentRepositoryAPI.SelectTournamentPlayerByPosition(ctx, tournamentID, position)
		if err == nil {
			return nil, errors.Errorf("this tournament already has a player with a final position of %d", position)
		}

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "failed to validate existing positions in this tournament")
		}

		// At this point Position does not exist in the tournament
	}

	player.FinalPosition = position

	err = s.TournamentRepositoryAPI.UpdateTournamentPlayer(ctx, tournamentID, playerID, player)

	return player, err

}
