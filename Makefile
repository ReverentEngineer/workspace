SOURCES = $(wildcard *.go)
INSTALL_PREFIX := /usr/local

workspace: $(SOURCES)
	go build -o $@

install: workspace
	cp $^ $(INSTALL_PREFIX)/bin/$^
	chmod 755 $(INSTALL_PREFIX)/bin/$^

clean:
	rm -f ./workspace
.PHONY: clean
