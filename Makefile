CC      	= go
PROGRAM		= validator
BUILDDIR	= build

.PHONY: build clean cleanBuild run

build: $(PROGRAM)

$(PROGRAM):
	$(CC) build

clean:
	rm -f $(PROGRAM)

cleanBuild:
	rm -rf $(BUILDDIR)


run: cleanBuild clean build

