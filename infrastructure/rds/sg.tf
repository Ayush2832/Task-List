resource "aws_security_group" "rds_sg" {
  name = "rds_sg"
  vpc_id = data.aws_vpc.main.id
}

resource "aws_vpc_security_group_ingress_rule" "allow_vpc_connection" {
    security_group_id = aws_security_group.rds_sg.id
    cidr_ipv4 = data.aws_vpc.main.cidr_block
    from_port         = 5432
    ip_protocol       = "tcp"
    to_port           = 5432
}

resource "aws_vpc_security_group_ingress_rule" "allow_vpn_connection" {
    security_group_id = aws_security_group.rds_sg.id
    cidr_ipv4 = "165.22.223.6/32"
    from_port         = 5432
    ip_protocol       = "tcp"
    to_port           = 5432
}

resource "aws_vpc_security_group_egress_rule" "rds_allow_all_traffic_ipv4" {
  security_group_id = aws_security_group.rds_sg.id
  cidr_ipv4         = "0.0.0.0/0"
  ip_protocol       = "-1"
}

