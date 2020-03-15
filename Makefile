PROTOC_VERSION 	= 3.7.0

MY_GOARCH 	= $(shell go env GOARCH | sed -e 's/amd64/x86_64/' -e 's/386/x86_32/')
MY_GOOS 		= $(shell go env GOOS | sed -e 's/darwin/osx/')

dependencies:
	mkdir -p /tmp/protoc && \
	curl -sLk https://github.com/google/protobuf/releases/download/v$(PROTOC_VERSION)/protoc-$(PROTOC_VERSION)-$(MY_GOOS)-$(MY_GOARCH).zip | \
		tar -xzv -C /tmp/protoc
	mv /tmp/protoc/bin/protoc $(GOPATH)/bin
	rm -rf /usr/local/include/google
	mv /tmp/protoc/include/google /usr/local/include/
	go get -u github.com/grpc-ecosystem/grpc-gateway/{protoc-gen-grpc-gateway,protoc-gen-swagger}
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	# $(RM) -rf /tmp/protoc