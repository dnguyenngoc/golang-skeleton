#====================================================
# Start from golang base image
#====================================================
FROM golang:alpine as builder

LABEL maintainer="Nghia Nguyen <ngocnghia128@gmail.com>"

WORKDIR /src

COPY src/go.mod ./
COPY src/main.go ./

RUN go mod download 

COPY ./src ./

# Create swager docs
RUN go get github.com/swaggo/swag/cmd/swag
RUN swag init -g main.go

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


#====================================================
# Deploy golang app
#====================================================
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /src/main .

# Copy template (not need)
COPY --from=builder /src/template template
# Copy loggers folder
COPY --from=builder /src/loggers loggers
# Copy loggers folder
COPY --from=builder /src/storage storage
# Copy settings env
COPY --from=builder /src/settings/environment.toml environment.toml
# Copy swagger build
COPY --from=builder /src/docs docs

EXPOSE 8080

CMD ["./main"]
