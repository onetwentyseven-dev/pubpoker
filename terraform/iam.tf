resource "aws_iam_role" "lambda_execution_role" {
  name               = "lambda_execution_role"
  assume_role_policy = data.aws_iam_policy_document.lambda_execution_role_policy_doc.json
}

data "aws_iam_policy_document" "lambda_execution_role_policy_doc" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      identifiers = ["lambda.amazonaws.com"]
      type        = "Service"
    }
    effect = "Allow"
  }
}

# resource "aws_iam_role" "lambda_cloudwatch_access" {
#   name = "lambda_cloudwatch_access"
# }

resource "aws_iam_role_policy" "lambda_cloudwatch_access" {
  name   = "lambda_cloudwatch_access_policy"
  role   = aws_iam_role.lambda_execution_role.id
  policy = data.aws_iam_policy_document.lambda_write_cloudwatch_policy_doc.json
}

data "aws_iam_policy_document" "lambda_write_cloudwatch_policy_doc" {
  statement {
    effect  = "Allow"
    actions = ["logs:CreateLogStream", "logs:PutLogEvents"]
    resources = [
      "${aws_cloudwatch_log_group.players_handler.arn}:*",
      "${aws_cloudwatch_log_group.leaderboard_handler.arn}:*",
      "${aws_cloudwatch_log_group.seasons_handler.arn}:*",
    ]
  }
}

resource "aws_iam_role_policy" "lambda_ssm_read_pub_poker" {
  name   = "lambda_ssm_read_pub_poker"
  role   = aws_iam_role.lambda_execution_role.id
  policy = data.aws_iam_policy_document.lambda_read_pub_poker_ssm.json
}

data "aws_iam_policy_document" "lambda_read_pub_poker_ssm" {
  statement {
    effect  = "Allow"
    actions = ["ssm:GetParameters"]
    resources = [
      "arn:aws:ssm:us-east-1:847870459364:parameter/pubpoker",
      "arn:aws:ssm:us-east-1:847870459364:parameter/pubpoker/*"
    ]
  }
}
