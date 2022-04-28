resource "random_password" "rds_master_password" {
  length  = 16
  special = false
}

resource "aws_ssm_parameter" "db_pass" {
  name        = "/database/pub_poker_root_password"
  description = "Aurora Mysql Root Password"
  type        = "SecureString"
  overwrite   = true
  value       = random_password.rds_master_password.result
}

resource "aws_rds_cluster" "pub_serverless_aurora" {
  cluster_identifier      = "pub-serverless-aurora"
  engine                  = "aurora-mysql"
  engine_mode             = "serverless"
  database_name           = "pubpoker"
  enable_http_endpoint    = false
  master_username         = "pokeradmin"
  master_password         = random_password.rds_master_password.result
  backup_retention_period = 1

  vpc_security_group_ids = [
    aws_security_group.allow_rds_connections.id
  ]

  db_subnet_group_name = aws_db_subnet_group.poker.name

  skip_final_snapshot = true

  scaling_configuration {
    auto_pause               = true
    min_capacity             = 1
    max_capacity             = 2
    seconds_until_auto_pause = 300
    timeout_action           = "ForceApplyCapacityChange"
  }
}

resource "aws_db_subnet_group" "poker" {
  name       = "pub_poker_db_subnet"
  subnet_ids = [aws_subnet.db1b.id, aws_subnet.db1d.id]
}

