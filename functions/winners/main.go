package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/onetwentyseven-dev/pubpoker/internal/apigw/rest"
	"github.com/onetwentyseven-dev/pubpoker/internal/leaderboard"
	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

type handler struct {
	logger      *logrus.Logger
	leaderboard *leaderboard.Client
}

func (h *handler) handleGetWinners(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	rows, err := h.leaderboard.SeasonRecentWinners(ctx)
	if err != nil {
		fmt.Println(err)
		return rest.RespondJSONError(http.StatusBadRequest, errors.Wrap(err, "failed to fetch season leaderboard").Error(), nil, err)
	}

	return rest.RespondJSON(http.StatusOK, rows, map[string]string{})

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
			Path:   "/winners",
		}: h.handleGetWinners,
	}

	lambda.Start(rest.UseMiddleware(rest.HandleRoutes(routes), rest.Cors(rest.DefaultCorsOpt)))

}
