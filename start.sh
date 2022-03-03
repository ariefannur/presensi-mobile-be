#!/bin/sh

set -e
echo "run migrations"
/app/migrate -path /app/migration -database "$DB_SERVER_URL" --verbose up

echo "start the app"
exec "$@"