# simple-microservice

# # To run the whole setup : 

# 1. Go in the website directory and open the index.html in live server.
# 2. Go in the cmd/name_service/ and run command "go run name.go".
# 3. Go in the cmd/greet_service/ and run command "go run greet.go".

## After this you can make the request and it will greet you with Hello + Your Name. The names is geeting stored in the mongodb atlas database.

## Note : there should be an env file inside the cmd/name_service/ from which your MONGODB_URI is going to get load in the name.go file.
