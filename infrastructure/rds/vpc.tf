data "aws_vpc" "main"{
    filter {
      name = "tag:Name"
      values = ["*vpc*"]
    }
}

data "aws_subnets" "subnets"{
    filter {
      name = "tag:Name"
      values = ["*rds*"]
    }
}