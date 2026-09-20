output "vpc_name" {
    value = data.aws_vpc.main_vpc.tags
}

output "pub_sub"{
    value = data.aws_subnets.pub_subnet.ids
}

output "priv_sub"{
    value = data.aws_subnets.priv_subnet.ids
}

output "ssl" {
    value = data.aws_acm_certificate.ssl.domain
}

output "frontend_url" {
  value = var.frontend_url
}

output "database_url" {
  value = var.db_url
}