QUERY="SELECT cql_version from system.local;"

until echo $QUERY | cqlsh; do
  echo "cqlsh: cassandra is unavaliable - retry..."
  sleep 2
done &

exec /docker-entrypoint.sh "$@"
