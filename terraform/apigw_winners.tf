resource "aws_apigatewayv2_route" "get_winners" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "GET /winners"

  target = "integrations/${aws_apigatewayv2_integration.winners_handler.id}"
}

resource "aws_apigatewayv2_integration" "winners_handler" {
  api_id = aws_apigatewayv2_api.http.id

  connection_type        = "INTERNET"
  integration_type       = "AWS_PROXY"
  integration_method     = "POST"
  payload_format_version = "2.0"
  integration_uri        = aws_lambda_function.winners_handler.arn
  passthrough_behavior   = "WHEN_NO_MATCH"

}

