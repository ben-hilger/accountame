DB_URL="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?search_path=public&sslmode=disable"

atlas schema apply \
  -u "$DB_URL" \
  --to file://database/schema.sql \
  --dev-url "docker://postgres/15/dev?search_path=public"

atlas schema inspect \
  --url "$DB_URL" \
  --format '{{ sql . }}' > database/generated_schema.sql
