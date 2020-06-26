brew install mysql
brew services start mysql
mysql -uroot
CREATE USER IF NOT EXISTS 'cards_api'@'localhost' IDENTIFIED BY 'cards_api';
GRANT ALL ON cards_db.* TO 'cards_api'@'localhost';
mysql -u cards_api -p

npm install
npm audit fix
npm run serve

./set_env.sh
go get ./...
go run main.go

scp user@host:saves/jackmarshall-200626.sql .
mysql -u cards_api -p cards_db < jackmarshall-200626.sql


export CGO_CFLAGS_ALLOW='-Xpreprocessor'
go run bin/fetch-cards-pdf/main.go --dest-dir assets/images/front




cp assets/images/front/0.png /srv/jackmarshall/assets/pdf_generator/images/front/