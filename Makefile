.PHONY: migrate-dev
migrate-dev:
	atlas migrate hash
	atlas migrate apply --revisions-schema public --url "postgres://postgres:postgres@localhost:5432?search_path=public&sslmode=disable"
