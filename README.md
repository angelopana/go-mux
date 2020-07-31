# Testing postgres docker container using Golang


I wanted to get more familiar with docker and postgres. 
These are my notes for future me or anyone on how to access the postgres container from localhost.



## Tip to map your postgres docker container to localhost

[Link to Original post](https://stackoverflow.com/questions/37694987/connecting-to-postgresql-in-a-docker-container-from-outside)

Create your postgres container, you can run Postgres this way (map a port):

``` 
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres 
```

Peek at the docker container that was created:

```
docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                    NAMES
c02585a214a7        postgres            "docker-entrypoint.s…"   2 seconds ago       Up 1 second         0.0.0.0:5432->5432/tcp   some-postgres
```

Go inside the container & create a database:

```
docker exec -it c02585a214a7 bash
root@0c02585a214a7:/# psql -U postgres
postgres-# CREATE DATABASE mytest;
postgres-# \q

```

Go to the local terminal and map public-ip to to 5432

```
psql -h public-ip-server -p 5432 -U postgres
```
