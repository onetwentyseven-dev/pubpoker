package apigw

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/sirupsen/logrus"
)

func Log(logger *logrus.Logger) Middleware {
	return func(h Handler) Handler {
		return func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
			return events.APIGatewayV2HTTPResponse{}, nil
		}
	}
}
