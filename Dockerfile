FROM golang:1.11.4 as builder
WORKDIR /go/src/github.com/ksimir/grpc-go-api
COPY cmd/server cmd/server
COPY pkg pkg
WORKDIR /go/src/github.com/ksimir/grpc-go-api/cmd/server

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM scratch
ARG projectid
ARG instance
ARG database
COPY --from=builder /go/src/github.com/ksimir/grpc-go-api/cmd/server/server . 
EXPOSE 50051
ENTRYPOINT ["./server", "-grpc-port=50051", "-project=${projectid}", "-instance=${instance}", "-database=${database}"]