FROM golang:1.17-alpine AS builder

WORKDIR /go/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o music-temperature-api ./main.go

FROM gcr.io/distroless/base

COPY --from=builder /go/app/music-temperature-api .

CMD ["/music-temperature-api"]