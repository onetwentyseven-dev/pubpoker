package db

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/onetwentyseven-dev/pubpoker"
	"github.com/pkg/errors"
)

type TournamentRepository struct {
	db *sqlx.DB
}

type TournamentRepositoryAPI interface {
	GetTournamentByID(ctx context.Context, id string) (*pubpoker.Tournament, error)
	CreateTournament(ctx context.Context, tournament *pubpoker.Tournament) error
}

var _ TournamentRepositoryAPI = (*TournamentRepository)(nil)

func NewTournamentRepository(db *sqlx.DB) TournamentRepositoryAPI {
	return &TournamentRepository{
		db: db,
	}
}

func (r *TournamentRepository) GetTournamentByID(ctx context.Context, id string) (*pubpoker.Tournament, error) {

	query, args, err := sq.Select("id", "tournamentID", "venueID", "approved", "approvedByUserID", "createdAt").
		From("tournaments").
		Where(sq.Eq{"id": id}).
		OrderBy("createdAt desc").
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate sql for fetch tournaments by id query")
	}

	var results = new(pubpoker.Tournament)
	err = r.db.SelectContext(ctx, results, query, args)
	return results, err

}

func (r *TournamentRepository) CreateTournament(ctx context.Context, tournament *pubpoker.Tournament) error {

	tournament.CreatedAt = time.Now()

	query, args, err := sq.Insert("tournaments").SetMap(map[string]interface{}{
		"id":               tournament.ID,
		"tournamentID":     tournament.TournamentID,
		"venueID":          tournament.VenueID,
		"approved":         tournament.Approved,
		"approvedByUserID": tournament.ApprovedByUserID,
		"createdAt":        tournament.CreatedAt,
	}).ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to generate sql for insert tournaments query")
	}

	_, err = r.db.ExecContext(ctx, query, args...)

	return errors.Wrap(err, "failed to generate sql for insert tournaments query")

}
