resource "aws_apigatewayv2_route" "get_players_search" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "GET /players/search"

  target = "integrations/${aws_apigatewayv2_integration.players_handler.id}"
}

resource "aws_apigatewayv2_integration" "players_handler" {
  api_id = aws_apigatewayv2_api.http.id

  connection_type        = "INTERNET"
  integration_type       = "AWS_PROXY"
  integration_method     = "POST"
  payload_format_version = "2.0"
  integration_uri        = aws_lambda_function.players_handler.arn
  passthrough_behavior   = "WHEN_NO_MATCH"

}

resource "aws_lambda_permission" "allow_apig_players_search_handler" {
  statement_id  = "AllowAPIGatewayPlayersSearchHandlerInvocation"
  function_name = aws_lambda_function.players_handler.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.http.execution_arn}/*/*/players/search"
  action        = "lambda:InvokeFunction"
}

