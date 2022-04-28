package main

import (
	"context"
	"net/http"
	"time"

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

	var routes = map[apigw.Route]apigw.Handler{}

	lambda.Start(apigw.UseMiddleware(apigw.HandleRoutes(routes), apigw.Cors(apigw.DefaultCorsOpt)))

}
