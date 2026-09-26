resource "aws_secretsmanager_secret" "tasklist" {
    name = "task-list/secrets1"
}


resource "aws_secretsmanager_secret_version" "task_list" {
  secret_id = aws_secretsmanager_secret.tasklist.id

  secret_string = jsonencode({
    DATABASE_URL = var.db_url
  })
}