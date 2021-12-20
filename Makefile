l/init:
	mkdir -p ./pkg/logs
	cp -n .env.example ./pkg/.env
	cp -n ./challenge/cities.txt ./pkg/cities.txt

l/test:
	@cd pkg && go mod tidy && go get -d -u && go test -v ./... -cover

l/run:
	@cd pkg && go mod tidy && go get -d -u && go build -o paackgo && ./paackgo

l/docker:
	docker-compose --env-file ./pkg/.env down --remove-orphans --rmi all
	docker-compose --env-file ./pkg/.env up --detach --build --force-recreate --remove-orphans
