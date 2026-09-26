data "aws_vpc" "main_vpc"{
    filter {
        name="tag:Name"
        values=["*vpc*"]
    }
}

data "aws_subnets" "pub_subnet"{
    filter{
        name="tag:Name"
        values=["*pub*"]
    }
}

data "aws_subnets" "priv_subnet"{
    filter{
        name="tag:Name"
        values=["*priv*"]
    }
}

data "aws_route_table" "priv_Rt" {
    filter {
      name = "tag:Name"
      values = ["priv-rt2"]
    }
}