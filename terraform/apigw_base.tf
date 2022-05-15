resource "aws_apigatewayv2_api" "http" {
  name          = "ppc-api"
  protocol_type = "HTTP"
}

# resource "aws_apigatewayv2_domain_name" "http" {
#   domain_name = format("api.%s", aws_route53_zone.zone.name)
#   domain_name_configuration {
#     certificate_arn = aws_acm_certificate.certificate.arn
#     endpoint_type   = "REGIONAL"
#     security_policy = "TLS_1_2"
#   }
# }

resource "aws_apigatewayv2_stage" "http" {
  name        = "$default"
  api_id      = aws_apigatewayv2_api.http.id
  auto_deploy = true

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.apigateway_access_logs.arn
    format = jsonencode({
      httpMethod              = "$context.httpMethod"
      integrationErrorMessage = "$context.integrationErrorMessage"
      ip                      = "$context.identity.sourceIp"
      protocol                = "$context.protocol"
      requestId               = "$context.requestId"
      requestTime             = "$context.requestTime"
      responseLength          = "$context.responseLength"
      routeKey                = "$context.routeKey"
      status                  = "$context.status"
    })
  }


  lifecycle {
    ignore_changes = [deployment_id]
  }
}

# resource "aws_apigatewayv2_api_mapping" "http" {
#   api_id      = aws_apigatewayv2_api.http.id
#   domain_name = aws_apigatewayv2_domain_name.http.id
#   stage       = aws_apigatewayv2_stage.http.id
# }
