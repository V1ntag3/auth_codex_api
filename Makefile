build: go build -o server main.go

run: build ./server

watch: /home/marcos/go/bin/reflex  '\.go$$'  make