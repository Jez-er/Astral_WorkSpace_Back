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