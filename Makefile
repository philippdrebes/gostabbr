test:
	@go test -v -cover ./...

cover:
	@go test -coverprofile=c.out -v -cover ./...
	@go tool cover -html=c.out

watch:
	@reflex -r '\.go$$' -s -- go test -v -cover ./...

run:
	@go run -v ./...

graph:
	@go test ./...
	@fdp -Tsvg -O engine/world.dot
	@mv engine/world.dot.svg ./world.svg
