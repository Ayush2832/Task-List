output "vpc_name_id"{
    value = data.aws_vpc.main.id
}

output "vpc_name"{
    value = data.aws_vpc.main.tags
}

output "private_subnet_ids" {
  value = data.aws_subnets.subnets.ids
}