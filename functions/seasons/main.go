package main

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/onetwentyseven-dev/pubpoker/internal/apigw"
	"github.com/onetwentyseven-dev/pubpoker/internal/leaderboard"

	"github.com/sirupsen/logrus"
)

type handler struct {
	logger      *logrus.Logger
	leaderboard *leaderboard.Client
}

func (h *handler) handleGetSeasons(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	seasons, err := h.leaderboard.Seasons(ctx)
	if err != nil {
		return apigw.RespondJSON(http.StatusBadRequest, map[string]string{
			"error": "failed to fetch players",
		}, map[string]string{})
	}
	return apigw.RespondJSON(http.StatusOK, seasons, nil)

}

func (h *handler) handleGetCurrentSeason(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	seasons, err := h.leaderboard.Seasons(ctx)
	if err != nil {
		return apigw.RespondJSON(http.StatusBadRequest, map[string]string{
			"error": "failed to fetch seasons",
		}, map[string]string{})
	}

	season, err := h.leaderboard.CurrentSeason(ctx, seasons)
	if err != nil {
		return apigw.RespondJSON(http.StatusBadRequest, map[string]string{
			"error": "failed to fetch current season",
		}, map[string]string{})
	}

	return apigw.RespondJSON(http.StatusOK, season, nil)

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

	var routes = map[apigw.Route]apigw.Handler{
		{
			Method: http.MethodGet,
			Path:   "/seasons",
		}: h.handleGetSeasons,
		{
			Method: http.MethodGet,
			Path:   "/seasons/current",
		}: h.handleGetCurrentSeason,
	}

	lambda.Start(apigw.UseMiddleware(apigw.HandleRoutes(routes), apigw.Cors(apigw.DefaultCorsOpt)))

}
