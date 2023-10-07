
clean:
	@rm -rf build

flash:
	tinygo flash -target=pico .	

smoke-test:
	@mkdir -p build
	tinygo build -size short -o ./build/test.hex -target=pico .

test: clean smoke-test
	go test
