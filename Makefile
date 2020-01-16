all: build-all

build-front:
	npm --prefix front install
	npm run --prefix front build

deploy-front: build-front
	rm -R /srv/jackmarshall/cards-editor/
	mv front/dist /srv/jackmarshall/cards-editor

build-back:
	go mod download
	go build -o bin/api ./bin/api/

deploy-back: build-back
	mv bin/api/api /srv/jackmarshall/cards
	sudo systemctl restart jackmarshall-cards

build-all: build-front build-back

deploy-all: deploy-front deploy-back

download-assets:
	rm -R /srv/jackmarshall/assets/pdf_generator/
	cp -R assets /srv/jackmarshall/assets/pdf_generator/
	CGO_CFLAGS_ALLOW="-Xpreprocessor" go run ./bin/fetch-cards-pdf/main.go -dest-dir /srv/jackmarshall/assets/pdf_generator/images/front -workers 10
