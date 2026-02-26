.PHONY: work-sync fmt build build-dataplane build-controlplane build-diff-worker build-ops-console up down cp-migrate-up cp-migrate-down cp-migrate-create

DB_URL ?= postgres://janus:janus@localhost:5432/janus?sslmode=disable

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
	docker local -f infra/local/docker-local.yml up -d

down:
	docker local -f infra/local/docker-local.yml down

fmt:
	gofmt -w .

cp-migrate-up:
	migrate -path services/controlplane/migrations -database "$(DB_URL)" up

cp-migrate-down:
	migrate -path services/controlplane/migrations -database "$(DB_URL)" down 1

cp-migrate-create:
	migrate create -ext sql -dir services/controlplane/migrations -seq $(name)
