# Proxy Server

## Описание
Это HTTP-сервер для проксирования запросов к сторонним сервисам. Сервер принимает запросы от клиента, отправляет их к указанным сервисам и возвращает ответы клиенту в формате JSON. Запросы и ответы сохраняются локально.

## Требования
- Docker
- Docker Compose
- Make

## Установка и запуск

### Клонирование репозитория
```sh
git clone https://github.com/yrss1/hl_easy
cd hl_easy
```


### Сборка и запуск с помощью Docker Compose
```sh
make deploy
```
### Остановка сервиса
```sh
make down
```
### Логи
```sh
make logs
```

### Использование
Сервер слушает на порту 8080. Для проксирования запроса отправьте POST-запрос на /proxy с телом в формате JSON:
```json
{
"id": "unique-request-id",
"status": 200,
"headers": {
"Content-Type": "application/json"
},
"length": 1234
}
```
Ответ будет в формате:

```json
{
"id": "unique-request-id",
"status": 200,
"headers": {
"Content-Type": "application/json"
},
"length": 1234
}
```
