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

func (h *handler) handleGetTournament(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{}, nil
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

	_ = &handler{
		logger:      logger,
		leaderboard: lbc,
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
			Method: http.MethodGet,
			Path:   "/tournaments",
		}: nil,
	}

	lambda.Start(apigw.HandleRoutes(routes))

}
