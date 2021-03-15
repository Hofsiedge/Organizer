#!bin/sh
psql -U postgres -d $POSTGRES_DB -f source/schema.pgsql > /dev/null


for f in source/**/functions.pgsql; do
	psql -U postgres -d $POSTGRES_DB -f $f > /dev/null;
done
