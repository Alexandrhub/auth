#!/bin/bash
source .env


sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DOCKER_DSN}" up -v