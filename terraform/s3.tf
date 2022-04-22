resource "aws_s3_bucket" "ppc_lambda_functions" {
  bucket = "ppc-lambda-functions"
}

resource "aws_s3_bucket" "ppc_terraform_state" {
  bucket = "ppc-terraform-state"

  lifecycle {
    prevent_destroy = true
  }
}
