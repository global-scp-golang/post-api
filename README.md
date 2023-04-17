## 서버 실행

Server와 DB 모두 Docker Container로 실행

```
docker-compose up -d
```

- 중지 & 삭제
```
docker-compose down --rmi all
```


## DB 접속
```
docker exec -it post-api_db_1 /bin/bash
 
psql -U postgres
```