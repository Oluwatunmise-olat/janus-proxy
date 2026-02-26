.PHONY: work-sync fmt build build-dataplane build-controlplane build-diff-worker build-ops-console up down

work-sync:
	go work sync

build: build-dataplane build-controlplane build-diff-worker build-ops-console

build-dataplane:
	cd services/dataplane && go build ./...

build-controlplane:
	cd services/controlplane && go build ./...

build-diff-worker:
	cd services/diff-worker && go build ./...

build-ops-console:
	cd services/ops-console && go build ./...

up:
	docker compose -f deploy/compose/docker-compose.yml up -d

down:
	docker compose -f deploy/compose/docker-compose.yml down
