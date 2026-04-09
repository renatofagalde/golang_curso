#!/bin/bash

START_DATE="2026-04-09"
END_DATE="2026-04-09"

start_ts=$(date -u -d "$START_DATE 00:00:00" +%s)
end_ts=$(date -u -d "$END_DATE 00:00:00" +%s)

current_ts=$start_ts

echo "START_TS: $start_ts"
echo "END_TS: $end_ts"

while [ "$current_ts" -le "$end_ts" ]; do

  current_date=$(date -u -d "@$current_ts" +%Y-%m-%d)

  echo "👉 PROCESSANDO: $current_date"

  skip_day=$((RANDOM % 8))
  if [ "$skip_day" -eq 0 ]; then
    echo "😴 $current_date -> sem commits"
    current_ts=$((current_ts + 86400))
    continue
  fi

  day_of_week=$(date -u -d "$current_date" +%u)

  if [ "$day_of_week" -ge 6 ]; then
    commits=$((RANDOM % 3))
  else
    commits=$((RANDOM % 8 + 5))
  fi

  echo "📅 $current_date -> $commits commits"

  for ((i=1; i<=commits; i++)); do
    hour=$((RANDOM % 9 + 9))
    min=$((RANDOM % 60))
    sec=$((RANDOM % 60))

    commit_date="$current_date $hour:$min:$sec"

    echo "commit $i - $current_date" >> fake.txt

    GIT_AUTHOR_DATE="$commit_date" \
    GIT_COMMITTER_DATE="$commit_date" \
    git add .

    git commit -m "commit $current_date $i" --date="$commit_date" >/dev/null
  done

  current_ts=$((current_ts + 86400))
done

echo "✅ Agora vai pegar desde janeiro"