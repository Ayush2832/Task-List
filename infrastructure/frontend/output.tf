output "bucket_name" {
  value = aws_s3_bucket.frontend.bucket
}

output "cloudfront_domain" {
  value = aws_cloudfront_distribution.frontend.domain_name
}

output "distribution_id" {
  value = aws_cloudfront_distribution.frontend.id
}//

output "certificate_arn"{
  value = var.certificate_arn.arn
}