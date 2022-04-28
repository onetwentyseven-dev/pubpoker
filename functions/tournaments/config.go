package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var envConfig struct {
	SSMPrefix string `envconfig:"SSM_PREFIX" required:"true"`
}

var ssmConfig struct {
	LeaderboardUsername string `ssm:"leaderboard/username" required:"true"`
	LeaderboardPassword string `ssm:"leaderboard/password" required:"true"`
	LeagueID            string `ssm:"leaderboard/leagueID" required:"true"`
	MySQLHost           string `ssm:"leaderboard/mysql_host" required:"true"`
	MySQLDB             string `ssm:"leaderboard/mysql_db" required:"true"`
	MySQLUsername       string `ssm:"leaderboard/mysql_username" required:"true"`
	MySQLPassword       string `ssm:"leaderboard/mysql_password" required:"true"`
}

func loadConfig(awsConfig aws.Config) {

	_ = godotenv.Load(".env")

	err := envconfig.Process("", &envConfig)
	if err != nil {
		panic(fmt.Sprintf("envconfig: %s", err))
	}

	// err = ssmconfig.Process(context.TODO(), awsConfig, envConfig.SSMPrefix, &ssmConfig)
	// if err != nil {
	// 	panic(fmt.Sprintf("ssmconfig: %s", err))
	// }
}
