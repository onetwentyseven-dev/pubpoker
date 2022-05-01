package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/onetwentyseven-dev/pubpoker"
	"github.com/onetwentyseven-dev/pubpoker/internal/apigw"
	"github.com/onetwentyseven-dev/pubpoker/internal/db"
	"github.com/onetwentyseven-dev/pubpoker/internal/leaderboard"
	"github.com/onetwentyseven-dev/pubpoker/internal/tournaments"
	"github.com/sirupsen/logrus"
)

var (
	dbConn *sqlx.DB
	logger *logrus.Logger
)

type handler struct {
	logger      *logrus.Logger
	leaderboard *leaderboard.Client
	tournament  tournaments.API
}

type handlePatchTournamentPlayerRequestBody struct {
	FinalPosition uint `json:"finalPosition"`
}

func (h *handler) handlePatchTournamentPlayer(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	tournamentIDStr, ok := input.PathParameters["tournamentID"]
	if !ok {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to parse tournmentID from path", nil, nil)
	}

	playerIDStr, ok := input.PathParameters["playerID"]
	if !ok {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to parse playerID from path", nil, nil)
	}

	tournamentID, err := uuid.FromString(tournamentIDStr)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "tournamentID is not a valid uuid", nil, nil)
	}

	playerID, err := uuid.FromString(playerIDStr)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "playerID is not a valid uuid", nil, nil)
	}

	var body = new(handlePatchTournamentPlayerRequestBody)
	err = json.Unmarshal([]byte(input.Body), body)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to decode request", nil, nil)
	}

	player, err := h.tournament.UpdateTournamentPlayerPosition(ctx, tournamentID, playerID, body.FinalPosition)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to update player", nil, err)
	}

	return apigw.RespondJSON(http.StatusOK, player, nil)

}

type handlePostTournamentPlayerRequestBody struct {
	PlayerID uuid.UUID `json:"playerID"`
}

func (h *handler) handlePostTournamentPlayer(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	tournamentIDStr, ok := input.PathParameters["tournamentID"]
	if !ok {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to parse tournmentID from path", nil, nil)
	}

	tournementID, err := uuid.FromString(tournamentIDStr)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "tournamentID is not a valid uuid", nil, nil)
	}

	var body = new(handlePostTournamentPlayerRequestBody)
	err = json.Unmarshal([]byte(input.Body), body)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to decode request", nil, nil)
	}

	player, err := h.tournament.CreateTournamentPlayer(ctx, tournementID, body.PlayerID)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to register player with tournment", nil, err)
	}

	return apigw.RespondJSON(http.StatusOK, player, nil)

}

func (h *handler) handleGetTournament(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	tournamentIDStr, ok := input.PathParameters["tournamentID"]
	if !ok {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to parse tournmentID from path", nil, nil)
	}

	tournementID, err := uuid.FromString(tournamentIDStr)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "tournamentID is not a valid uuid", nil, nil)
	}

	tournament, err := h.tournament.Tournament(ctx, tournementID)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "no tournment found for supplied tournamentID", nil, err)
	}

	return apigw.RespondJSON(http.StatusOK, tournament, nil)

}

func (h *handler) handleGetTournamentPlayers(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	tournamentIDStr, ok := input.PathParameters["tournamentID"]
	if !ok {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to parse tournmentID from path", nil, nil)
	}

	tournementID, err := uuid.FromString(tournamentIDStr)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "tournamentID is not a valid uuid", nil, nil)
	}

	players, err := h.tournament.TournamentPlayers(ctx, tournementID)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "no tournment found for supplied tournamentID", nil, err)
	}

	return apigw.RespondJSON(http.StatusOK, players, nil)

}

func (h *handler) handleGetTournamentVenue(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	tournamentIDStr, ok := input.PathParameters["tournamentID"]
	if !ok {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to parse tournmentID from path", nil, nil)
	}

	tournementID, err := uuid.FromString(tournamentIDStr)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "tournamentID is not a valid uuid", nil, nil)
	}

	venue, err := h.tournament.TournamentVenue(ctx, tournementID)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "no tournment found for supplied tournamentID", nil, err)
	}

	return apigw.RespondJSON(http.StatusOK, venue, nil)

}

func (h *handler) handlePatchTournament(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	tournamentIDStr, ok := input.PathParameters["tournamentID"]
	if !ok {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to parse tournmentID from path", nil, nil)
	}

	tournementID, err := uuid.FromString(tournamentIDStr)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "tournamentID is not a valid uuid", nil, nil)
	}

	var body = new(pubpoker.TournamentPatchFields)
	err = json.Unmarshal([]byte(input.Body), body)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to decode request", nil, err)
	}

	tournament, err := h.tournament.UpdateTournament(ctx, tournementID, body)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "no tournment found for supplied tournamentID", nil, err)
	}

	return apigw.RespondJSON(http.StatusOK, tournament, nil)

}

func (h *handler) handleGetTournaments(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	tournaments, err := h.tournament.Tournaments(ctx)
	if err != nil {
		return apigw.RespondJSONError(http.StatusInternalServerError, "failed to fetch tournaments", nil, nil)
	}

	return apigw.RespondJSON(http.StatusOK, tournaments, nil)
}

type handlePostTournamentRequestBody struct {
	VenueID uuid.UUID `json:"venueID"`
}

func (h *handler) handlePostTournament(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	var body = new(handlePostTournamentRequestBody)
	err := json.Unmarshal([]byte(input.Body), body)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to decode request", nil, err)
	}

	tournament, err := h.tournament.CreateTournament(ctx, body.VenueID)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to create tournament", nil, err)
	}

	return apigw.RespondJSON(http.StatusCreated, tournament, nil)
}

func main() {

	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	apigw.SetLogger(logger)

	awsConf, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.WithError(err).Fatal("failed to load default config for aws sdk")
	}

	loadConfig(awsConf)
	loadMySQL()

	lbc := leaderboard.New(&http.Client{
		Timeout: time.Second * 5,
	}, leaderboard.Credentials{
		Username: ssmConfig.LeaderboardUsername,
		Password: ssmConfig.LeaderboardPassword,
	}, leaderboard.Config{
		LeagueID: ssmConfig.LeagueID,
	})

	tournamentRepo := db.NewTournamentRepository(dbConn)
	tournamentService := tournaments.NewTournamentService(lbc, tournamentRepo)

	h := &handler{
		logger:      logger,
		leaderboard: lbc,
		tournament:  tournamentService,
	}

	var routes = map[apigw.Route]apigw.Handler{
		{
			Method: http.MethodPatch,
			Path:   "/tournaments/{tournamentID}/players/{playerID}",
		}: h.handlePatchTournamentPlayer,
		{
			Method: http.MethodGet,
			Path:   "/tournaments/{tournamentID}",
		}: h.handleGetTournament,
		{
			Method: http.MethodGet,
			Path:   "/tournaments/{tournamentID}/players",
		}: h.handleGetTournamentPlayers,
		{
			Method: http.MethodGet,
			Path:   "/tournaments/{tournamentID}/venue",
		}: h.handleGetTournamentVenue,
		{
			Method: http.MethodPatch,
			Path:   "/tournaments/{tournamentID}",
		}: h.handlePatchTournament,
		{
			Method: http.MethodPost,
			Path:   "/tournaments/{tournamentID}/players",
		}: h.handlePostTournamentPlayer,
		{
			Method: http.MethodPost,
			Path:   "/tournaments",
		}: h.handlePostTournament,
		{
			Method: http.MethodGet,
			Path:   "/tournaments",
		}: h.handleGetTournaments,
	}

	lambda.Start(
		apigw.UseMiddleware(
			apigw.HandleRoutes(routes),
			apigw.Cors(apigw.DefaultCorsOpt),
			apigw.ContentType("application/json"),
		),
	)

}
