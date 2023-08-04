# UserService

The service handles client requests.
The docker-compose file is used to build the application.
We use nginx to distribute the load between application replicas.

Endpoints:
#CreateUser
  POST http//:localhost:8082/user
  
#UpdateUSerAge
  PUT http//:localhost:8082/user
  
#DeleteUSer
  DELETE http//:localhost:8082/user
  
#CreateFriendship
  POST http//:localhost:8082/friends
  
#GetFriends
  GET http//:localhost:8082/friends

