variable "ami_builder_name" {
  description = "Base name used for the AMI-builder IAM role, policy, and instance profile."
  type        = string
  default     = "ec2-ami-builder"
}
