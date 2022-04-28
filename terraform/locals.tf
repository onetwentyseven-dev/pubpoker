locals {
  ssm_prefix               = "/pubpoker"
  base_domain              = "ppc.onetwentyseven.dev"
  lambda_vpc_access_policy = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
}
