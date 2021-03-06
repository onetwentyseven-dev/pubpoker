##################################
######                      ######
######  Github Action OIDC  ######
######                      ######
##################################
resource "aws_iam_role" "github_actions_oidc" {
  name                 = "GithubActionOIDC"
  description          = "Role Used by Github OIDC Provider"
  max_session_duration = "3600"
  assume_role_policy   = data.aws_iam_policy_document.github_actions_oidc_assume_role.json
}

data "aws_iam_policy_document" "github_actions_oidc_assume_role" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]
    effect  = "Allow"
    condition {
      test     = "StringLike"
      values   = ["repo:onetwentyseven-dev/pubpoker:*"]
      variable = "token.actions.githubusercontent.com:sub"
    }

    principals {
      identifiers = [aws_iam_openid_connect_provider.github.arn]
      type        = "Federated"
    }
  }
}

resource "aws_iam_role_policy" "github_action_oidc_access" {
  name   = "github_action_oidc_access"
  role   = aws_iam_role.github_actions_oidc.id
  policy = data.aws_iam_policy_document.github_actions_odic.json
}

data "aws_iam_policy_document" "github_actions_odic" {
  statement {
    effect = "Allow"
    actions = [
      "s3:Put*",
      "s3:Get*",
      "s3:List*"
    ]
    resources = [
      aws_s3_bucket.main.arn,
      aws_s3_bucket.play.arn,
      "${aws_s3_bucket.main.arn}/*",
      "${aws_s3_bucket.play.arn}/*",
      "${aws_s3_bucket.lambda_functions.arn}/*",
      aws_s3_bucket.lambda_functions.arn,
    ]
  }
  # statement {
  #   effect = "Allow"
  #   actions = [
  #     "cloudfront:CreateInvalidation",
  #   ]
  #   resources = [
  #     aws_cloudfront_distribution.home.arn,
  #     aws_cloudfront_distribution.play.arn,

  #   ]
  # }
  statement {
    effect    = "Allow"
    actions   = ["cloudfront:ListDistributions"]
    resources = ["*"]
  }
  statement {
    effect = "Allow"
    actions = [
      "lambda:UpdateFunctionCode"
    ]
    resources = [
      aws_lambda_function.tournaments_handler.arn,
      aws_lambda_function.players_handler.arn,
      aws_lambda_function.seasons_handler.arn,
      aws_lambda_function.winners_handler.arn,
    ]
  }
}


resource "aws_iam_openid_connect_provider" "github" {
  client_id_list = concat(
    ["https://github.com/onetwentyseven-dev"],
    ["sts.amazonaws.com"]
  )

  thumbprint_list = ["6938fd4d98bab03faadb97b34396831e3780aea1"]
  url             = "https://token.actions.githubusercontent.com"
}


