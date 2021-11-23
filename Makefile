.SILENT:

#=====MAIN_COMMAND=====================

.PHONY: up
up: up_docker info

.PHONY: up_see
up_see: 
	docker-compose up

.PHONY: info
info: ps info_domen

.PHONY: up_docker
up_docker:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down --remove-orphans

# флаг -v удаляет все volume (очищает все данные)
.PHONY: down-clear
down-clear:
	docker-compose down -v --remove-orphans

.PHONY: build
build:
	docker-compose build

.PHONY: ps
ps:
	docker-compose ps

#=====MAIN_COMMAND=====================
.PHONY: info_domen
info_domen:
	echo '---------------------------------';
	echo '----------DEV--------------------';
	echo GRAYLOG - http://localhost:9000;
	echo TG API - http://localhost:8080;
	echo '---------------------------------';


