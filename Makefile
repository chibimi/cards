all: build-all

.PHONY: help
help: ## Get help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build-front: ## Build the client app
	npm --prefix front install
	npm run --prefix front build

deploy-front: build-front ## Deploy the client app
	rm -R /srv/jackmarshall/cards-editor/
	mv front/dist /srv/jackmarshall/cards-editor

build-back: ## Build the server app
	go mod download
	go build -o bin/api ./bin/api/

deploy-back: build-back ## Deploy the server app
	mv bin/api/api /srv/jackmarshall/cards
	cp -R assets/* /srv/jackmarshall/assets
	sudo systemctl restart jackmarshall-cards

build-all: build-front build-back ## Build both the client and the server

deploy-all: deploy-front deploy-back ## Deploy both the client and the server

fronts_dir ?= ./assets/fronts

download-assets: ## Download the cards front from the privateer press cards database
	mkdir -p $(fronts_dir)
	CGO_CFLAGS_ALLOW="-Xpreprocessor" go run ./bin/fetch-cards-pdf/main.go -dest-dir $(fronts_dir) -dsn "jackmarshall:$(password)@tcp(localhost:3306)/jackmarshall"

download-assets-light: ## Only downlaod attachment and secondary cards
	CGO_CFLAGS_ALLOW="-Xpreprocessor" go run ./bin/fetch-cards-pdf/main.go -secondary-only -dest-dir $(fronts_dir) -dsn "jackmarshall:$(password)@tcp(localhost:3306)/jackmarshall"
	mv $(fronts_dir)/* /srv/jackmarshall/assets/fronts/

find-ids:
	go run ./bin/find-ids/main.go -dsn "jackmarshall:$(password)@tcp(localhost:3306)/jackmarshall"
