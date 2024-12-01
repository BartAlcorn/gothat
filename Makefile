
# tg - templ generate - stand alone
.PHONY: templ
templ:
	templ generate --path ./pkg

# tmpwatch - templ generate --watch w proxy
.PHONY: templw
templw:
	templ generate --watch --path ./pkg --proxy="http://localhost:9090" --proxybind="localhost" --proxyport="9191" --open-browser=true -v

# generate tailwind css, minify and watch changes   --minify
.PHONY: css
css:
	npx tailwindcss -i ./pkg/embedded/assets/css/tailwind.css -o ./pkg/embedded/assets/css/styles.css

# generate tailwind css, minify and watch changes   --minify
.PHONY: cssw
cssw:
	npx tailwindcss -i ./pkg/embedded/assets/css/tailwind.css -o ./pkg/embedded/assets/css/styles.css --watch

.PHONY: air
air:
	air

# run both templ and air concurrently
.PHONY: dev
dev:
	make -j3  cssw air templw

# Go Targets
.PHONY: build
build:
	npx tailwindcss -i ./pkg/embedded/assets/css/tailwind.css -o ./pkg/embedded/assets/css/styles.css -minify
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ./main -tags lambda.norpc ./cmd/app/main.go

.PHONY: update/air
update/air:
	go install github.com/air-verse/air@latest

.PHONY: update/templ
update/templ:
	go get -u github.com/a-h/templ
	go install github.com/a-h/templ/cmd/templ@latest
