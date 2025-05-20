.PHONY: gen-proto run-users run-posts run-comments run-web run-all build-all

gen-proto: gen-proto-comments gen-proto-users gen-proto-posts

gen-proto-comments:
	protoc --proto_path=. --micro_out=comments --go_out=comments comments/proto/comments.proto

gen-proto-users:
	protoc --proto_path=. --micro_out=users --go_out=users users/proto/users.proto

gen-proto-posts:
	protoc --proto_path=. --micro_out=posts --go_out=posts posts/proto/posts.proto

run-users:
	cd users && go run main.go

run-posts:
	cd posts && go run main.go

run-comments:
	cd comments && go run main.go

run-web:
	cd web && go run main.go



build-all: build-users build-posts build-comments build-web

build-users:
	cd users && go build -o ../bin/users

build-posts:
	cd posts && go build -o ../bin/posts

build-comments:
	cd comments && go build -o ../bin/comments

build-web:
	cd web && go build -o ../bin/web


clean:
	rm -rf bin/
	mkdir -p bin/