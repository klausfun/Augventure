#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
shift
cmd="$@"

until PGPASSWORD=$DB_PASSWORD psql -h "$host" -U "postgres" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd
##!/bin/sh
## wait-for-postgres.sh
#
#set -e
#
#host="$1"
#shift
#cmd="$@"
#
#until PGPASSWORD=$DB_PASSWORD psql -h "$host" -U "postgres" -c '\q'; do
#  >&2 echo "Postgres is unavailable - sleeping"
#  sleep 1
#done
#
#>&2 echo "Postgres is up - running migrations"
#migrate -path ./schema -database "postgres://postgres:$DB_PASSWORD@$host:$DB_PORT/postgres?sslmode=disable" up
#
#>&2 echo "Migrations completed - executing command"
#exec $cmd
#
