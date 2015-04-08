#! /bin/bash

call(){
  METHOD=$1
  URL=$2
  DATA=$3

  echo "curl -X $METHOD '$URL' -d '$DATA'"
  curl -X $METHOD -H 'Content-type: application/json' "$URL" -d "$DATA" -v
  echo ""
  echo "-----------------------------------"
}

# keep it simple, check with your eyes

call 'POST' 'localhost:4000/register.json' '{}'
call 'POST' 'localhost:4000/register.json' '{"name":"mark"}'
call 'POST' 'localhost:4000/register.json' '{"name":"mark"}'
