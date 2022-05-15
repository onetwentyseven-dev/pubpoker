package main

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/onetwentyseven-dev/pubpoker/internal/apigw/rest"
	"github.com/onetwentyseven-dev/pubpoker/internal/leaderboard"

	"github.com/sirupsen/logrus"
)

type handler struct {
	logger      *logrus.Logger
	leaderboard *leaderboard.Client
}

func (h *handler) handleGetVenues(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	venues, err := h.leaderboard.Venues(ctx)
	if err != nil {
		return rest.RespondJSON(http.StatusBadRequest, map[string]string{
			"error": "failed to fetch venues",
		}, map[string]string{})
	}

	return rest.RespondJSON(http.StatusOK, venues, nil)
}

// func (h *handler) handleGetVenues(ctx context.Context, input events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

// 	fmt.Printf("%T\n", input)
// 	inputData, _ := json.Marshal(input)
// 	fmt.Println("input :: ", string(inputData))
// 	return events.APIGatewayV2HTTPResponse{
// 		StatusCode: http.StatusOK,
// 	}, nil

// }

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
			Path:   "/venues",
		}: h.handleGetVenues,
	}

	lambda.Start(rest.UseMiddleware(rest.HandleRoutes(routes), rest.Cors(rest.DefaultCorsOpt)))
}
