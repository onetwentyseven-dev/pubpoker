resource "aws_vpc" "poker" {
  instance_tenancy     = "default"
  enable_dns_hostnames = true
  enable_dns_support   = true


  tags = {
    "Name" = "Pub Poker Championship"
  }
}

resource "aws_eip" "ngw" {
  count = 1
  vpc   = true
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.poker.id
}

resource "aws_nat_gateway" "ngw" {
  allocation_id = element(aws_eip.ngw.*.id, 0)
  subnet_id     = element(aws_subnet.dmz.*.id, 0)
  depends_on = [
    aws_internet_gateway.igw
  ]
}

# Rout Tables
resource "aws_route_table" "pub" {
  vpc_id = aws_vpc.poker.id
  tags = {
    "pub"  = "true",
    "priv" = "false"
  }
}

resource "aws_route" "pub_default_gateway" {
  route_table_id         = aws_route_table.pub.id
  gateway_id             = aws_internet_gateway.igw.id
  destination_cidr_block = "0.0.0.0/0"
}

resource "aws_route_table" "priv" {
  vpc_id = aws_vpc.poker.id
  tags = {
    "pub"  = "false",
    "priv" = "true"
  }
}

resource "aws_route" "priv_default_gateway" {
  destination_cidr_block = "0.0.0.0/0"
  route_table_id         = aws_route_table.priv.id
  nat_gateway_id         = aws_nat_gateway.ngw.id
}

resource "aws_subnet" "dmz" {
  count = var.az_count

  vpc_id                  = aws_vpc.poker.id
  cidr_block              = cidrsubnet(var.vpc_cidr, 4, count.index)
  availability_zone       = element(var.azs, count.index)
  map_public_ip_on_launch = true

  tags = {
    pub = "true"
  }

}

resource "aws_route_table_association" "pub_dmz" {
  count = var.az_count

  subnet_id      = element(aws_subnet.dmz.*.id, count.index)
  route_table_id = aws_route_table.pub.id
}

# DB's/Redis/Cache
resource "aws_subnet" "db" {
  count = var.az_count

  vpc_id                  = aws_vpc.poker.id
  cidr_block              = cidrsubnet(var.vpc_cidr, 4, 2 + count.index)
  availability_zone       = element(var.azs, count.index)
  map_public_ip_on_launch = false

  tags = {
    priv = "true"
  }

}

resource "aws_route_table_association" "priv_db" {
  count = var.az_count

  subnet_id      = element(aws_subnet.db.*.id, count.index)
  route_table_id = aws_route_table.priv.id
}

resource "aws_subnet" "app" {
  count = var.az_count

  vpc_id                  = aws_vpc.poker.id
  cidr_block              = cidrsubnet(var.vpc_cidr, 4, 4 + count.index)
  availability_zone       = element(var.azs, count.index)
  map_public_ip_on_launch = false

  tags = {
    priv = "true"
  }


}

resource "aws_route_table_association" "priv_app" {
  count = var.az_count

  subnet_id      = element(aws_subnet.app.*.id, count.index)
  route_table_id = aws_route_table.priv.id
}
