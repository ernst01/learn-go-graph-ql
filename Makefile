MAKEFLAGS += --silent

local:
	@echo "  >  Runs container..."
	bash -c "docker-compose up"