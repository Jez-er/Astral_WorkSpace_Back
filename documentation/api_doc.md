#### Пример
Описание...
- `/пример/ссылки` авторизация по логину и паролю
```javascript
await ( await fetch("https://ДОМЕН/пример/ссылки", {
  "method": "POST",
  (Метод запроса)
  "body": JSON.stringify({
    "username": "Gravita",
    "password": "1111"
    (Тут все параметры что передаються в тело запроса)
  }),
  "headers": {
    "Content-Type": "application/json"
    (Тут все хедеры что нужны)
  }
}) ).json()
```

------

#### Singup
Описание...
- `/api/public/signup` регистрация по логину, паролю и почте(что бы потом можно было войти по почте или по логину)
```javascript
await ( await fetch("https://localhost:3001/api/public/signup", {
  "method": "POST",
  "body": JSON.stringify({
    "name": "Gravita",
    "email": "gravita@gmail.ru",
    "password": "1111"
  }),
  "headers": {
    "Content-Type": "application/json"
    (Тут все хедеры что нужны)
  }
}) ).json()
```

------

#### Login
Описание...
- `/api/public/login` авторизация по почте и паролю (По логину не сделал, не уверен, что надо)
```javascript
await ( await fetch("https://localhost:3001/api/public/login", {
  "method": "POST",
  (Метод запроса)
  "body": JSON.stringify({
    "email": "gravita@gmail.ru",
    "password": "1111"
  }),
  "headers": {
    "Content-Type": "application/json"
    (Тут все хедеры что нужны)
  }
}) ).json()
```


Есть еще 2 апи, но они нужны для отправки токена и проверки(никакого осуществимого функционала), есть ли юзер в базе. 