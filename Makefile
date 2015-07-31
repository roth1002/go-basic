all:
	sh ./goinstall
	go run build.go ./
build:
	go run build.go ./
test:
	go test
vender:
	sh ./goinstall
