all:
	go build -v ./cmd/tectonic-offline/
	./tectonic-offline

.PHONY: all