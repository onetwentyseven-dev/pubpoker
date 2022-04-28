package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/onetwentyseven-dev/pubpoker/internal/apigw"
	"github.com/onetwentyseven-dev/pubpoker/internal/leaderboard"
	"github.com/onetwentyseven-dev/pubpoker/internal/tournaments"
	"github.com/sirupsen/logrus"
)

var (
	dbConn *sqlx.DB
)

type handler struct {
	logger      *logrus.Logger
	leaderboard *leaderboard.Client
	tournament  tournaments.API
}

type handlePostTournamentRequestBody struct {
	VenueID uuid.UUID `json:"venueID"`
}

// func (h *handler) handleGetTournament(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
// 	return events.APIGatewayV2HTTPResponse{}, nil
// }

func (h *handler) handlePostTournament(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	resp, err := http.Get("https://icanhazip.com/")
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to execute http get", nil, nil)
	}

	_, err = io.Copy(os.Stderr, resp.Body)
	if err != nil {
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to copy resp body to stderr", nil, nil)

	}

	// var body = new(handlePostTournamentRequestBody)
	// err := json.Unmarshal([]byte(input.Body), body)
	// if err != nil {
	// 	return apigw.RespondJSONError(http.StatusBadRequest, "failed to decode request", nil, nil)
	// }

	// tournament, err := h.

	return events.APIGatewayV2HTTPResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func main() {

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	awsConf, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.WithError(err).Fatal("failed to load default config for aws sdk")
	}

	loadConfig(awsConf)
	// loadMySQL()

	lbc := leaderboard.New(&http.Client{
		Timeout: time.Second * 5,
	}, leaderboard.Credentials{
		Username: ssmConfig.LeaderboardUsername,
		Password: ssmConfig.LeaderboardPassword,
	}, leaderboard.Config{
		LeagueID: ssmConfig.LeagueID,
	})

	// tournamentRepo := db.NewTournamentRepository(dbConn)
	// tournamentService := tournaments.NewTournamentService(lbc, tournamentRepo)

	h := &handler{
		logger:      logger,
		leaderboard: lbc,
		// tournament:  tournamentService,
	}

	var routes = map[apigw.Route]apigw.Handler{
		{
			Method: http.MethodGet,
			Path:   "/tournaments/{tournamentID}",
		}: nil,
		{
			Method: http.MethodPost,
			Path:   "/tournaments/{tournamentID}/players",
		}: nil,
		{
			Method: http.MethodPost,
			Path:   "/tournaments",
		}: h.handlePostTournament,
	}

	lambda.Start(
		apigw.UseMiddleware(
			apigw.HandleRoutes(routes),
			apigw.Cors(apigw.DefaultCorsOpt),
			apigw.ContentType("application/json"),
		),
	)

}
