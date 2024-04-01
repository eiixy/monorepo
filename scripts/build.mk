## Example
example.build: example.build
example.image: example.image
example.publish: example.publish

account.build: account.build
account.image: account.image
account.publish: account.publish


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
	docker build -t $(SERVICE):$(VERSION) -f ./deploy/build/$(SERVICE)/Dockerfile .

%.publish:
	$(eval SERVICE:= $*)
	@$(MAKE) $(SERVICE).image
	@echo "publish $(SERVICE)"
	docker tag $(SERVICE):$(VERSION) $(IMAGE_REGISTRY)/$(SERVICE):$(VERSION)
	docker push $(IMAGE_REGISTRY)/$(SERVICE):$(VERSION)