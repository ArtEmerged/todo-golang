build:
	docker build -t todo-app .
run:
	docker run -p 8000:8000 todo-app
migrate:
	sqlite3 todo.sqlite3 < migrations/init.sql