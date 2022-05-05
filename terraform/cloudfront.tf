resource "aws_cloudfront_origin_access_identity" "root" {
  comment = "Pub Poker Static Assets"
}

resource "aws_cloudfront_distribution" "ppc_main" {
  origin {
    domain_name = aws_s3_bucket.ppc_main_site.bucket_regional_domain_name
    origin_id   = aws_s3_bucket.ppc_main_site.id

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.root.cloudfront_access_identity_path
    }
  }

  enabled             = true
  is_ipv6_enabled     = true
  comment             = "Pub Poker Main Site Static Assets"
  default_root_object = "index.html"

  custom_error_response {
    error_code         = 403
    response_code      = 200
    response_page_path = "/index.html"
  }

  custom_error_response {
    error_code         = 404
    response_code      = 200
    response_page_path = "/index.html"
  }

  aliases = [local.base_domain]

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = aws_s3_bucket.ppc_main_site.id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  price_class = "PriceClass_100"

  restrictions {
    geo_restriction {
      restriction_type = "none"
      locations        = []
    }

  }

  viewer_certificate {
    acm_certificate_arn      = aws_acm_certificate.ppc_certificate.arn
    minimum_protocol_version = "TLSv1.2_2021"
    ssl_support_method       = "sni-only"
  }

}

resource "aws_cloudfront_distribution" "ppc_play" {
  origin {
    domain_name = aws_s3_bucket.ppc_play_site.bucket_regional_domain_name
    origin_id   = aws_s3_bucket.ppc_play_site.id

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.root.cloudfront_access_identity_path
    }
  }

  enabled             = true
  is_ipv6_enabled     = true
  comment             = "Pub Poker Play Site Static Assets"
  default_root_object = "index.html"

  custom_error_response {
    error_code         = 403
    response_code      = 200
    response_page_path = "/index.html"
  }

  custom_error_response {
    error_code         = 404
    response_code      = 200
    response_page_path = "/index.html"
  }

  aliases = ["play.${local.base_domain}"]

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = aws_s3_bucket.ppc_play_site.id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  price_class = "PriceClass_100"

  restrictions {
    geo_restriction {
      restriction_type = "none"
      locations        = []
    }

  }

  viewer_certificate {
    acm_certificate_arn      = aws_acm_certificate.ppc_certificate.arn
    minimum_protocol_version = "TLSv1.2_2021"
    ssl_support_method       = "sni-only"
  }

}
