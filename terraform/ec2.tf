# resource "aws_security_group" "allow_lambda_egress" {
#   name        = "AllowLambdaEgress"
#   description = "Allows Lambda to Egress to anywhere on any port"
#   vpc_id      = aws_vpc.poker.id

#   egress {
#     from_port   = 0
#     to_port     = 0
#     protocol    = "-1"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
# }
