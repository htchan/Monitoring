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

backup:
	docker-compose run backup sh -c 'cp -Rv /grafana_source/* /grafana_dest'

restore:
	docker-compose run backup sh -c 'cp -Rv /grafana_dest/* /grafana_source'

jaeger:
	docker-compose --profile jaeger up -d

cadvisor:
	docker-compose --profile cadvisor up -d