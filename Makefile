all: bin bin/unix2json bin/aws2json
	@echo "Build complete"

bin:
	mkdir -p bin

bin/aws2json:
	cd aws2json ; go build -o ../bin/aws2json

bin/unix2json:
	cd unix2json ; go build -o ../bin/unix2json
