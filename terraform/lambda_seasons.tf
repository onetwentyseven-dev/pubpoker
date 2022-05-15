resource "aws_lambda_function" "seasons_handler" {
  function_name = "seasons_handler"
  role          = aws_iam_role.lambda_execution_role.arn
  handler       = "seasons"
  s3_bucket     = aws_s3_bucket.lambda_functions.bucket
  s3_key        = "seasons_handler.zip"


  runtime = "go1.x"

  environment {
    variables = {
      SSM_PREFIX = "/pubpoker"
    }
  }

  lifecycle {
    ignore_changes = [
      filename,
      s3_bucket,
      s3_key,
      s3_object_version,
      source_code_hash,
      version,
      qualified_arn,
      last_modified,
      package_type,
      image_uri,
      source_code_size
    ]
  }
}


