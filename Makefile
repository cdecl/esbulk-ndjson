
GOPATH=$(CURDIR)
GOBIN=$(GOPATH)/bin
GOFILES=esbulk
EXEC=esbulk.exe

build:
	@SET GOPATH=$(GOPATH)& SET GOBIN=$(GOBIN)&  go build -o $(GOBIN)/$(EXEC) $(GOFILES)

get:
	-SET GOPATH=$(GOPATH)& SET GOBIN=$(GOBIN)&  go get -d "github.com/elastic/go-elasticsearch"
	
cc:
	@SET GOPATH=$(GOPATH)& SET GOBIN=$(GOBIN)& SET GOOS=linux& SET GOARCH=amd64&  go build -o $(GOBIN)/linux/$(EXEC) $(GOFILES)
	@SET GOPATH=$(GOPATH)& SET GOBIN=$(GOBIN)& SET GOOS=windows& SET GOARCH=amd64&  go build -o $(GOBIN)/windows/$(EXEC) $(GOFILES)

clean:
	del bin\*.exe
