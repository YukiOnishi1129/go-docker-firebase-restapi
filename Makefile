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
	swag init -g app/cmd/main.go  --output docs/swagger

swagger-up:
	docker-compose -f docs/swagger/docker-compose.yml up -d

swagger-down:
	docker-compose -f docs/swagger/docker-compose.yml down