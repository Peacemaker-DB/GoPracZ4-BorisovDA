<img width="519" height="911" alt="image" src="https://github.com/user-attachments/assets/490920a8-34f1-40a4-984a-e6dc1b3b327b" /># Практическое задание № 4 Борисов Денис Александрович ЭФМО-01-25

Цели:

1.	Освоить базовую маршрутизацию HTTP-запросов в Go на примере роутера chi.
2.	Научиться строить REST-маршруты и обрабатывать методы GET/POST/PUT/DELETE.
3.	Реализовать небольшой CRUD-сервис «ToDo» (без БД, хранение в памяти).
4.	Добавить простое middleware (логирование, CORS).
5.	Научиться тестировать API запросами через curl/Postman/HTTPie.

Выполнение работы:

1. Создаём скелет проекта

В ходе работы над практической работой были созданы следующий скелет проекта:

<img width="291" height="270" alt="image" src="https://github.com/user-attachments/assets/69d5c9c2-7903-483d-bea0-e6c2f7da1aad" />


После были инициализация модулей и установка зависимостей

<img width="732" height="199" alt="Снимок экрана 2025-10-07 152241" src="https://github.com/user-attachments/assets/58cbff8c-6cff-4543-95ef-d2d798c36492" />

2. Модель и «хранилище» в памяти

После был запущен сервер и проверен в браузере:

<img width="1918" height="719" alt="image" src="https://github.com/user-attachments/assets/29276571-1443-421f-b5cf-9bfab3d31b4d" />

После был написан код для model.go и для repo.go

Код model.go

<img width="422" height="229" alt="image" src="https://github.com/user-attachments/assets/2f52a741-ef86-4040-bbae-803870408242" />

Код repo.go

<img width="506" height="955" alt="image" src="https://github.com/user-attachments/assets/8684edb4-69f8-4925-b4b8-48638c5f01ec" />


3. Handlers: JSON API

Для выполнения шага практики был написан код для handler.go:

<img width="519" height="911" alt="image" src="https://github.com/user-attachments/assets/57555616-28ab-4d98-9e1b-7bda873cf173" />

4. Middleware: логирование и CORS

После был написан код для logger.go и cors.go

Код logger.go

<img width="916" height="502" alt="image" src="https://github.com/user-attachments/assets/9a1b8190-b128-4a4b-9313-37710d308f0c" />

Код cors.go

<img width="1046" height="507" alt="image" src="https://github.com/user-attachments/assets/1f48922b-8549-403c-935e-8152799062d8" />

5. main.go — сборка приложения

Затем быо написан код для main.go:

<img width="702" height="846" alt="image" src="https://github.com/user-attachments/assets/d9562347-5dff-470e-b8af-052bd4478185" />

6. main.go — сборка приложения

После выполнения всех шагов было выполнено тестирование:

Создание

<img width="716" height="770" alt="image" src="https://github.com/user-attachments/assets/3acfb106-76df-4e14-89b0-b9535f398efc" />

Список

<img width="725" height="930" alt="image" src="https://github.com/user-attachments/assets/a73fd83e-05fe-4547-a395-8cd4b1e1a6ba" />

Получить по id

<img width="717" height="780" alt="image" src="https://github.com/user-attachments/assets/1c8d4655-8c97-4607-b7f2-5ef17b5a9213" />

Обновить

<img width="723" height="785" alt="image" src="https://github.com/user-attachments/assets/fc477d3c-a713-4031-b948-61dedcf86097" />

Удалить

<img width="719" height="695" alt="image" src="https://github.com/user-attachments/assets/4a0fea0f-6b3c-4dee-abf5-76aa640c74ef" />

Список после удаления

<img width="721" height="813" alt="image" src="https://github.com/user-attachments/assets/32ca29df-7a21-4550-8e1b-b1dc263ade94" />

Дополнительные задания

1. Валидация длины title (минимум 3, максимум 100 символов).

Для валидации длины был дополнен код функции create и update в файле handler.go

Обновленный код функции create и update

<img width="589" height="672" alt="image" src="https://github.com/user-attachments/assets/3fc7368a-959d-4862-87c3-108096d53471" />


Проверка

Успех

<img width="715" height="769" alt="image" src="https://github.com/user-attachments/assets/aaa1d1a4-69ac-4208-889e-d852c42e05fb" />

Ошибка

<img width="707" height="698" alt="image" src="https://github.com/user-attachments/assets/b343762f-af01-4dd6-95f4-c0cc9914e642" />




