#go build -o helloworld
# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o helloworld .
CGO_ENABLED=0 go build -o helloworld