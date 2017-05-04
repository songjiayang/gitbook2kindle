all: gobuild gotest

godev:
	cd app && go run main.go

gobuild: goclean goinstall

gorebuild: goclean goreinstall

goclean:
	rm -rf ./bin ./pkg

goinstall:
	go get github.com/jordan-wright/email

goreinstall:
	go get -a -v github.com/jordan-wright/email

gotest:

gopackage:
	mkdir -p bin && go build -a -o bin/gitbook2kindle src/github.com/gitbook2kindle/main.go

travis: gobuild gotest

release:
	sh package.sh
	ghr -replace v0.1.0 pkg/
