# make light
.PHONY: default image

default: image

image: bin
	docker build -t jackxu/mapper-cpu:v1.0.0 .
	docker push jackxu/mapper-cpu:v1.0.0

.PHONY: bin
bin:
	go build -o mapper-cpu

.PHONY: run
run: bin
	./mapper-cpu
