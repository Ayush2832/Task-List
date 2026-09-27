# Task List

A simple multi-user to-do application deployed on AWS. A static frontend served through CloudFront, a Go REST API running on ECS Fargate, and a PostgreSQL database on RDS — all provisioned with Terraform and deployed via GitHub Actions.

- Frontend: `https://task.pingayush.in`
- API: `https://apitasklist.pingayush.in`

## Architecture

```
Browser
  ├── task.pingayush.in ──────► CloudFront ──(OAC)──► S3 (private, static site)
  └── apitasklist.pingayush.in ► ALB (443) ─► ECS Fargate (Go API :8080) ─► RDS PostgreSQL
                                                     │
                                                     ├── Secrets Manager (DATABASE_URL)
                                                     └── NAT Gateway (outbound)
```

- **Frontend** — Single `index.html` (vanilla JS, no build step) plus a runtime `config.js` holding the API URL. Stored in a private S3 bucket, served only through CloudFront using Origin Access Control (OAC). TLS via an ACM `*.pingayush.in` certificate.
- **Backend** — Go service using the standard library `net/http`. Runs as a container on ECS Fargate in private subnets, reachable only through an internet-facing Application Load Balancer.
- **Database** — RDS PostgreSQL 16 (`db.t3.micro`), private, not publicly accessible.
- **Networking** — VPC with public subnets (ALB, NAT) and private subnets (ECS tasks, RDS) across two AZs.
- **Secrets** — `DATABASE_URL` is read from AWS Secrets Manager at task start; never committed.

## Project Structure

```
.
├── backend/            # Go API (main.go, db.go, Dockerfile)
├── frontend/           # Static site (index.html, config.js)
├── local/              # Local dev via Nginx + Docker Compose
├── infrastructure/
│   ├── frontend/       # S3 + CloudFront + OAC + bucket policy (Terraform)
│   ├── backend/        # VPC data, ALB, ECS, IAM, Secrets Manager (Terraform)
│   └── rds/            # RDS PostgreSQL (Terraform)
└── .github/workflows/  # CI/CD pipelines
```

## API

All endpoints scope data to the caller via an `X-User-ID` header (a UUID generated per browser and stored in `localStorage`).

| Method | Path              | Description        |
|--------|-------------------|--------------------|
| GET    | `/health`         | Health check       |
| GET    | `/api/tasks`      | List the user's tasks |
| POST   | `/api/tasks`      | Create a task      |
| DELETE | `/api/tasks/{id}` | Delete a task      |

Environment variables: `PORT`, `ALLOWED_ORIGIN`, `DATABASE_URL`.

## Local Development

Runs the backend and Postgres in Docker Compose. Nginx serves the frontend and proxies `/api` to the backend.

```bash
cd local
docker compose up --build
```

Notes:
- Local `DATABASE_URL` uses `sslmode=disable` (the local Postgres container has no TLS). RDS uses `sslmode=require`.
- `site.local` and `api.site.local` are mapped to `127.0.0.1` via `/etc/hosts` to mirror the production frontend/API split.

## Deployment

Infrastructure is managed with Terraform and deployed through GitHub Actions using OIDC (region `us-east-2`).

Terraform stacks (apply in order):

```bash
cd infrastructure/frontend && terraform init && terraform apply
cd infrastructure/rds      && terraform init && terraform apply
cd infrastructure/backend  && terraform init && terraform apply
```

GitHub Actions workflows:
- `infra-frontend.yml` — orchestrates frontend infra, frontend deploy, RDS infra, and image build/push.
- `frontend-deploy.yml` — syncs `frontend/` to S3.
- `infra-rds.yml` — provisions RDS.
- `image-Dhub.yml` — builds and pushes the backend image to Docker Hub.
- `ecs.yml` — applies the backend infrastructure (ECS service, ALB, etc.).

Required CI config:
- Secret: `AWS_ROLE_ARN`, `DB_URL`, `BUCKET_NAME`
- Variable: `FRONTEND_URL` (must include scheme, e.g. `https://task.pingayush.in`)

## Tech Stack

Go · PostgreSQL · Docker · AWS (S3, CloudFront, ECS Fargate, ALB, RDS, Secrets Manager, ACM) · Terraform · GitHub Actions
