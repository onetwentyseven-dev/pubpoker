resource "aws_lambda_function" "seasons_handler" {
  function_name = "seasons_handler"
  role          = aws_iam_role.lambda_execution_role.arn
  handler       = "seasons"
  s3_bucket     = aws_s3_bucket.ppc_lambda_functions.bucket
  s3_key        = "seasons_handler.zip"


  runtime = "go1.x"

  environment {
    variables = {
      SSM_PREFIX = "/pubpoker"
    }
  }
}

resource "aws_apigatewayv2_integration" "seasons_handler" {
  api_id = aws_apigatewayv2_api.ppc_api.id

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
  source_arn    = "${aws_apigatewayv2_api.ppc_api.execution_arn}/*/*/seasons"
  action        = "lambda:InvokeFunction"
}

resource "aws_lambda_permission" "allow_apig_current_seasons_handler" {
  statement_id  = "AllowAPIGatewayCurrentSeasonHandlerInvocation"
  function_name = aws_lambda_function.seasons_handler.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.ppc_api.execution_arn}/*/*/seasons/current"
  action        = "lambda:InvokeFunction"
}
