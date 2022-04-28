# resource "aws_ssm_parameter" "mysql_host" {
#   name  = "/pubpoker/leaderboard/mysql_host"
#   type  = "String"
#   value = aws_rds_cluster.pub_serverless_aurora.endpoint
# }

# resource "aws_ssm_parameter" "mysql_db" {
#   name  = "/pubpoker/leaderboard/mysql_db"
#   type  = "String"
#   value = aws_rds_cluster.pub_serverless_aurora.database_name
# }

# resource "aws_ssm_parameter" "mysql_username" {
#   name  = "/pubpoker/leaderboard/mysql_username"
#   type  = "String"
#   value = aws_rds_cluster.pub_serverless_aurora.master_username
# }

# resource "aws_ssm_parameter" "mysql_password" {
#   name  = "/pubpoker/leaderboard/mysql_password"
#   type  = "SecureString"
#   value = aws_rds_cluster.pub_serverless_aurora.master_password
# }
