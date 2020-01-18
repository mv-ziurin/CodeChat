# Добавить тестовые данные

Для проверки бэка нужно сделать GET запрос на URL/api/CodeChat

1)Если хочешь потестить codechat.ru, то заходишь codechat.ru:8080

* login : inwady@gmail.com
* password : password

выполняешь скрипты ниже, обращение к апи делаешь по api.codechat.ru

2)Если хочешь у себя поднять бэк и потестить на локальной базе, то заходишь в appseting.json, меняешь имя хоста базы на localhost, делаешь Update-Database и выполняешь скрипты ниже

```

INSERT INTO public."users" VALUES('Webber1580', 'Webber@gmail.com', '18138372FAD4B94533CD4881F03DC6C69296DD897234E0CEE83F727E2E6B1F63', '123path.png');
INSERT INTO public."users" VALUES('inwady', 'inwady@gmail.com', '18138372FAD4B94533CD4881F03DC6C69296DD897234E0CEE83F727E2E6B1F63', '123path.png');
INSERT INTO public."users" VALUES('darkmaxZ', 'darkmaxZ@gmail.com', '18138372FAD4B94533CD4881F03DC6C69296DD897234E0CEE83F727E2E6B1F63', '123path.png');
INSERT INTO public."users" VALUES('Webbik', 'Webbik@gmail.com', '18138372FAD4B94533CD4881F03DC6C69296DD897234E0CEE83F727E2E6B1F63', '123path.png');
INSERT INTO public."users" VALUES('assinfire', 'assinfire@gmail.com', '18138372FAD4B94533CD4881F03DC6C69296DD897234E0CEE83F727E2E6B1F63', '123path.png');

INSERT INTO public."Chats" VALUES(1, 'BMSTU');
INSERT INTO public."Chats" VALUES(2, 'MSU');
INSERT INTO public."Chats" VALUES(3, 'HACKER_CHAT');
INSERT INTO public."Chats" VALUES(4, 'BlueOyster');
INSERT INTO public."Chats" VALUES(5, 'SlavKings');

INSERT INTO public."UserChats" VALUES(1, 'Webber1580', 1);
INSERT INTO public."UserChats" VALUES(2, 'Webber1580', 3);
INSERT INTO public."UserChats" VALUES(3, 'Webber1580', 4);
INSERT INTO public."UserChats" VALUES(4, 'darkmaxZ', 1);
INSERT INTO public."UserChats" VALUES(5, 'darkmaxZ', 3);
INSERT INTO public."UserChats" VALUES(6, 'darkmaxZ', 4);
INSERT INTO public."UserChats" VALUES(7, 'darkmaxZ', 5);
INSERT INTO public."UserChats" VALUES(8, 'inwady', 1);
INSERT INTO public."UserChats" VALUES(9, 'inwady', 2);
INSERT INTO public."UserChats" VALUES(10, 'inwady', 3);
INSERT INTO public."UserChats" VALUES(11, 'inwady', 4);
INSERT INTO public."UserChats" VALUES(12, 'inwady', 5);
INSERT INTO public."UserChats" VALUES(13, 'assinfire', 4);


INSERT INTO public."CodeChats" VALUES(1, 'CodeBMSTU', 1);
INSERT INTO public."CodeChats" VALUES(2, 'CodeBlueOyster', 4);
INSERT INTO public."CodeChats" VALUES(3, 'CodeHACKER_CHAT', 3);
INSERT INTO public."CodeChats" VALUES(4, 'CTF', 3);

INSERT INTO public."Messages" VALUES(1, 'Webber1580', 4, 'Whats up!', '2016-06-22 19:10:25');
INSERT INTO public."Messages" VALUES(2, 'darkmaxZ', 4, 'Hello', '2016-06-22 19:10:35');
INSERT INTO public."Messages" VALUES(3, 'Webber1580', 4, 'What about go to the party this week', '2016-06-22 19:10:39');
INSERT INTO public."Messages" VALUES(4, 'darkmaxZ', 4, 'Sounds good. I am in', '2016-06-22 19:10:45');
INSERT INTO public."Messages" VALUES(5, 'Webber1580', 4, 'Right decision)))', '2016-06-22 19:10:55');
INSERT INTO public."Messages" VALUES(6, 'darkmaxZ', 4, 'See u there', '2016-06-22 19:10:59');
```

# Удалить тестовые данные

```
DELETE FROM public."Messages" * ;
DELETE FROM public."CodeChats" * ;
DELETE FROM public."UserChats" * ;
DELETE FROM public."Chats" * ;
DELETE FROM public."users" * ;
```

### Scaffold DB

```
Scaffold-DbContext "Host=codechat.ru;Port=5432;Database=codechat;Username=postgres;Password=postgres" Npgsql.EntityFrameworkCore.PostgreSQL -o ScaffoldModels
```

### Build and run the Docker image

```
docker build -t CodeChatApp .
docker run -d -p 5000:5000 --name codeChat CodeChatApp
```

# Описание API

### Статус Коды
```
* 20000 - успешный запрос
* 40001 - нет на бэкенде функции с таким именем
* 40002 - ошибка валидации
* 40003 - неверно переданы параметры для функции
* 40004 - ошибка, связанная с обновлением данных в бд 
```

### Получение списка чатов и код чатов GetChats

* Запрос

```
{
	"method" : "GetChats",
	"token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Ikl2YW4iLCJlbWFpbCI6ImFsZ3lzLmlldmxldkBnbWFpbC5jb20iLCJpYXQiOjE1MzI4ODMxNDgxMjMsInByb2plY3QiOiJhcGkifQ.uzox1eOWLF31TlCfWVinFUCkK4lAvBW-5bNo3gfHE9A",
	"params" : {}
}
```

* Ответ

```
{
    "status": 20000,
    "result": {
        "channels": [
            {
                "chatId": 4,
                "name": "BlueOyster",
                "codeChats": [
                    {
                        "codeChatId": 2,
                        "mainChatName": "BlueOyster",
                        "name": "CodeBlueOyster"
                    }
                ]
            },
            {
                "chatId": 1,
                "name": "BMSTU",
                "codeChats": [
                    {
                        "codeChatId": 1,
                        "mainChatName": "BMSTU",
                        "name": "CodeBMSTU"
                    }
                ]
            },
            {
                "chatId": 3,
                "name": "HACKER_CHAT",
                "codeChats": [
                    {
                        "codeChatId": 3,
                        "mainChatName": "HACKER_CHAT",
                        "name": "CodeHACKER_CHAT"
                    },
                    {
                        "codeChatId": 4,
                        "mainChatName": "HACKER_CHAT",
                        "name": "CTF"
                    }
                ]
            }
        ]
    }
}
```

### Получение истории сообщений GetMessageHistory

* Запрос

```
{
	"method" : "GetMessageHistory",
	"token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Ikl2YW4iLCJlbWFpbCI6ImFsZ3lzLmlldmxldkBnbWFpbC5jb20iLCJpYXQiOjE1MzI4ODMxNDgxMjMsInByb2plY3QiOiJhcGkifQ.uzox1eOWLF31TlCfWVinFUCkK4lAvBW-5bNo3gfHE9A",
	"params" : {
		"chatId" : "3"
	}
}
```

* Ответ

```
{
    "status": 20000,
    "result": {
        "messages": [
            {
                "id": 1,
                "userName": "Webber1580",
                "chatId": 4,
                "text": "Whats up!",
                "time": "2016-06-22T19:10:25"
            },
            {
                "id": 2,
                "userName": "darkmaxZ",
                "chatId": 4,
                "text": "Hello",
                "time": "2016-06-22T19:10:35"
            },
            {
                "id": 3,
                "userName": "Webber1580",
                "chatId": 4,
                "text": "What about go to the party this week",
                "time": "2016-06-22T19:10:39"
            },
            {
                "id": 4,
                "userName": "darkmaxZ",
                "chatId": 4,
                "text": "Sounds good. I am in",
                "time": "2016-06-22T19:10:45"
            },
            {
                "id": 5,
                "userName": "Webber1580",
                "chatId": 4,
                "text": "Right decision)))",
                "time": "2016-06-22T19:10:55"
            },
            {
                "id": 6,
                "userName": "darkmaxZ",
                "chatId": 4,
                "text": "See u there",
                "time": "2016-06-22T19:10:59"
            }
        ]
    }
}
```

### Добавить чат PostChat

* Запрос

```
{
	"method" : "PostChat",
	"token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Ikl2YW4iLCJlbWFpbCI6ImFsZ3lzLmlldmxldkBnbWFpbC5jb20iLCJpYXQiOjE1MzI4ODMxNDgxMjMsInByb2plY3QiOiJhcGkifQ.uzox1eOWLF31TlCfWVinFUCkK4lAvBW-5bNo3gfHE9A",
	"params" : {
		"name" : "Sharaga"
	}
}
```

* Ответ

```
{
    "status": 20000,
    "result": {
        "chatId": 6
    }
}
```

### Добавить Кодчат PostCodeChat

* Запрос

```
{
	"method" : "PostChat",
	"token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Ikl2YW4iLCJlbWFpbCI6ImFsZ3lzLmlldmxldkBnbWFpbC5jb20iLCJpYXQiOjE1MzI4ODMxNDgxMjMsInByb2plY3QiOiJhcGkifQ.uzox1eOWLF31TlCfWVinFUCkK4lAvBW-5bNo3gfHE9A",
	"params" : {
		"name" : "LOLCodeChat",
		"chatId" : 4
	}
}
```

* Ответ

```
{
    "status": "200",
    "result": {
        "codeChatId" : 5
    }
}
```

### Добавить Пользователя в чат AddUserToChat

* Запрос

```
{
	"method" : "AddUserToChat",
	"token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Ikl2YW4iLCJlbWFpbCI6ImFsZ3lzLmlldmxldkBnbWFpbC5jb20iLCJpYXQiOjE1MzI4ODMxNDgxMjMsInByb2plY3QiOiJhcGkifQ.uzox1eOWLF31TlCfWVinFUCkK4lAvBW-5bNo3gfHE9A",
	"params" : {
		"username" : "darkmaxZ",
		"chatId" : 4
	}
}
```

* Ответ

```
{
    "status": "200",
    "result": User was successfully added to the chat"
}
```

### Покинуть чат LeaveChannel

* Запрос

```
{
	"method" : "LeaveChannel",
	"token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Ikl2YW4iLCJlbWFpbCI6ImFsZ3lzLmlldmxldkBnbWFpbC5jb20iLCJpYXQiOjE1MzI4ODMxNDgxMjMsInByb2plY3QiOiJhcGkifQ.uzox1eOWLF31TlCfWVinFUCkK4lAvBW-5bNo3gfHE9A",
	"params" : {
		"chatId" : 4
	}
}
```

* Ответ

```
{
    "status": "200",
    "result":"User has successfuly leaved the channel"
}
```
