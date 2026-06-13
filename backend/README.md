# For db 

- First create the container 

```Dockerfile
docker run -d \
  --name postgres \
  -e POSTGRES_DB=taskdb \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=NyQ8jC5eLCZhGGWd \
  -p 5432:5432 \
  -v postgres-data:/var/lib/postgresql/data \
  postgres:17
```

- Then ssh inside the container 
`docker exec -it postgres psql -U postgres -d taskdb`

- Then crate table

```
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    title TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```