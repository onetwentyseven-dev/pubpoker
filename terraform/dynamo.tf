resource "aws_dynamodb_table" "ppc_terraform_state" {
  name         = "ppc-terraform-state"
  hash_key     = "LockID"
  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "LockID"
    type = "S"
  }
}
