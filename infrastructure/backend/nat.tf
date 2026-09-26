
resource "aws_eip" "nat" {
  domain = "vpc"

  tags = {
    Name = "task-list-nat-eip"
  }
}


resource "aws_nat_gateway" "main" {
  allocation_id = aws_eip.nat.id
  subnet_id     = data.aws_subnets.pub_subnet.ids[0]

  tags = {
    Name = "task-list-nat"
  }

  depends_on = [aws_eip.nat]
}

resource "aws_route" "private_nat" {
  route_table_id         = data.aws_route_table.priv_Rt.id
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = aws_nat_gateway.main.id
}
