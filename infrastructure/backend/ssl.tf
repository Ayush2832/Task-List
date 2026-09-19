data "aws_acm_certificate" "ssl"{
    domain = "*.pingayush.in"
    statuses = ["ISSUED"]
    most_recent = true
}