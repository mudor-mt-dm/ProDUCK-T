#!/bin/sh
# wait-for-postgres.sh Костыль ожидающий запуска постгрески. Хз может и не нужен, но пока пусть валяется.
set -e
host="$1"
shift
cmd="$@"
until PGPASSWORD=$DB_PASSWORD psql -h "$host" -u "postgres" -c '\q';do
    >&2 echo "Postgres is unavailible - sleeping"
    sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd