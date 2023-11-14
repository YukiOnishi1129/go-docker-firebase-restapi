include .env

empty:
	echo "empty"

# コンテナ環境へsshログイン
backend-ssh:
	docker exec -it ${BACKEND_CONTAINER_NAME} sh


# ローカル開発用
# go library install
go-add-library:
	docker exec -it ${BACKEND_CONTAINER_NAME} sh -c "go get ${name}"