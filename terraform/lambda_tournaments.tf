resource "aws_lambda_function" "tournaments_handler" {
  function_name = "tournaments_handler"
  role          = aws_iam_role.lambda_execution_role.arn
  handler       = "tournaments"
  s3_bucket     = aws_s3_bucket.ppc_lambda_functions.bucket
  s3_key        = "tournaments_handler.zip"
  vpc_config {
    subnet_ids = [
      aws_subnet.public1a.id,
      aws_subnet.public1c.id,
    ]
    security_group_ids = [
      aws_security_group.allow_lambda_egress.id
    ]
  }

  runtime = "go1.x"

  environment {
    variables = {
      SSM_PREFIX = "/pubpoker"
    }
  }
}

resource "aws_apigatewayv2_integration" "tournaments_handler" {
  api_id = aws_apigatewayv2_api.ppc_api.id

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
  source_arn    = "${aws_apigatewayv2_api.ppc_api.execution_arn}/*/*/tournaments*"
  action        = "lambda:InvokeFunction"
}
