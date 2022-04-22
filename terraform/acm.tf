resource "aws_acm_certificate" "ppc_certificate" {
  domain_name       = aws_route53_zone.ppc_zone.name
  validation_method = "DNS"
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate_validation" "ppc_certificate" {
  certificate_arn         = aws_acm_certificate.ppc_certificate.arn
  validation_record_fqdns = [for record in aws_route53_record.ppc_acm_validation_records : record.fqdn]
}
