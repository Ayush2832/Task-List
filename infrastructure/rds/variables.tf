variable "db_username" {
  description = "mention db name"
}
variable "db_password" {
  description = "mention db password. strong one"
  sensitive = true
}