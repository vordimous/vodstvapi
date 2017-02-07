go api for esvods

install:
follow steps to install godep
make sure vodstv/core is in your src folder
then run godep save ./... 

to import changes, commit then run godep update ./...

REDIS_URL: redis://h:p6e2kef2lalf188l57onlcqgtcg@ec2-54-243-188-149.compute-1.amazonaws.com:15199
REDIS_URL: redis://h:127.0.0.1:6379
DATABASE_URL: postgres://super:eoM9WJvSPl6lPfKO@127.0.0.1/vodstv?sslmode=disable

docker run -d --name="pg-vodstv" \
             -p 127.0.0.1:5432:5432 \
             -v /Users/ajdanelz/pg-data/vodstv:/data \
             -e USER="vodsapi" \
             -e DB="vodstv" \
             -e PASS="$(pwgen -s -1 16)" \
             paintedfox/postgresql

docker logs postgresql

docker run --name redis -p 127.0.0.1:6379:6379 -d redis

docker run -d --name="pg-vodstv" \
             -p 127.0.0.1:5432:5432 \
             -v /Users/ajdanelz/pg-data/vodstv:/data \
             -e USER="vodsapi" \
             -e DB="vodstv" \
             -e PASS="vodsapi" \
             paintedfox/postgresql

docker run --name="pg-vodstv" -itd --restart always \
  --publish 5432:5432 \
  --volume /Users/ajdanelz/pg-data/vodstv:/var/lib/postgresql \
  --env 'PG_TRUST_LOCALNET=true' \
  --env 'DB_USER=vodsapi' --env 'DB_PASS=vodsapi' \
  --env 'DB_NAME=vodstv' \
  sameersbn/postgresql:9.6-2