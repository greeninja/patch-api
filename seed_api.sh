#!/bin/bash

for i in {1..500}; do
  ## Server Name
  server=server$(( ( RANDOM % 50 )  + 1 )).testing.lab

  ## Random Date time
  days_in_front=$(shuf -i 7-11 -n 1)
  hours_in_front=$(shuf -i 0-24 -n 1)
  minutes_array=("00" "15" "30" "45")
  minute=$(shuf -e "${minutes_array[@]}" -n 1)
  #echo "Minutes = $minute"

  scheduled_update_time=$(date '+%Y-%m-%d %H' -d "+$days_in_front days +$hours_in_front hours")

  #echo "$server $scheduled_update_time:$minute:00"

  curl -X POST http://localhost:8080/patch/ --header "Content-Type: application/json" --data "{\"server\": \"$server\", \"PatchStart\": \"$scheduled_update_time:$minute:00\", \"PreCheckScheduled\": \"0\", \"PreCheckStatus\": \"\", \"PatchScheduled\": \"0\", \"Status\": \"Booked\"}"
done