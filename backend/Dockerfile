FROM golang:1.24 AS prod

WORKDIR /apps

COPY . .

RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o /taskbackend

FROM scratch

COPY --from=prod /taskbackend /taskbackend

EXPOSE 8080

ENTRYPOINT [ "/taskbackend"]
