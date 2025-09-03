# Chirpy

This is for my [Boot.dev](https://www.boot.dev/lessons/50f37da8-72c0-4860-a7d1-17e4bda5c243) course  
Learning http servers in Go

Set you .env file

```
DB_URL="postgres://<path_to_your_postges_db>?sslmode=disable"
PLATFORM="dev" if or "prod" to enable/disable reset endpoint
SECRET="<a_32_bit_random_char_string>"
POLKA_API="a_fake_but_functional_API"
```

## routes

### GET
/api/healthz -- returns 200 ok if route is good  
/api/chirps -- get all chirps or add ?{user_id} to get chirps for that user  
/api/chirps/{chirpID} -- get a one chirp by id  
/admin/metrics -- get total app hits since server was started  
GET requests do not require authentication  

### POST
/api/login -- log in if user account exists  
/api/users -- create a user  
```
{
  "email": "walt@breakingbad.com",
  "password": "123456"
}
```
Login and users don't require a token. requests should be sent with json shape above

/api/chirps -- post a chirp if logged in as a user
```
{
  "body": "I'm the one who knocks!"
}
```
posting a chirp requires a token obtained by logging in. send with json shape above.

/api/refresh -- refresh access token if you have a valid refresh token  
/api/revoke -- revoke your access token if you have a valid refresh token  
refresh and revoke require a token and take no input  

/admin/reset -- resets the db for development.  
must sep platform="dev" in .env and requires no input  


### PUT
/api/users -- change email and or password  
```
{
  "email": "walt@breakingbad.com",
  "password": "123456"
}
```
Changing email requires a token. requests should be sent with json shape above  

### POST
/api/polka/webhooks -- webhook for fictitious payment provider to upgrade account to preimium "Red" account  
requires an api from the fake payment platform Polka  

### DELETE
/api/chirps/{chirpID} delete regretful chirp  
must have a valid token for the owner of the chirp  

