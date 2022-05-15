resource "aws_apigatewayv2_route" "get_seasons" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "GET /seasons"

  target = "integrations/${aws_apigatewayv2_integration.seasons_handler.id}"
}

resource "aws_apigatewayv2_route" "get_current_season" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "GET /seasons/current"

  target = "integrations/${aws_apigatewayv2_integration.seasons_handler.id}"
}

resource "aws_apigatewayv2_route" "get_season_seasonID_leaderboard" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "GET /seasons/{seasonID}/leaderboard"

  target = "integrations/${aws_apigatewayv2_integration.seasons_handler.id}"
}

resource "aws_apigatewayv2_integration" "seasons_handler" {
  api_id = aws_apigatewayv2_api.http.id

  connection_type        = "INTERNET"
  integration_type       = "AWS_PROXY"
  integration_method     = "POST"
  payload_format_version = "2.0"
  integration_uri        = aws_lambda_function.seasons_handler.arn
  passthrough_behavior   = "WHEN_NO_MATCH"

}

resource "aws_lambda_permission" "allow_apig_seasons_handler" {
  statement_id  = "AllowAPIGatewaySeasonsHandlerInvocation"
  function_name = aws_lambda_function.seasons_handler.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.http.execution_arn}/*/*/seasons"
  action        = "lambda:InvokeFunction"
}

resource "aws_lambda_permission" "allow_apig_current_seasons_handler" {
  statement_id  = "AllowAPIGatewayCurrentSeasonHandlerInvocation"
  function_name = aws_lambda_function.seasons_handler.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.http.execution_arn}/*/*/seasons/current"
  action        = "lambda:InvokeFunction"
}
