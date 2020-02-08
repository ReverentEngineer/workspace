SOURCES = $(wildcard *.go)
INSTALL_PREFIX := $(HOME)/go

workspace: $(SOURCES)
	go build -o $@

install: workspace
	go install

clean:
	rm -f ./workspace
.PHONY: clean
