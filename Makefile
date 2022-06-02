.PHONY: build all collector db visualize

build:
	docker-compose --profile all build

all:
	docker-compose --profile all up -d

collector:
	docker-compose --profile collector up -d --force-recreate 

db:
	docker-compose --profile db up -d

visualize:
	docker-compose --profile visualize up -d
