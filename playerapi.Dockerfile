FROM golang:1.11.4 as builder
WORKDIR /go/src/github.com/ksimir/grpc-go-api
COPY cmd/player-server cmd/player-server
COPY pkg pkg
WORKDIR /go/src/github.com/ksimir/grpc-go-api/cmd/player-server

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
ENV PROJECTID=""
ENV INSTANCE=""
ENV DATABASE=""
COPY --from=builder /go/src/github.com/ksimir/grpc-go-api/cmd/player-server/player-server . 
EXPOSE 50051
ENTRYPOINT [ "sh", "-c", "./player-server -grpc-port=50051 -project=$PROJECTID -instance=$INSTANCE -database=$DATABASE"]