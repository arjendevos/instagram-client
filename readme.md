# Instagram Client

Instagram Client is not really a client but more a tool for myself.
Instagram Client uses the instagram web api endpoints to make requests.

## Features

- [x] Fetch profile with username
- [x] Fetch recommended accounts from account
- [x] Convert and save data to json file
- [x] Save data to database (postgresql)
- [ ] Login with username + password (if possible)
- [ ] Create web dashboard to view data (next)
- [ ] Make API endpoints
- [ ] Find out language (last 3 posts)
- [ ] More stuff

## Setup

1. Extract the `cookie` + `x-ig-app-id` from a web request on [instagram-web](https://instagram.com)
2. Copy .env.example, rename to .env and add `cookie` + `x-ig-app-id`
3. Run it with `go run main.go -username=yourusername`
