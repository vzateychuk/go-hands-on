# Работа с модулями go

## Задача
Модули являются переносимыми библиотеками.  Пример: хотим создать vez/logging и вызвать его из main, вот так:

```go
package main

import "vez/logging"

func main() {
    logging.Debug(true)

    logging.Log("This is a debug statement...")
}
```

Создаем в папку **src/main**, в ней сохраняем main.go и инициализируем модуль
```shell
go mod init main
```
при этом в директории **src/main** из которой выполнялась команда _go mod init_, создается файл go.mod.

Рядом создаем папку **src/logging**, выполняем в ней аналогично
```shell
go mod init vez/logging
```

будет создан **src/logging/go.mod** содержанием:
```go
module vez/logging

go 1.16
```

здесь же создаем **src/logging/logging.go**
```go
package logging

import (
    "fmt"
    "time"
)

var debug bool

func Debug(b bool) {
    debug = b
}

func Log(statement string) {
    if !debug {
        return
    }

    fmt.Printf("%s %s\n", time.Now().Format(time.RFC3339), statement)
}
```

Если сейчас в директории **src/main** выполнить:
```shall
go run main.go
```
Тогда запуск не произойдет, поскольку main.go пытается импортировать зависимость vez/logger из %GOROOT%.
Обычные dependency мы получаем выполняя download командой
```shell
go mod download vez/logging                                                       
```
Но здесь мы получим ошибку _go mod download: module vez/logging: not a known dependency_ ,  поскольку go будет пытаться получить этот dependency с сайта https://vez/logging, и конечно обломается.
Мы можем "подсказать" go что искать эту dependency можно в соседней папке ../logging, для этого добавить в src/main/go.mod строку
```go
replace vez/logging => ../logging
```
теперь при попытке запуска go run main.go, golang найдет нужный модуль и мы получим сообщение что:
```shell
module vez/logging provides package vez/logging and is replaced but not required; to add it:                                                                                            go get vez/logging
```

Можно запустить **go get vez\logging** и тогда сообщение:
```shell
go get: added vez/logging v0.0.0-00010101000000-000000000000 
```
В результате в **src/main/go.mod** атоматически будет добавлен импорт модуля vez/logging.
Полностью **src/main/go.mod** станет таким:
```go
module main

go 1.16

replace vez/logging => ../logging

require vez/logging v0.0.0-00010101000000-000000000000 // indirect
```

И тогда можно запускать main.go на исполнение:
```shall
go run main.go

2021-09-29T23:42:42+03:00 This is a debug statement... 
```

## Links

[Введение в систему модулей Go](https://habr.com/ru/post/421411/)

[Похожий пример в TUTORIAL "Видимость пакетов в Go" ](https://www.digitalocean.com/community/tutorials/importing-packages-in-go)

[Туториал по импорту packages на DigitalOcean](https://www.digitalocean.com/community/tutorials/importing-packages-in-go)