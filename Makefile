.PHONY: build all collector db grafana temperature-collector container-resource-collector

build:
	docker-compose --profile all build

all:
	docker-compose --profile all up -d

collector:
	docker-compose --profile collector up -d --force-recreate 

db:
	docker-compose --profile db up -d

grafana:
	docker-compose --profile grafana up -d

backup:
	docker-compose run backup sh -c 'cp -Rv /grafana_source/* /grafana_dest'

restore:
	docker-compose run backup sh -c 'cp -Rv /grafana_dest/* /grafana_source'

jaeger:
	docker-compose --profile jaeger up -d

cadvisor:
	docker-compose --profile cadvisor up -d --force-recreate

prometheus:
	docker-compose --profile prometheus up -d --force-recreate

temperature-collector:
	docker-compose up -d --force-recreate temperature-collector

container-resource-collector:
	docker-compose up -d --force-recreate --build container-resource-collector

