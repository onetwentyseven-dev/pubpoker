resource "aws_apigatewayv2_api" "ppc_api" {
  name          = "ppc-api"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_domain_name" "ppc_api" {
  domain_name = "api.${aws_route53_zone.ppc_zone.name}"
  domain_name_configuration {
    certificate_arn = aws_acm_certificate.ppc_certificate.arn
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"
  }
}

resource "aws_apigatewayv2_stage" "ppc_api" {
  name        = "$default"
  api_id      = aws_apigatewayv2_api.ppc_api.id
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

resource "aws_apigatewayv2_api_mapping" "api" {
  api_id      = aws_apigatewayv2_api.ppc_api.id
  domain_name = aws_apigatewayv2_domain_name.ppc_api.id
  stage       = aws_apigatewayv2_stage.ppc_api.id
}
