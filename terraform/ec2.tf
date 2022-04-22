locals {
  default_ami = "ami-04505e74c0741db8d"
}

resource "aws_network_interface" "ssh_ni" {
  subnet_id = aws_subnet.public1a.id
  security_groups = [
    aws_security_group.allow_ssh_from_home.id
  ]

}

data "aws_key_pair" "default" {
  key_name = "ddouglas-20220409"
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

resource "aws_security_group" "allow_ssh_from_home" {
  name        = "allow_ssh_from_home"
  description = "Allows SSH From Home (Created With Terraform)"
  vpc_id      = aws_vpc.poker.id

  ingress {
    description = "SSH from Home"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["208.104.177.231/32"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
