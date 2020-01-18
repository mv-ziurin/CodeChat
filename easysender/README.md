# easysender

## Build project

```
$ make all
$ make run
```

## Use

```
$ curl -i "http://127.0.0.1:5555/send?type=email&module=auth" -X POST --data '{"email": "inwady@gmail.com", "subject": "test", "data": "hello"}'
```