.PHONY: build all collector ingester frontend log metrics trace backup restore

build:
	docker-compose --profile all build

all:
	docker-compose --profile all up -d

collector:
	docker-compose --profile collector up -d --force-recreate 

ingester:
	docker-compose --profile ingester up -d --force-recreate 

frontend:
	docker-compose --profile frontend up -d

log:
	docker-compose --profile logs up -d

metrics:
	docker-compose --profile logs up -d

trace:
	docker-compose --profile trace up -d

backup:
	docker-compose run backup sh -c 'cp -Rv /grafana_source/* /grafana_dest'

restore:
	docker-compose run backup sh -c 'cp -Rv /grafana_dest/* /grafana_source'
