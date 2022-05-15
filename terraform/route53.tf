resource "aws_route53_zone" "zone" {
  name = "ppc.onetwentyseven.dev"
}

# resource "aws_route53_record" "api_a_record" {
#   name    = aws_apigatewayv2_domain_name.http.domain_name
#   type    = "A"
#   zone_id = aws_route53_zone.zone.zone_id

#   alias {
#     name                   = aws_apigatewayv2_domain_name.http.domain_name_configuration[0].target_domain_name
#     zone_id                = aws_apigatewayv2_domain_name.http.domain_name_configuration[0].hosted_zone_id
#     evaluate_target_health = false
#   }
# }

resource "aws_route53_record" "acm_validation_records" {
  for_each = {
    for dvo in aws_acm_certificate.certificate.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
      type   = dvo.resource_record_type
    }
  }

  name    = each.value.name
  type    = each.value.type
  records = [each.value.record]
  zone_id = aws_route53_zone.zone.id
  ttl     = 60
}

# resource "aws_route53_record" "home" {
#   zone_id = aws_route53_zone.zone.id
#   name    = local.base_domain
#   type    = "A"

#   alias {
#     zone_id                = aws_cloudfront_distribution.home.hosted_zone_id
#     name                   = aws_cloudfront_distribution.home.domain_name
#     evaluate_target_health = false
#   }
# }

# resource "aws_route53_record" "play" {
#   zone_id = aws_route53_zone.zone.id
#   name    = "play.${local.base_domain}"
#   type    = "A"

#   alias {
#     zone_id                = aws_cloudfront_distribution.play.hosted_zone_id
#     name                   = aws_cloudfront_distribution.play.domain_name
#     evaluate_target_health = false
#   }
# }

