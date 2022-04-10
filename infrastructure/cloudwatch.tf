resource "aws_cloudwatch_log_group" "apigateway_access_logs" {
  name              = "/aws/apigatewayv2/logs"
  retention_in_days = 3
}
