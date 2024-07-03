build:
	echo "Building lambda binaries"
	env GOOS=linux GOARCH=arm64 go build -o build/lambda/mssbp/bootstrap cmd/lambda/mssbp/main.go

zip:
	zip -j build/lambda/mssbp.zip build/lambda/mssbp/bootstrap
