CRUD endpoints
--- 

- Проекте создается таблица БД из нескольких полей,
- Таблица заполняется фейковыми данными
- Создается сервер, endpoints: '/people/?id=...', 
    http://localhost:8080/people/ - список всех people
    http://localhost:8080/people/?id=xxx - информация по одному из people
    http://localhost:8080/ - детальная информация по пользователю
    http://localhost:8080/form - добавление пользователя
Запуск
---
1. Запустить БД Postgresql с помощью Docker и вместе с ним PGadmin:
```
docker run -tid --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=root postgres
docker run -tid --name pgadmin -p 5555:80 -e PGADMIN_DEFAULT_EMAIL=vlad@vlad.com -e PGADMIN_DEFAULT_PASSWORD=root dpage/pgadmin4
```
2. Уточнить ip адрес БД:
```
docker inspect postgres
```

