# Simple CRUD Applicatio
Web-приложение обогащает данные созданного пользователя(возраст, пол, страна). В этом проекте реализована пагинация и выборка данных по различным фильтрам, также имеется миграция базы данных. 

## Зависимости
- Golang (версия 1.21.3)
- База данных - PostgreSQL

## Установка и запуск проекта
  1) Клонируйте репозиторий.
  ```
  git clone https://github.com/Fluffi1235/apiDataEnrichment.git
  ```
  2) Заполните информацию о базе данных (измените файл docker-compose.yaml и config.evn).
     ### Пример:
  ```
  # Изменение config.evn
  connectDb = host=localhost port=5430 user=user password=password dbname=db sslmode=disable

  # Изменение docker-compose.yaml
  POSTGRESQL_USERNAME=user
  POSTGRESQL_PASSWORD=password
  POSTGRESQL_DATABASE=db
  ```
  3) Запустите БД помощью команды
  ```
  docker-compose up
  ```
  4) Соберите проект
  ```
  go build ./cmd/api/main.go
  ```
  5) Запустите ваше приложение командой
  ```
  go run ./cmd/api/main.go 
  ```
  **Примечание:** Убедитесь, что установлен Docker и Docker Compose на вашем компьютере.
## Использование
  - Post : `http://localhost:8080/api/createUser` - Создание пользователя. В параметрах надо ввести name(обязательный параметр), surname(обязательный параметр), patronymic(необязательно)
    ##### Примеры: 
    ```
    http://localhost:8080/api/createUser?name=Dmitriy&surname=Ushakov&patronymic=Vasilevich
    http://localhost:8080/api/createUser?name=Anna&surname=Karamash
    ```
  - Post : `http://localhost:8080/api/search` - Получение данных по фильтрам(name,surname,patronymic,age,gender,country). В теле запроса необходимо передать json структуру, для этого можно использовать такие инструменты как insomnia или postman. Параметр page отвечает за пагинацию данных(выводит по 5 пользователей), отсчет страниц начинается с 1.
    ##### Примеры: 
    ```
    Запрос: 
    http://localhost:8080/api/search
  
    Примеры json структур:
      Поле page обязательный параметр, для пагинации данных, нумерация страниц начинается с 1
  
      №1 = выведет пользователей с параметрами name и surname 
      {
        "name":"Dmitriy",
        "surname":"Ushakov",
        "page":1
      }
      №2 - выведет всех пользователей, так как фильтры пользователей не указаны
      {
        "page":1
      }
     ```
  - Put : `http://localhost:8080/api/update` - Обновление данных пользователя по id. В запросе надо заполнить все поля пользователя(id, name, surname, patronymic, age, gender, country).
    ##### Примеры: 
    ``` 
    http://localhost:8080/api/update?id=2&name=Dmitriy&surname=Ushakov&patronymic=Vasilevich&age=22&gender=male&countru=AA
    http://localhost:8080/api/update?id=10&name=Anna&surname=Sergeevna&patronymic=Artemovna&age=35&gender=female&countru=RU
    ```
  - Delete : `http://localhost:8080/api/deleteUserById` - Удаление пользователя. В запросе введите id пользователя. 
    ##### Примеры: 
    ``` 
    http://localhost:8080/api/deleteUserById?id=10
    http://localhost:8080/api/deleteUserById?id=2
    ```
