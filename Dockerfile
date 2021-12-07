# Start from golang base image
FROM golang:alpine as builder

LABEL maintainer="Nghia Nguyen <ngocnghia128@gmail.com>"


WORKDIR /api

COPY api/go.mod ./
COPY api/main.go ./

RUN go mod download 

COPY api ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


# Deploy golang app
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /api/main .
# COPY --from=builder /app/.env .   
COPY --from=builder /api/template template

EXPOSE 8080

CMD ["./main"]
