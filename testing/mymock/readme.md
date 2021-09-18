# For the purpose of demonstrating the creation and use of mocks

Для того чтобы воспользоваться mockgen, необходимо:
1. Код был модулем (присутствовал файл go.mod);
2. Установить **mockgen**;
3. Объявить зависимость текущего модуля от **mockgen/model**;
```shell
go get github.com/golang/mock/mockgen/model
```
Тогда автоматически будет добавлен файл go.sum и в go.mod будет добавлено:
> require github.com/golang/mock v1.6.0 // indirect
4. Запустить mockgen из папки
```shell
mockgen -destination mock/api-mock.go . API
```
где:
- . - указание в какой директории будет поиск интерфейсов;
- API - имя interface для которого будет создан mock;
Будет автоматически создана подпапка **mock** и в ней будет создан **api-mock.go**   
 
