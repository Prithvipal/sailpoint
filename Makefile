
.PHONY: build
build: 
	GOOS=linux GOARCH=amd64 go build -o server

.PHONY: image
image: build
	docker build . -t sailpoint

.PHONY: run
run: image
	docker run sailpoint  -cfg.mail.pass=''

