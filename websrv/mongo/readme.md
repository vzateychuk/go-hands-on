# Использование Mongo

#№ Запустить Mongo локально на порту 27017
```shell
docker run -tid --name mongo -p 27017:27017 mongo:latest
```

## Подключиться из Mongo Compass
```shell
mongodb://localhost:27017
```

## Links
[Coursera Разработка веб сервисов на Golang, part 2](https://www.coursera.org/learn/golang-webservices-2/lecture/69Ylo/document-store-mongodb)
[Описание параметров подключения Mongo Compass](https://docs.mongodb.com/manual/reference/connection-string/)
