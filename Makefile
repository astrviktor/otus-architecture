tests:
	cd ./hw01 && go test -v ./... && cd ..
	cd ./hw02 && go test -v ./... && cd ..
	cd ./hw03 && go test -v ./... && cd ..
	cd ./hw04 && go test -v ./... && cd ..
	cd ./hw05 && go test -v ./... && cd ..
	cd ./hw06 && go test -v ./... && cd ..
	cd ./hw07 && go test -v ./... && cd ..
	cd ./hw08 && go test -v ./... && cd ..
	cd ./hw09/apps/auth && go test -v ./... && cd ../../..
	cd ./hw09/apps/server && go test -v ./... && cd ../../..
	cd ./hw12 && go test -v ./... && cd ..
