resource "aws_acm_certificate" "certificate" {
  domain_name       = aws_route53_zone.zone.name
  validation_method = "DNS"
  lifecycle {
    create_before_destroy = true
  }
}

# resource "aws_acm_certificate_validation" "certificate" {
#   certificate_arn         = aws_acm_certificate.certificate.arn
#   validation_record_fqdns = [for record in aws_route53_record.acm_validation_records : record.fqdn]
# }
