package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/gofrs/uuid"
	"github.com/onetwentyseven-dev/pubpoker/internal/apigw/rest"
	"github.com/onetwentyseven-dev/pubpoker/internal/leaderboard"
	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

type handler struct {
	logger      *logrus.Logger
	leaderboard *leaderboard.Client
}

func (h *handler) handleGetLeaderboard(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	seasonID, err := uuid.FromString(input.PathParameters["seasonID"])
	if err != nil {
		return rest.RespondJSONError(http.StatusBadRequest, "failed to parse seasonID to valid uuid", nil, err)
	}

	rows, err := h.leaderboard.SeasonLeaderboard(ctx, 1, 25, seasonID.String(), input.QueryStringParameters["search"])
	if err != nil {
		fmt.Println(err)
		return rest.RespondJSONError(http.StatusBadRequest, errors.Wrap(err, "failed to fetch season leaderboard").Error(), nil, err)
	}

	return rest.RespondJSON(http.StatusOK, rows, map[string]string{})

}

func (h *handler) handleGetSeasons(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	seasons, err := h.leaderboard.Seasons(ctx)
	if err != nil {
		return rest.RespondJSON(http.StatusBadRequest, map[string]string{
			"error": "failed to fetch seasons",
		}, map[string]string{})
	}
	return rest.RespondJSON(http.StatusOK, seasons, nil)

}

func (h *handler) handleGetCurrentSeason(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	seasons, err := h.leaderboard.Seasons(ctx)
	if err != nil {
		return rest.RespondJSON(http.StatusBadRequest, map[string]string{
			"error": "failed to fetch seasons",
		}, map[string]string{})
	}

	season, err := h.leaderboard.CurrentSeason(ctx, seasons)
	if err != nil {
		return rest.RespondJSON(http.StatusBadRequest, map[string]string{
			"error": "failed to fetch current season",
		}, map[string]string{})
	}

	return rest.RespondJSON(http.StatusOK, season, nil)

}

func main() {

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	awsConf, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.WithError(err).Fatal("failed to load default config for aws sdk")
	}

	loadConfig(awsConf)

	lbc := leaderboard.New(&http.Client{
		Timeout: time.Second * 5,
	}, leaderboard.Credentials{
		Username: ssmConfig.LeaderboardUsername,
		Password: ssmConfig.LeaderboardPassword,
	}, leaderboard.Config{
		LeagueID: ssmConfig.LeagueID,
	})

	h := &handler{
		logger:      logger,
		leaderboard: lbc,
	}

	var routes = map[rest.Route]rest.Handler{
		{
			Method: http.MethodGet,
			Path:   "/seasons",
		}: h.handleGetSeasons,
		{
			Method: http.MethodGet,
			Path:   "/seasons/current",
		}: h.handleGetCurrentSeason,
		{
			Method: http.MethodGet,
			Path:   "/seasons/{seasonID}/leaderboard",
		}: h.handleGetLeaderboard,
	}

	lambda.Start(rest.UseMiddleware(rest.HandleRoutes(routes), rest.Cors(rest.DefaultCorsOpt)))

}
