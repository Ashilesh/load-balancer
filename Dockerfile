# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY ./src ./

RUN go build -o /LB

CMD ["/LB", "-config", "config.json", "-log", "./"]

# docker run -v D:\Work\Project\golang_project\Balancer\config.json:/app/config.json -v D:\Work\Project\golang_project\Balancer\LB.log:/app/LB.log -p 8080:8080 --add-host=host.docker.internal:host-gateway --name lb-test-container lb-test