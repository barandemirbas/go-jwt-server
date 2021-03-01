format:
	gofmt -l -s -w .

run:
	./bin/server

build:
	go mod tidy && go mod vendor
	go build -o ./bin/server && chmod +X ./bin/*
	echo "JWT Auth Boilerplate ready to use in ./bin/server"

cross-compile:
	env GOOS=linux GOARCH=arm go build -o ./release/server-linux-arm32 -ldflags "-s -w" -trimpath -mod=readonly
	env GOOS=linux GOARCH=arm64 go build -o ./release/server-linux-arm64 -ldflags "-s -w" -trimpath -mod=readonly
	env GOOS=darwin GOARCH=amd64 go build -o ./release/server-mac-x64 -ldflags "-s -w" -trimpath -mod=readonly
	env GOOS=linux GOARCH=386 go build -o ./release/server-linux-x32 -ldflags "-s -w" -trimpath -mod=readonly
	env GOOS=linux GOARCH=amd64 go build -o ./release/server-linux-x64 -ldflags "-s -w" -trimpath -mod=readonly
	env GOOS=windows GOARCH=386 go build -o ./release/server-windows-x32.exe -ldflags "-s -w" -trimpath -mod=readonly
	env GOOS=windows GOARCH=amd64 go build -o ./release/server-windows-x64.exe -ldflags "-s -w" -trimpath -mod=readonly

m1:
	env GOOS=darwin GOARCH=arm64 go build -o ./release/server-mac-arm64 -ldflags "-s -w" -trimpath -mod=readonly

clean:
	rm -rf ./bin
	rm -rf ./release
	rm -rf ./vendor
