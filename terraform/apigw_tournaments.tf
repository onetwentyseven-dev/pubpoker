
resource "aws_apigatewayv2_route" "post_tournaments" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "POST /tournaments"

  target = "integrations/${aws_apigatewayv2_integration.tournaments_handler.id}"
}

resource "aws_apigatewayv2_route" "options_tournaments" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "OPTIONS /tournaments"

  target = "integrations/${aws_apigatewayv2_integration.tournaments_handler.id}"
}

resource "aws_apigatewayv2_route" "get_tournaments" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "GET /tournaments"

  target = "integrations/${aws_apigatewayv2_integration.tournaments_handler.id}"
}

resource "aws_apigatewayv2_route" "get_tournament" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "GET /tournaments/{tournamentID}"

  target = "integrations/${aws_apigatewayv2_integration.tournaments_handler.id}"
}

resource "aws_apigatewayv2_route" "get_tournament_venue" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "GET /tournaments/{tournamentID}/venue"

  target = "integrations/${aws_apigatewayv2_integration.tournaments_handler.id}"
}

resource "aws_apigatewayv2_route" "patch_tournament" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "PATCH /tournaments/{tournamentID}"

  target = "integrations/${aws_apigatewayv2_integration.tournaments_handler.id}"
}

resource "aws_apigatewayv2_route" "options_tournament" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "OPTIONS /tournaments/{tournamentID}"

  target = "integrations/${aws_apigatewayv2_integration.tournaments_handler.id}"
}


resource "aws_apigatewayv2_route" "post_tournaments_players" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "POST /tournaments/{tournamentID}/players"

  target = "integrations/${aws_apigatewayv2_integration.tournaments_handler.id}"
}

resource "aws_apigatewayv2_route" "options_tournaments_players" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "OPTIONS /tournaments/{tournamentID}/players"

  target = "integrations/${aws_apigatewayv2_integration.tournaments_handler.id}"
}

resource "aws_apigatewayv2_route" "get_tournament_players" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "GET /tournaments/{tournamentID}/players"

  target = "integrations/${aws_apigatewayv2_integration.tournaments_handler.id}"
}

resource "aws_apigatewayv2_route" "patch_tournaments_player" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "PATCH /tournaments/{tournamentID}/players/{playerID}"

  target = "integrations/${aws_apigatewayv2_integration.tournaments_handler.id}"
}

resource "aws_apigatewayv2_route" "options_tournaments_player" {
  api_id    = aws_apigatewayv2_api.http.id
  route_key = "OPTIONS /tournaments/{tournamentID}/players/{playerID}"

  target = "integrations/${aws_apigatewayv2_integration.tournaments_handler.id}"
}

resource "aws_apigatewayv2_integration" "tournaments_handler" {
  api_id = aws_apigatewayv2_api.http.id

  connection_type        = "INTERNET"
  integration_type       = "AWS_PROXY"
  integration_method     = "POST"
  payload_format_version = "2.0"
  integration_uri        = aws_lambda_function.tournaments_handler.arn
  passthrough_behavior   = "WHEN_NO_MATCH"

}

resource "aws_lambda_permission" "allow_apig_tournament_handler" {
  statement_id  = "AllowAPIGatewayHandlerInvocation"
  function_name = aws_lambda_function.tournaments_handler.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.http.execution_arn}/*/*/tournaments*"
  action        = "lambda:InvokeFunction"
}
