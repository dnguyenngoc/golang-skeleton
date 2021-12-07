# Start from golang base image
FROM golang:alpine as builder

LABEL maintainer="Nghia Nguyen <ngocnghia128@gmail.com>"


WORKDIR /src

COPY src/go.mod ./
COPY src/main.go ./

RUN go mod download 

COPY ./src ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


# Deploy golang app
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /src/main .
# COPY --from=builder /app/.env .   
COPY --from=builder /src/template template

EXPOSE 8080

CMD ["./main"]
