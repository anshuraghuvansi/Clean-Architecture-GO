# Clean-Architecture-GO

This is a Rest Api(Gin, Go and MYSQL) project based on clean architecture.
In this project i have implemented 4 API's:
1. `signup` api(POST) 
    On successful signup this api return a JWT Web Token.
2. `signin` api(POST)
    On successful signin this api return a JWT Web Token.
3. `Get profile` api(GET)
    This api use the JWT token from the `Authorization Header` from the request to autorize and return user data.
4. `Update profile` api(PUT)
    This api use the JWT token from the `Authorization Header` from the request to autorize and update user data.

# How to Use

1. Clone this project on your machine.
2. Run MYSQL db on your local machine or change the database details in `development.env` file.
3. Then run project in the VSCode

