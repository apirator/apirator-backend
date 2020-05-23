# Builder Image
FROM golang:1.14-stretch as builder
ENV GO111MODULE=on
COPY . /apirator-backend
WORKDIR /apirator-backend
RUN go get github.com/google/wire/cmd/wire && \
    wire && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o bin/application

# Run Image
FROM scratch
COPY --from=builder /apirator-backend/bin/application /apirator-backend/bin
EXPOSE 9999
ENTRYPOINT ["./apirator-backend/bin"]