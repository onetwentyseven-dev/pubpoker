resource "aws_route53_zone" "ppc_zone" {
  name = "ppc.onetwentyseven.dev"
}

resource "aws_route53_record" "ppc_api_a_record" {
  name    = aws_apigatewayv2_domain_name.ppc_api.domain_name
  type    = "A"
  zone_id = aws_route53_zone.ppc_zone.zone_id

  alias {
    name                   = aws_apigatewayv2_domain_name.ppc_api.domain_name_configuration[0].target_domain_name
    zone_id                = aws_apigatewayv2_domain_name.ppc_api.domain_name_configuration[0].hosted_zone_id
    evaluate_target_health = false
  }
}

resource "aws_route53_record" "ppc_acm_validation_records" {
  for_each = {
    for dvo in aws_acm_certificate.ppc_certificate.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
      type   = dvo.resource_record_type
    }
  }

  name    = each.value.name
  type    = each.value.type
  records = [each.value.record]
  zone_id = aws_route53_zone.ppc_zone.id
  ttl     = 60
}



resource "aws_route53_record" "ppc_main" {
  zone_id = aws_route53_zone.ppc_zone.id
  name    = local.base_domain
  type    = "A"

  alias {
    zone_id                = aws_cloudfront_distribution.ppc_main.hosted_zone_id
    name                   = aws_cloudfront_distribution.ppc_main.domain_name
    evaluate_target_health = false
  }
}
