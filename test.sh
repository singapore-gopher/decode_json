#! /bin/bash

call(){
  METHOD=$1
  URL=$2
  DATA=$3


  if [[ -n $DATA ]]; then
    echo "curl -X $METHOD '$URL' -d '$DATA'"
    curl -X $METHOD -H 'Content-type: application/json' "$URL" -d "$DATA" -v
  else
    echo "curl -X $METHOD '$URL'"
    curl -X $METHOD -H 'Content-type: application/json' "$URL" -v
  fi
  echo ""
  echo "-----------------------------------"
}

# keep it simple, do manual check

test_register(){
  call 'POST' 'localhost:4000/register.json' '{}'
  call 'POST' 'localhost:4000/register.json' '{"name":"marvin"}' # good
  call 'POST' 'localhost:4000/register.json' '{"name":"marvin"}'
}

# stage 1
#			{`{"first":5,"second":10}`, `{"sum":15}`},
#			{`{"first":7,"second":234}`, `{"sum":241}`},
#			{`{"first":9,"second":8}`, `{"sum":17}`},
#			{`{"first":14,"second":84}`, `{"sum":98}`},

test_stage1(){
  call 'POST' 'localhost:4000/register.json' '{"name":"vorgon"}'
  call 'POST' 'localhost:4000/stage1/submit.json' '{"team":"vorgon","solutions":[{"sum":15},{"sum":241},{"sum":17},{"sum":9811}]}' # good (first 3 correct)
  call 'POST' 'localhost:4000/stage1/submit.json' '{"team":"vorgon","solutions":[{"sum":15},{"sum":241},{"sum":17}]}'
  call 'POST' 'localhost:4000/stage1/submit.json' '{"solutions":[{"sum":15},{"sum":241},{"sum":17},{"sum":9811}]}'
  call 'POST' 'localhost:4000/stage1/submit.json' '{"team":"random","solutions":[{"sum":15},{"sum":241},{"sum":17},{"sum":9811}]}'
}

test_leaderboard(){
  call 'GET' 'localhost:4000/leaderboard.json'
}

test_register
test_stage1
test_leaderboard
