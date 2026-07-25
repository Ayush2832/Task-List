# Frontend
- Change url to your backend url in config.js

# Backend
- change the domain in the backend .env. Add the Frontend url there.
- Port 8080.
- Dockerfile 

```yml
FROM golang:1.24 AS prod

WORKDIR /apps

COPY . .

RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o /taskbackend

FROM scratch

COPY --from=prod /taskbackend /taskbackend

EXPOSE 8080

ENTRYPOINT [ "/taskbackend"]
```

# Cloud
- Now we our code is ready. To run this my simple plan is
    - Push the Frontend in the S3 and then serve it with CDN.
    - Then put the backend behind the load balancers. 
    - We can scale this with the ASG.

# Strucuture
```
project/
│
├── frontend/
│   ├── index.html
│   ├── config.js
│   └── ...
│
├── backend/
│   ├── main.go
│   ├── Dockerfile
│   └── ...
│
├── infrastructure/
│   │
│   ├── frontend/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   └── outputs.tf
│   │
│   ├── backend/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   └── outputs.tf
│   │
│   ├── database/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   └── outputs.tf
│   │
│   
│
└── .github/
    └── workflows/
        ├── frontend-deploy.yml
        ├── backend.yml
        └── infra-frontend.yml
```