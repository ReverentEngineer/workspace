SOURCES = $(wildcard *.go)
ifeq ($(PREFIX),)
    PREFIX := /usr/local
endif

workspace: $(SOURCES)
	go build -o $@

install: workspace
	cp $^ $(PREFIX)/bin/$^
	chmod 755 $(PREFIX)/bin/$^

clean:
	rm -f ./workspace
.PHONY: clean
