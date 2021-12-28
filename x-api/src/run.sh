echo "[docs] run swager ..."
swag init -g main.go

echo "[docs] done! -> start server ..."
go run main.go

echo "[server] run on port 8080!"

