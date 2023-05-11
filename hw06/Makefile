tests:
	go test -v ./...

build:
	go build -o "./generator" "./cmd/generator/generator.go"
	go build -o "./summator" "./cmd/summator/summator.go"

run_generator:
	./generator -input F0.txt -output F2.txt

run_summator:
	./summator -input F0.txt -output F1.txt

clear:
	rm F0.txt F1.txt F2.txt