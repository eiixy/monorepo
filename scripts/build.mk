## Example
example.build: example.build
example.image: example.image
example.publish: example.publish
example.deploy: example.deploy


.PHONY: generate
generate:
	wire ./cmd/...

build.all: generate
	go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./cmd/...

release.all:
	for service in $(shell ls ./cmd/); do \
	 $(MAKE) release.$$service; \
	done

test:
	go test -v ./internal/... -cover

%.gen:
	$(eval SERVICE:= $*)
	go generate ./cmd/$(SERVICE)/main.go

%.build:
	$(eval SERVICE:= $*)
	@echo "build: $(SERVICE):$(VERSION)"
	go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64
	go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/$(SERVICE) ./cmd/$(SERVICE)/

%.image:
	$(eval SERVICE:= $*)
	@$(MAKE) $(SERVICE).build
	sudo docker tag $(SERVICE):$(VERSION) $(REGISTRY)/$(SERVICE):$(VERSION)
	sudo docker push $(REGISTRY)/$(SERVICE):$(VERSION)

%.publish:
	$(eval SERVICE:= $*)
	@$(MAKE) $(SERVICE).image
	@echo "publish golang-$(SERVICE)"