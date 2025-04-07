run:
	docker-compose up -d

stop:
	docker-compose down

build:
	docker-compose up -d --build

hot:
	curl -v -X POST http://localhost:8888/schedule -d '{"schedule_type": "hot"}'

cold:
	curl -v -X POST http://localhost:8888/schedule -d '{"schedule_type": "cold"}'