.PHONY: build tool cli clean

build: clean bin

tool cli:
	$(MAKE) -C $@

bin: tool cli
	mkdir bin
	rm tool/tool
	mv cli/cli bin/govel
	cd bin && ./govel new test

clean:
	rm -fR ./bin
