.PHONY: build

build:
	@mkdir -p build
	@rm -f ./build/*
	@CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o build/vmux .

clean:
	@rm -rf build