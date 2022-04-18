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
	"github.com/onetwentyseven-dev/pubpoker/internal/apigw"
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
		return apigw.RespondJSONError(http.StatusBadRequest, "failed to parse seasonID to valid uuid", nil, err)
	}

	rows, err := h.leaderboard.SeasonLeaderboard(ctx, 1, 25, seasonID.String(), input.QueryStringParameters["search"])
	if err != nil {
		fmt.Println(err)
		return apigw.RespondJSONError(http.StatusBadRequest, errors.Wrap(err, "failed to fetch season leaderboard").Error(), nil, err)
	}

	return apigw.RespondJSON(http.StatusOK, rows, map[string]string{})

}

func (h *handler) handleGetRecentWinners(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	rows, err := h.leaderboard.SeasonRecentWinners(ctx)
	if err != nil {
		fmt.Println(err)
		return apigw.RespondJSONError(http.StatusBadRequest, errors.Wrap(err, "failed to fetch season leaderboard").Error(), nil, err)
	}

	return apigw.RespondJSON(http.StatusOK, rows, map[string]string{})

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
			Path:   "/seasons/{seasonID}/leaderboard",
		}: h.handleGetLeaderboard,
		{
			Method: http.MethodGet,
			Path:   "/recent-winners",
		}: h.handleGetRecentWinners,
	}

	lambda.Start(apigw.UseMiddleware(apigw.HandleRoutes(routes), apigw.Cors(apigw.DefaultCorsOpt)))

}
