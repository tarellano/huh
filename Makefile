build:
	go build -o bin/huh ./cmd/main.go

clean:
	rm -rf bin/

test: build
	./bin/huh --api-key=$$LASTFM_API_KEY --username=tarellano8 --artist "Tame Impala"