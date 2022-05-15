resource "aws_apigatewayv2_route" "get_venues" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "GET /venues"

  target = "integrations/${aws_apigatewayv2_integration.venues_handler.id}"
}


resource "aws_apigatewayv2_integration" "venues_handler" {
  api_id = aws_apigatewayv2_api.http.id

  connection_type        = "INTERNET"
  integration_type       = "AWS_PROXY"
  integration_method     = "POST"
  payload_format_version = "2.0"
  integration_uri        = aws_lambda_function.venues_handler.arn
  passthrough_behavior   = "WHEN_NO_MATCH"

}

resource "aws_lambda_permission" "allow_apig_venues_handler" {
  statement_id  = "AllowAPIGatewaySeasonLeaderboardHandlerInvocation"
  function_name = aws_lambda_function.venues_handler.function_name
  principal     = "apigateway.amazonaws.com"
  action        = "lambda:InvokeFunction"

  source_arn = "${aws_apigatewayv2_api.http.execution_arn}/*/*/venues"
}


