resource "aws_cloudwatch_log_group" "apigateway_access_logs" {
  name              = "/aws/apigatewayv2/logs"
  retention_in_days = 3
}

resource "aws_cloudwatch_log_group" "players_handler" {
  name = "/aws/lambda/${aws_lambda_function.players_handler.function_name}"
}

resource "aws_cloudwatch_log_group" "seasons_handler" {
  name = "/aws/lambda/${aws_lambda_function.seasons_handler.function_name}"
}

resource "aws_cloudwatch_log_group" "tournaments_handler" {
  name = "/aws/lambda/${aws_lambda_function.tournaments_handler.function_name}"
}

resource "aws_cloudwatch_log_group" "venues_handler" {
  name = "/aws/lambda/${aws_lambda_function.venues_handler.function_name}"
}

resource "aws_cloudwatch_log_group" "winners_handler" {
  name = "/aws/lambda/${aws_lambda_function.winners_handler.function_name}"
}
