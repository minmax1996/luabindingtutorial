build:
	go build -o hasher .

install: build
	cp hasher /usr/local/bin/hasher
	
build-libhasher:
	go build -o lib/libhasher.so -buildmode=c-shared .

install-libhasher: build-libhasher
	sudo cp lib/libhasher.so /usr/local/lib/libhasher.so
