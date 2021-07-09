
build:
	@echo "Building confd..."
	@mkdir -p bin
	@cd src/main && go build  -o ../../bin/confd .

clean:
	@rm -f bin/*