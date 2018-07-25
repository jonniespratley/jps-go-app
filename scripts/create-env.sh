#!/bin/bash
# create-env.sh
sess_secret=$(uuidgen)
uuid=$(uuidgen)
db_passwd=${uuid:0:12}
echo -e SESSION_SECRET=${sess_secret} > dep.env 
echo -e DATABASE_URL=\"postgres://postgres:${db_passwd}@pg_db:5432/jps_go_app_development?sslmode=disable\" >> dep.env;
echo -e POSTGRES_PASSWORD=${db_passwd} >> dep.env;