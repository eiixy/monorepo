# go.proxy
go.proxy:
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct

# install dependents
install.dependents:
	go install github.com/google/wire/cmd/wire@latest

# init env
init:
	git submodule init
	git submodule update
	$(MAKE) go.proxy
	$(MAKE) install.dependents
	go mod tidy
	cp .env.example .env
	cp deploy/components/.env.example deploy/components/.env
