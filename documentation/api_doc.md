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

# Авторизация 

#### Singup
Описание...
- `/api/public/signup` регистрация по логину, паролю и почте(что бы потом можно было войти по почте или по логину)
```javascript
await ( await fetch("https://localhost:3001/api/public/signup", {
  "method": "POST",
  "body": JSON.stringify({
     "Name": "Test1",
      "Email": "Test2@test.com",
      "Password": "123123",
      "Displayname": "Test21",
      "Howdid": "TikTok"
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

------

#### Oauth2 Google
Описание...
- `/api/oauth/google` авторизация через google (Тестить надо на странице в браузере)
```javascript
await ( await fetch("https://localhost:3001/api/oauth/google", {
  "method": "GET",
  (Метод запроса)
  "body": ,
  "headers": {
    "Content-Type": "application/json"
    (Тут все хедеры что нужны)
  }
}) ).json()
Здесь могу перенаправлять куда угодно после входа и отправлять кучу всего
`Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
	Hd            string `json:"hd"``
```

------


#### Oauth2 GitHub
Описание...
- `/api/oauth/github` авторизация через github 
```javascript
await ( await fetch("https://localhost:3001/api/oauth/github", {
  "method": "GET",
  (Метод запроса)
  "body": ,
  "headers": {
    "Content-Type": "application/json"
    (Тут все хедеры что нужны)
  }
}) ).json()
```

------

#### Create WorkSpace
Описание...
- `/api/workspace/createSpace` Создание пространства  
```javascript
await ( await fetch("https://localhost:3001/api/workspace/createSpace", {
  "method": "POST",
  "body":JSON.stringify({
    "UserID": uuid пользователя, с аккаунта которого выполняется запрос,
    "description": "Бла бла бла",
    "title": "Name",
    "logoColor":"#ffffff"
  }), ,
  "headers": {
    "Content-Type": "application/json"
    (Тут все хедеры что нужны)
  }
}) ).json()
```

------

#### Update WorkSpace
Описание...
- `/api/workspace/updateSpace/:id` Обновление пространства
```javascript
await ( await fetch("https://localhost:3001/api/workspace/updateSpace/:id", { // ID рабочего пространства
  "method": "PUT",
  "body":JSON.stringify({
    "description": "Новый Бла бла бла",
    "title": "новый Name",
    "logoColor":"новый #ffffff"
  }), ,
  "headers": {
    "Content-Type": "application/json"
  }
}) ).json()
```

------

#### Delete WorkSpace
Описание...
- `/api/workspace/deleteSpace/:id` Удаление пространства
```javascript
await ( await fetch("https://localhost:3001/api/workspace/deleteSpace/:id", { // ID рабочего пространства
  "method": "DELETE",
  "body": ,
  "headers": {
    "Content-Type": "application/json"
  }
}) ).json()
```

------

#### Get WorkSpaces
Описание...
- `/api/workspace/getSpaces/:user_id` Запрос на все пространства
```javascript
await ( await fetch("https://localhost:3001/api/workspace/getSpaces/:user_id", { // ID пользователя
  "method": "GET",
  "body": ,
  "headers": {
    "Content-Type": "application/json"
  }
}) ).json()
```

------

#### Get WorkSpace
Описание...
- `/api/workspace/getSpace/:user_id` Запрос на 1 пространствo
```javascript
await ( await fetch("https://localhost:3001/api/workspace/getSpace/:id", { // ID рабочего пространства
  "method": "GET",
  "body": ,
  "headers": {
    "Content-Type": "application/json"
  }
}) ).json()
```

