ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: serve
serve:
	go run main.go serve