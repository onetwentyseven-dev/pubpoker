resource "aws_vpc" "poker" {
  instance_tenancy     = "default"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    "Name" = "Pub Poker Championship"
  }
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.poker.id
}

resource "aws_route_table" "pubilc" {
  vpc_id = aws_vpc.poker.id
}

resource "aws_route" "default_gateway" {
  route_table_id         = aws_route_table.pubilc.id
  gateway_id             = aws_internet_gateway.igw.id
  destination_cidr_block = "0.0.0.0/0"
}

resource "aws_subnet" "public1a" {
  vpc_id                  = aws_vpc.poker.id
  availability_zone       = "us-east-1a"
  cidr_block              = "10.0.0.0/28" // 16 Addresses
  map_public_ip_on_launch = true
}

resource "aws_subnet" "public1c" {
  vpc_id                  = aws_vpc.poker.id
  availability_zone       = "us-east-1a"
  cidr_block              = "10.0.0.16/28" // 16 Addresses
  map_public_ip_on_launch = true
}


resource "aws_route_table_association" "public1a" {
  subnet_id      = aws_subnet.public1a.id
  route_table_id = aws_route_table.pubilc.id
}

resource "aws_route_table_association" "public1c" {
  subnet_id      = aws_subnet.public1c.id
  route_table_id = aws_route_table.pubilc.id
}


// If we need private subnets and routes, add them here
resource "aws_route_table" "priv" {
  vpc_id = aws_vpc.poker.id
}

resource "aws_subnet" "db1b" {
  vpc_id            = aws_vpc.poker.id
  cidr_block        = "10.0.0.32/28" // 16 Addresses
  availability_zone = "us-east-1b"
}

resource "aws_subnet" "db1d" {
  vpc_id            = aws_vpc.poker.id
  cidr_block        = "10.0.0.48/28" // 16 Addresses
  availability_zone = "us-east-1d"
}

resource "aws_route_table_association" "db1b" {
  subnet_id      = aws_subnet.db1b.id
  route_table_id = aws_route_table.priv.id
}

resource "aws_route_table_association" "db1d" {
  subnet_id      = aws_subnet.db1d.id
  route_table_id = aws_route_table.priv.id
}
