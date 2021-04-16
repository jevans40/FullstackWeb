.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/test test/main.go test/auth.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/RDSPublicConnection RDSPublicConnection/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/AddProfile AddProfile/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose --stage dev --region us-east-2
