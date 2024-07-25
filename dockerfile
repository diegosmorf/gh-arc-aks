FROM golang:1.22.4 as builder

ARG TARGETOS
ARG TARGETARCH

USER 0
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY ./cmd/producer  ./cmd/producer
RUN CGO_ENABLED=0 GOOS=linux GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o main -v -mod=readonly ./cmd/producer 

FROM alpine:3
COPY --from=builder /app/main /main
CMD ["/main"]
