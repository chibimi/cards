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