DATA='{"state": ''}'

$ echo '{"type": "note", "title": "'"$TR_TORRENT_NAME"'", "body": "'"$TR_TORRENT_NAME completed"'."}'

curl -i -X POST \
-H "Authorization: Bearer $GITHUB_TOKEN" \
-H "Content-Type: application/json" \
--data '{"state":"$STATUS_STATE", \
         "context": "prombench-update", \
         "description": "prombench-update", \
         "target_url": "https://jhajdh/${{ github.sha }}"}' \
"https://api.github.com/repos/$GITHUB_REPOSITORY/statuses/$LAST_COMMIT_SHA"
