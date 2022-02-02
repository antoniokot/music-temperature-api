FROM golang:1.17-alpine AS builder

WORKDIR /go/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o music-temperature-api ./cmd/web/main.go

FROM gcr.io/distroless/base

COPY --from=builder /go/app/music-temperature-api .

HEALTHCHECK CMD curl --fail http://localhost:3333 || exit 1 

CMD ["/music-temperature-api"]