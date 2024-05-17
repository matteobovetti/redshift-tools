.PHONY: build run-retention run-copy run-unload

build:
	@docker build -t redshift-tool .

run-retention:
	@docker run --env-file=./.env.dev redshift-tool retention

run-copy:
	@docker run --env-file=./.env.dev redshift-tool copy

run-unload:
	@docker run --env-file=./.env.dev redshift-tool unload
