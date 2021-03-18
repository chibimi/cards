brew install mysql  
brew services start mysql  
mysql -uroot  
CREATE USER IF NOT EXISTS 'jackmarshall'@'localhost' IDENTIFIED BY 'jackmarshall';  
GRANT ALL ON jackmarshall.* TO 'jackmarshall'@'localhost'; 
mysql -u jackmarshall -p  

npm install  
npm audit fix  
npm run serve  

./set_env.sh  
go get ./...  
go run main.go  

scp user@host:saves/jackmarshall-210317.sql .  
mysql -u jackmarshall -p jackmarshall < jackmarshall-210317.sql  
GRANT ALL ON jackmarshall.* TO 'jackmarshall'@'localhost'; 

export CGO_CFLAGS_ALLOW='-Xpreprocessor'  
go run bin/fetch-cards-pdf/main.go --dest-dir assets/images/front  




cp assets/images/front/0.png /srv/jackmarshall/assets/pdf_generator/images/front/  
