resource "aws_security_group" "load_balancer_sg" {
    name = "lb_sg"
    vpc_id = data.aws_vpc.main_vpc.id
}

resource "aws_vpc_security_group_ingress_rule" "allow_80_traffic" {
    security_group_id = aws_security_group.load_balancer_sg.id
    cidr_ipv4 = "0.0.0.0/0"
    from_port = "80"
    ip_protocol = "tcp"
    to_port = "80"
}

resource "aws_vpc_security_group_ingress_rule" "allow_443_traffic" {
    security_group_id = aws_security_group.load_balancer_sg.id
    ip_protocol = "tcp"
    cidr_ipv4 = "0.0.0.0/0"
    from_port = "443"
    to_port = "443"
}

resource "aws_vpc_security_group_egress_rule" "elb_allow_all_traffic_ipv4" {
    security_group_id = aws_security_group.load_balancer_sg.id
    cidr_ipv4         = "0.0.0.0/0"
    ip_protocol       = "-1"
}


resource "aws_security_group" "ecs_sg" {
    name = "ecs_sg"
    vpc_id = data.aws_vpc.main_vpc.id
}

resource "aws_vpc_security_group_ingress_rule" "ecs_lb_allow" {
    security_group_id = aws_security_group.ecs_sg.id
    ip_protocol = "tcp"
    referenced_security_group_id = aws_security_group.load_balancer_sg.id
    from_port = "8080"
    to_port = "8080"
}

resource "aws_vpc_security_group_egress_rule" "ecs_allow_all_traffic_ipv4" {
    security_group_id = aws_security_group.ecs_sg.id
    cidr_ipv4         = "0.0.0.0/0"
    ip_protocol       = "-1"
}