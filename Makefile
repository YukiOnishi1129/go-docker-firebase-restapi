include .env

empty:
	echo "empty"

# コンテナ環境へsshログイン
backend-ssh:
	docker exec -it ${BACKEND_CONTAINER_NAME} sh


# ローカル開発用
# go library install
## 複数のライブラリを指定する場合は、name="xxx yyy" のように""で囲んで実行すること
go-add-library:
	docker exec -it ${BACKEND_CONTAINER_NAME} sh -c "go get ${name}"

# swagger
gen-swagger:
	swag init -g app/cmd/main.go  --output docs/swagger/config

swagger-build:
	docker-compose -f docker-compose.swagger.yml build

swagger-up:
	docker-compose -f docker-compose.swagger.yml up -d

swagger-down:
	docker-compose -f docker-compose.swagger.yml down