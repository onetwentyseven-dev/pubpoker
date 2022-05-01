package db

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/onetwentyseven-dev/pubpoker"
	"github.com/pkg/errors"
)

type TournamentRepository struct {
	db *sqlx.DB
}

type TournamentRepositoryAPI interface {
	SelectTournament(ctx context.Context, tournamentID uuid.UUID) (*pubpoker.Tournament, error)
	SelectTournaments(ctx context.Context) ([]*pubpoker.Tournament, error)
	SelectTournamentPlayer(ctx context.Context, tournamentID, playerID uuid.UUID) (*pubpoker.TournamentPlayer, error)
	SelectTournamentPlayerByPosition(ctx context.Context, tournamentID uuid.UUID, position uint) (*pubpoker.TournamentPlayer, error)
	SelectTournamentPlayers(ctx context.Context, tournamentID uuid.UUID) ([]*pubpoker.TournamentPlayer, error)
	CreateTournament(ctx context.Context, tournament *pubpoker.Tournament) error
	CreateTournamentPlayer(ctx context.Context, player *pubpoker.TournamentPlayer) error
	UpdateTournament(ctx context.Context, tournament *pubpoker.Tournament) error
	UpdateTournamentPlayer(ctx context.Context, tournamentID, playerID uuid.UUID, player *pubpoker.TournamentPlayer) error
}

var _ TournamentRepositoryAPI = (*TournamentRepository)(nil)

const (
	tournamentTable        = "tournaments"
	tournamentPlayersTable = "tournament_players"
)

var (
	tournamentTableColumns = []string{
		"id", "tournamentID", "venueID", "approved",
		"approvedByUserID", "isRegistrationClosed",
		"createdAtDate", "createdAt",
	}
	tournamentPlayersTableColumn = []string{
		"tournamentID", "playerID", "name",
		"finalPosition", "createdAtDate", "createdAt",
	}
)

func NewTournamentRepository(db *sqlx.DB) TournamentRepositoryAPI {
	return &TournamentRepository{
		db: db,
	}
}

func (r *TournamentRepository) SelectTournaments(ctx context.Context) ([]*pubpoker.Tournament, error) {

	query, _, err := sq.Select(tournamentTableColumns...).From(tournamentTable).OrderBy("createdAtDate DESC").ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate sql for fetch tournaments by id query")
	}

	var results = make([]*pubpoker.Tournament, 0)
	err = r.db.SelectContext(ctx, &results, query)

	return results, err

}

func (r *TournamentRepository) SelectTournament(ctx context.Context, tournamentID uuid.UUID) (*pubpoker.Tournament, error) {

	query, args, err := sq.Select(tournamentTableColumns...).
		From(tournamentTable).
		Where(sq.Eq{"id": tournamentID}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate sql for fetch tournaments by id query")
	}

	var result = new(pubpoker.Tournament)
	err = r.db.GetContext(ctx, result, query, args...)

	return result, err

}

func (r *TournamentRepository) SelectTournamentPlayer(ctx context.Context, tournamentID, playerID uuid.UUID) (*pubpoker.TournamentPlayer, error) {

	query, args, err := sq.Select(tournamentPlayersTableColumn...).
		From(tournamentPlayersTable).
		Where(sq.Eq{"tournamentID": tournamentID, "playerID": playerID}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate sql for fetch tournaments by id query")
	}

	var result = new(pubpoker.TournamentPlayer)
	err = r.db.GetContext(ctx, result, query, args...)

	return result, err

}

func (r *TournamentRepository) SelectTournamentPlayerByPosition(ctx context.Context, tournamentID uuid.UUID, position uint) (*pubpoker.TournamentPlayer, error) {

	query, args, err := sq.Select(tournamentPlayersTableColumn...).
		From(tournamentPlayersTable).
		Where(sq.Eq{"tournamentID": tournamentID, "finalPosition": position}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate sql for fetch tournaments by id query")
	}

	var result = new(pubpoker.TournamentPlayer)
	err = r.db.GetContext(ctx, result, query, args...)

	return result, err

}

func (r *TournamentRepository) CreateTournament(ctx context.Context, tournament *pubpoker.Tournament) error {

	now := time.Now()
	tournament.CreatedAt = now
	tournament.CreatedAtDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	query, args, err := sq.Insert("tournaments").SetMap(map[string]interface{}{
		"id":               tournament.ID,
		"tournamentID":     tournament.TournamentID,
		"venueID":          tournament.VenueID,
		"approved":         tournament.Approved,
		"approvedByUserID": tournament.ApprovedByUserID,
		"createdAtDate":    tournament.CreatedAtDate,
		"createdAt":        tournament.CreatedAt,
	}).ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to generate sql for insert tournaments query")
	}

	_, err = r.db.ExecContext(ctx, query, args...)

	return errors.Wrap(err, "failed to exec insert tournament query")

}

func (r *TournamentRepository) UpdateTournament(ctx context.Context, tournament *pubpoker.Tournament) error {

	query, args, err := sq.Update("tournaments").SetMap(map[string]interface{}{
		"isRegistrationClosed": tournament.IsRegistrationClosed,
	}).Where(sq.Eq{"id": tournament.ID}).ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to generate sql for insert tournaments query")
	}

	_, err = r.db.ExecContext(ctx, query, args...)

	return errors.Wrap(err, "failed to exec insert tournament query")

}

func (r *TournamentRepository) SelectTournamentPlayers(ctx context.Context, tournamentID uuid.UUID) ([]*pubpoker.TournamentPlayer, error) {

	query, args, err := sq.Select(tournamentPlayersTableColumn...).
		From(tournamentPlayersTable).
		Where(sq.Eq{"tournamentID": tournamentID}).
		OrderBy("name ASC").
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate sql for select tournament players query")
	}

	var players = make([]*pubpoker.TournamentPlayer, 0)
	err = r.db.SelectContext(ctx, &players, query, args...)

	return players, err

}

func (r *TournamentRepository) CreateTournamentPlayer(ctx context.Context, player *pubpoker.TournamentPlayer) error {

	now := time.Now()
	player.CreatedAt = now
	player.CreatedAtDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	query, args, err := sq.Insert(tournamentPlayersTable).SetMap(map[string]interface{}{
		"tournamentID":  player.TournamentID,
		"playerID":      player.PlayerID,
		"name":          player.Name,
		"finalPosition": player.FinalPosition,
		"createdAtDate": player.CreatedAtDate,
		"createdAt":     player.CreatedAt,
	}).ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to generate sql for insert tournaments query")
	}

	_, err = r.db.ExecContext(ctx, query, args...)

	return errors.Wrap(err, "failed to exec insert tournament player query")

}

func (r *TournamentRepository) UpdateTournamentPlayer(ctx context.Context, tournamentID, playerID uuid.UUID, player *pubpoker.TournamentPlayer) error {

	query, args, err := sq.Update(tournamentPlayersTable).SetMap(map[string]interface{}{
		"finalPosition": player.FinalPosition,
	}).Where(sq.Eq{"tournamentID": tournamentID, "playerID": playerID}).ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to generate sql for insert tournaments query")
	}

	_, err = r.db.ExecContext(ctx, query, args...)

	return errors.Wrap(err, "failed to exec update tournament player query")

}
