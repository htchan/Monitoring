.PHONY: build all collector db visualize

build:
	docker-compose --profile all build

all:
	docker-compose --profile all up -d --force-recreate --build

collector:
	docker-compose --profile fluentd up -d --force-recreate --build 

db:
	docker-compose --profile db up -d --force-recreate --build

visualize:
	docker-compose --profile visualize up -d --force-recreate --build