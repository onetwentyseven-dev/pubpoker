locals {
  default_ami = "ami-04505e74c0741db8d"
}

resource "aws_network_interface" "ssh_ni" {
  subnet_id = element(aws_subnet.dmz.*.id, 0)
  security_groups = [
    aws_security_group.allow_external_ssh.id
  ]

}

resource "aws_instance" "ssh_tunnel" {
  count         = var.enable_bastion ? 1 : 0
  ami           = local.default_ami
  instance_type = "t2.micro"

  network_interface {
    network_interface_id = aws_network_interface.ssh_ni.id
    device_index         = 0
  }

  key_name = data.aws_key_pair.default.key_name

}

resource "aws_security_group" "allow_external_ssh" {
  name        = "AllowExternalSSH"
  description = "Allow SSH Connection from External Sources"
  vpc_id      = aws_vpc.poker.id

  ingress {
    description = "SSH from Home"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["208.104.177.231/32", "66.191.182.17/32"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "allow_lambda_egress" {
  name        = "AllowLambdaEgress"
  description = "Allows Lambda to Egress to anywhere on any port"
  vpc_id      = aws_vpc.poker.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "allow_rds_connections" {
  name        = "AllowRDSConnections"
  description = "AllowRDSConnections"
  vpc_id      = aws_vpc.poker.id

  ingress {
    description = "MySQL TCP"
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    security_groups = [
      aws_security_group.allow_lambda_egress.id,
      aws_security_group.allow_external_ssh.id,
    ]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
