FROM golang:latest
WORKDIR /go/src/github.com/ksimir/grpc-go-api
COPY cmd/server cmd/server
WORKDIR /go/src/github.com/ksimir/grpc-go-api/cmd/server

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

#FROM scratch
#COPY --from=builder /go/src/github.com/ksimir/grpc-go-api/cmd/server/server . 
EXPOSE 50051
ENTRYPOINT ["./server -grpc-port=50051 -project=samirh-sandbox -instance=test-instance -database=game-a"]