.PHONY: start stop restart shell

start:
	docker-compose up --detach go

stop:
	docker-compose down --remove-orphans --volumes --timeout 0

restart: stop start

shell: start
	docker-compose exec go sh
