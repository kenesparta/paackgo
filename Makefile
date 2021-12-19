l/up:
	docker-compose --env-file ./pkg/.env down --remove-orphans --rmi all
	docker-compose --env-file ./pkg/.env up --detach --build --force-recreate --remove-orphans

l/init:
	mkdir -p logs
	cp -n .env.example ./pkg/.env

l/test:
	cd pkg && \
	go mod tidy && \
	go get -d -u && \
	go test -v ./... -cover

l/run:
	cd pkg && \
	go build . && \
	./paackgo
