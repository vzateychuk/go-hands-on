# Breaking circular dependencies

В этом примере мне показалось интересным не только сам прием "разрыва зависимости", предложенный автором, а так же механизм "внедрения зависимостей" с помощью функциональной переменной AcquirePacker (returns a Packer instance).
Этот прием можно использовать как минимум в юнит-тестах, определяя значение функциональной переменной AcquirePacker как функцию возвращающую mock вместо warehouse.AcquireRobot(ctx). 

## Link
[Книга Hands-On Software Engineering with Golang 2020 Anagnostopoulos Achilleas](https://vk.com/wall-51126445_71212)