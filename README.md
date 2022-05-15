### 起動方法
#### バックエンドとデータベース
- ```$ cd backend```
- ```$ docker compose up -d```
- ```$ docker compose exec db bash```
- ```$ mysql -u root -p```
- ```$ password: root_password```
- ```mysql> CREATE DATABASE test_database;```
- ```mysql> exit```
- ```$ docker compose exec go sh```
- ```/go/src # go run cmd/main.go -option=migrate```
    - tableを使うためのオプションなので、初回のみ。二回目以降は```go run cmd/main.go```でOK。

#### フロントエンド
- ```cd ../frontend```
- ```npm install```
- ```npm run start```