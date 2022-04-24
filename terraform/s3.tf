resource "aws_s3_bucket" "ppc_main_site" {
  bucket = "ppc-main-site"
}

data "aws_iam_policy_document" "ppc_static_assets" {
  statement {
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.ppc_main_site.arn}/*"]

    principals {
      type = "AWS"
      identifiers = [
        aws_cloudfront_origin_access_identity.root.iam_arn
      ]
    }
  }
}

resource "aws_s3_bucket_policy" "ppc_main_site" {
  bucket = aws_s3_bucket.ppc_main_site.id
  policy = data.aws_iam_policy_document.ppc_static_assets.json
}

resource "aws_s3_bucket" "ppc_lambda_functions" {
  bucket = "ppc-lambda-functions"
}

resource "aws_s3_bucket" "ppc_terraform_state" {
  bucket = "ppc-terraform-state"

  lifecycle {
    prevent_destroy = true
  }
}
