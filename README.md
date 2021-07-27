# Repository api

Сервис клонирования репозиториев в gitlab

* Запустить ```docker-compose up```
* Зайти на https://localhost/users/sign_in
* Залониниться как root с паролем adminadmin
* Настроить токен доступа для пользователя
* Создать репозиторий болванку 
* Создать группу, в которую будет создаваться клон

Если все нормально, то выполнить запрос

```
curl -i -X POST --location "http://localhost:8000/create-project-from-dull" \
    -H "Content-Type: application/json" \
    -d "{
          \"dull\": {
            \"provider_type\": \"gitlab\",
            \"project_id\": 34
          },
          \"new\": {
            \"provider_type\": \"gitlab\",
            \"name\": \"clone\",
            \"group_id\": 36
          }
        }"
```

Если всё прошло успешно, то в ответе будет
```
HTTP/1.1 204 No Content
Content-Type: application/json; charset=utf-8
Date: Mon, 26 Jul 2021 21:22:00 GMT
```
и в гитлабе появится новый репозиторий  


