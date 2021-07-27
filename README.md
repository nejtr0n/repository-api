# Repository api

Сервис клонирования репозиториев в gitlab

* Запустить ```docker-compose up```
* Зайти на https://localhost/users/sign_in
* Залониниться как root с паролем adminadmin

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

Если же гитлаб не поднялся, нужно удалить папку data, переподнять gitlab,
настроить токен доступа для пользователя, создать репозиторий болванку и группу, в которую
будет создаваться клон, и поменять эти параметры в запросе.
