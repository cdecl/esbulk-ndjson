
GOPATH=$(CURDIR)
GOBIN=$(GOPATH)/bin
GOFILES=esbulk-ndjson
EXEC=esbulk-ndjson.exe

build:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(GOBIN)/$(EXEC) $(GOFILES)

run:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOFILES) 

get:
	-GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -d "github.com/elastic/go-elasticsearch"
	
cc:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN)  GOOS=linux GOARCH=amd64 go build -o $(GOBIN)/linux/$(EXEC) $(GOFILES)
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN)  GOOS=windows GOARCH=amd64 go build -o $(GOBIN)/windows/$(EXEC) $(GOFILES)


