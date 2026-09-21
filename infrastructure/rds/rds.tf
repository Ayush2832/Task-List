resource "aws_db_subnet_group" "subnet_group" {
  name       = "rds_subnet_group"
  subnet_ids = data.aws_subnets.subnets.ids
}

resource "aws_db_instance" "postgres" {
  identifier     = "task-list-db"
  engine         = "postgres"
  engine_version = "16"
  instance_class = "db.t3.micro"

  allocated_storage = 20
  storage_type      = "gp3"

  db_name  = "taskdb"
  username = var.db_username
  password = var.db_password

  db_subnet_group_name   = aws_db_subnet_group.subnet_group.name
  vpc_security_group_ids = [aws_security_group.rds_sg.id]

  skip_final_snapshot = true
  publicly_accessible = false
}