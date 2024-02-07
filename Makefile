CC      	= go
PROGRAM		= validator

.PHONY: build clean run

build: $(PROGRAM)

$(PROGRAM):
	$(CC) build

clean:
	rm -f $(PROGRAM)

run: clean build

