FROM golang:alpine
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app ./cmd/api

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=0 /app /app
EXPOSE 4001 4001
ENTRYPOINT [ "/app"]