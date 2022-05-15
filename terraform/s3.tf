resource "aws_s3_bucket" "main" {
  bucket = "ppc-main-site"
}

resource "aws_s3_bucket" "play" {
  bucket = "ppc-play-site"
}

data "aws_iam_policy_document" "main_assets" {
  statement {
    actions = ["s3:GetObject"]
    resources = [
      format("%s/*", aws_s3_bucket.main.arn)
    ]

    principals {
      type = "AWS"
      identifiers = [
        aws_cloudfront_origin_access_identity.root.iam_arn
      ]
    }
  }
}

resource "aws_s3_bucket_policy" "main" {
  bucket = aws_s3_bucket.main.id
  policy = data.aws_iam_policy_document.main_assets.json
}


data "aws_iam_policy_document" "play_assets" {
  statement {
    actions = ["s3:GetObject"]
    resources = [
      format("%s/*", aws_s3_bucket.play.arn)
    ]

    principals {
      type = "AWS"
      identifiers = [
        aws_cloudfront_origin_access_identity.root.iam_arn
      ]
    }
  }
}

resource "aws_s3_bucket_policy" "play" {
  bucket = aws_s3_bucket.play.id
  policy = data.aws_iam_policy_document.play_assets.json
}


resource "aws_s3_bucket" "lambda_functions" {
  bucket = "ppc-lambda-functions"
}

resource "aws_s3_bucket" "ppc_terraform_state" {
  bucket = "ppc-terraform-state"

  lifecycle {
    prevent_destroy = true
  }
}
