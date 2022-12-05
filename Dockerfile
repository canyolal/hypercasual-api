FROM golang:alpine
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app ./cmd/api

FROM alpine
COPY --from=0 /app /app
ENV HYPERCASUAL_DSN=postgres://hypercasual:h4rdP4ssw0rd@localhost/hypercasual
EXPOSE 4001
ENTRYPOINT [ "/app"]