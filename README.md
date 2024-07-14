# Device Auth

## 1. Description

Кластер сервисов, обеспечивающих авторизацию пользователя через мобильный телефон, как это сделано в банковских приложениях

## 2. Deployment

```
    docker-compose up -d --build
```

## 3. Quick explanation / Flow

Есть два сервиса - user-service и device-service

Первый отвечает за работу с пользователя, второй - мобильными телефонами

Для начала нужно зарегистрировать новый device:
POST localhost:8081/api/v1/auth/sign-in
```
{
    "phoneNumber" : "+xxx"
}
```
Он возвращает:
```
{
    "codeToken" : "..."
}
```
Это токен, с которым нужно будет подтвердить устройство, используя код, который должен на него придти(в нашем случае он просто будет написан в консоль сервиса, денег на телефонию нет :skull:)

Подтверждаем устройство:
PATCH localhost:8081/api/v1/auth/verify-device?code={тут код} + codeToken
После запроса придёт уже device токен, с которым далее нужно трогать остальные эндпоинты
```
{
    "token" : "..."
}
```

Установка пина:
PATCH localhost:8081/api/v1/auth/set-pin + device token
```
{
    "pinCode" : "123"
}
```
Если всё ок - 200 и пустое тело

"Привязка" юзера:
Для этого необходимо получить токен юзера, которого необходимо привязать. Это можно сделать через /login || /sign-up в user-service //TODO: use something like "bind token" instead of user token
PATCH localhost:8081/api/v1/auth/bind-user + device token
```
{
"userToken" : "..."
}
```
Если всё ок - 200 и пустое тело

Далее можно логиниться:
POST localhost:8081/api/v1/auth/login
```
{
    "pin" : "123"
}
```
И в ответ получим токен user-a, которому устройство было привязано)
```
{
    "token" : "..."
}
```

Чтобы его проверить, можно получить инфо о юзере в user-service:
GET localhost:8080/api/v1/user/me

User-service sign-Up/login:

POST localhost:8080/api/v1/auth/sign-in:
```
{
"email" : "xxx@gmail.com",
"password" : "xxx",
"phoneNumber" : "xxx"
}
```

POST localhost:8080/api/v1/auth/login:
Если указан телефон - ищем пользователя используя его, иначе по почте
```
{
"email" : "xxx@gmail.com",
"password" : "xxx",
"phoneNumber" : "xxx"
}
```

В ответ везде получим:
```
{
   "token" : "..."
}
```