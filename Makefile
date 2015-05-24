export GOPATH:=$(GOPATH):/$(shell pwd)

amd64:
	GOARCH=amd64 go build -o bin/mycommander_amd64

386:
	GOARCH=386 go build -o bin/mycommander_386

arm:
	GOARCH=arm go build -o bin/mycommander_arm

clean:
	go clean -x && rm -rf pkg/ && rm -rf bin/*
