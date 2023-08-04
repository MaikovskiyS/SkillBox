# UserService

The service handles client requests.

The docker-compose file is used to build the application.

The nginx is used like load balancer between application replicas.
___
Endpoints:

CreateUser
  POST http//:localhost:8082/user
  
UpdateUserAge
  PUT http//:localhost:8082/user
  
DeleteUser
  DELETE http//:localhost:8082/user
  
CreateFriendship
  POST http//:localhost:8082/friends
  
GetFriends
  GET http//:localhost:8082/friends
___
