.PHONY: build deploy clean

build:
	GOOS=linux GOARCH=amd64 go build -o ./app/bootstrap ./app/main.go

deploy:
	cd deploy-scripts && cdk deploy

deploy-swap:
	cd deploy-scripts && cdk deploy --hotswap

clean:
	rm -rf ./app/bootstrap

all:
	make build
	make deploy
all-swap:
	make build
	make deploy-swap
