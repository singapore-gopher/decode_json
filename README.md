# decode_json
JSON Decode challenge

# how-to

## Register your team
### POST /register.json

```bash
> POST /register.json HTTP/1.1
> User-Agent: curl/7.35.0
> Host: localhost:4000
> Accept: */*
> Content-type: application/json
> Content-Length: 17
>

< HTTP/1.1 201 Created
< Date: Wed, 08 Apr 2015 08:14:12 GMT
< Content-Length: 27
< Content-Type: text/plain; charset=utf-8
<

{"ok":true,"name":"marvin"}
```
## Get leaderboard
### GET

```bash
> GET /leaderboard.json HTTP/1.1
> User-Agent: curl/7.35.0
> Host: localhost:4000
> Accept: */*
> Content-type: application/json
>

< HTTP/1.1 200 OK
< Date: Wed, 08 Apr 2015 08:14:12 GMT
< Content-Length: 161
< Content-Type: text/plain; charset=utf-8
<

{"marvin":{},"vorgon":{"stage_1":{"attempts":1,"passed":3,"first_try":"2015-04-08T16:14:12.212291619+08:00","latest_try":"2015-04-08T16:14:12.212291619+08:00"}}}
```

## Challenge 1

```bash
> GET /stage1/data.json HTTP/1.1
> User-Agent: curl/7.35.0
> Host: localhost:4000
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Wed, 08 Apr 2015 10:25:30 GMT
< Content-Length: 109
< Content-Type: text/plain; charset=utf-8
<

{"inputs":[{"first":5,"second":10},{"first":7,"second":234},{"first":9,"second":8},{"first":14,"second":84}]}
```

```bash
> POST /stage1/submit.json HTTP/1.1
> User-Agent: curl/7.35.0
> Host: localhost:4000
> Accept: */*
> Content-type: application/json
> Content-Length: 78
>

< HTTP/1.1 200 OK
< Date: Wed, 08 Apr 2015 08:14:12 GMT
< Content-Length: 12
< Content-Type: text/plain; charset=utf-8
<

{"passed":3}
```
