# decode_json
JSON Decode challenge

# how-to

## Registration

### POST http://10.0.2.235:4000/register.json
```bash
> POST /register.json HTTP/1.1
> User-Agent: curl/7.35.0
> Host: localhost:4000
> Accept: */*
> Content-type: application/json
> Content-Length: 17
>

{"name":"marvin"}

< HTTP/1.1 201 Created
< Date: Wed, 08 Apr 2015 08:14:12 GMT
< Content-Length: 27
< Content-Type: text/plain; charset=utf-8
<

{"ok":true,"name":"marvin"}
```


## Challenge 1

### The Warmup

Simply add the first and the second values in the data you get from the server. Return it in the specified format, along with your team's name.

### GET http://forgot_to_update_hostname/stage1/data.json
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

### POST http://10.0.2.235:4000/stage1/submit.json
```bash
> POST /stage1/submit.json HTTP/1.1
> User-Agent: curl/7.35.0
> Host: localhost:4000
> Accept: */*
> Content-type: application/json
> Content-Length: 78
>

{"team":"YOURTEAM","solutions":[{"sum":15},{"sum":241},{"sum":17},{"sum":9811}]}

< HTTP/1.1 200 OK
< Date: Wed, 08 Apr 2015 08:14:12 GMT
< Content-Length: 12
< Content-Type: text/plain; charset=utf-8
<

{"passed":3}
```

## Challenge2

### Faulty Sensors

You have a bunch of sensors in your smart home. Each sensor gives you exactly 4 values. You usually poll them using http GET. Some sensors started to fail, and did not send all the values anymore. Identify the failed sensors in your request to avoid a disaster!

### GET http://10.0.2.235:4000/stage1/data.json
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

{
   "sensors" : [
      {
         "id" : 2346,
         "humidity" : 40,
         "temperature" : 25
      },
      {
         "id" : 2871,
         "humidity" : 40,
         "temperature" : 25
      },
      {
         "have_smoke" : false,
         "id" : 22,
         "humidity" : 40,
         "temperature" : 25
      },
      {
         "have_smoke" : false,
         "id" : 234,
         "humidity" : 40,
         "temperature" : 25
      },
      {
         "have_smoke" : false,
         "id" : 39,
         "humidity" : 40,
         "temperature" : 25
      }
   ]
}
```

### POST http://10.0.2.235:4000/stage1/submit.json
```bash
> POST /stage1/submit.json HTTP/1.1
> User-Agent: curl/7.35.0
> Host: localhost:4000
> Accept: */*
> Content-type: application/json
> Content-Length: 78
>

{"team":"YOURTEAMNAME", "faulty":[2346,2871]}

< HTTP/1.1 200 OK
< Date: Wed, 08 Apr 2015 08:14:12 GMT
< Content-Length: 12
< Content-Type: text/plain; charset=utf-8
<

{"ok":true}
```


## Get leaderboard

### GET http://10.0.2.235:4000/leaderboard.json
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
