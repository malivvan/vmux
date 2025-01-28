.PHONY: build

test-run: build run


run:
	@killall vmux || true
	@rm -rf /tmp/vmux
	@kitty sh -c "./build/vmux -mouse"

build:
	@mkdir -p build
	@rm -f ./build/*
	@CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o build/vmux .

clean:
	@rm -rf build