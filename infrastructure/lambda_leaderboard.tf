resource "aws_lambda_function" "leaderboard_handler" {
  function_name = "leaderboard_handler"
  role          = aws_iam_role.lambda_execution_role.arn
  handler       = "leaderboard"
  s3_bucket     = aws_s3_bucket.ppc_lambda_functions.bucket
  s3_key        = "leaderboard_handler.zip"


  runtime = "go1.x"

  environment {
    variables = {
      SSM_PREFIX = "/pubpoker"
    }
  }
}

resource "aws_apigatewayv2_integration" "leaderboard_handler" {
  api_id = aws_apigatewayv2_api.ppc_api.id

  connection_type        = "INTERNET"
  integration_type       = "AWS_PROXY"
  integration_method     = "POST"
  payload_format_version = "2.0"
  integration_uri        = aws_lambda_function.leaderboard_handler.arn
  passthrough_behavior   = "WHEN_NO_MATCH"

}

resource "aws_lambda_permission" "allow_apig_leaderboard_handler" {
  statement_id  = "AllowAPIGatewaySeasonLeaderboardHandlerInvocation"
  function_name = aws_lambda_function.leaderboard_handler.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.ppc_api.execution_arn}/*/*/seasons/{seasonID}/leaderboard"
  action        = "lambda:InvokeFunction"
}
