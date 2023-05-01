define generate_mocks
    mockery --packageprefix mocked --keeptree --name=$(2) --recursive --case underscore --dir ./internal/$(1) --output ./internal/test/mocks/$(1)
endef

.phony: generate-mocks
generate-mocks:
	find ./test/mocks | xargs rm -rf
	$(call generate_mocks,"manager","Manager")
	$(call generate_mocks,"repository","Repository")

stop-test:
	docker-compose -f docker-compose-test.yml down --rmi all -v --remove-orphans
	docker volume prune -f

start-test: #stop-test
	docker-compose -f docker-compose-test.yml up --build --abort-on-container-exit --force-recreate --exit-code-from entain

test:start-test

.PHONY: compose-up
compose-up:
	docker-compose down --rmi all -v --remove-orphans
	docker-compose up  --force-recreate
