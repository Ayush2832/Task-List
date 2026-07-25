variable "region" {
  description = "Region where our frontend will deploy"
}

variable "bucket_name" {
  description = "Bucket name where we put our frontend code"
}

variable "oac_name" {
  description = "OAC name"
}

variable "certificate_arn" {
  type = string
}