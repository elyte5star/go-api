API with Go Fiber.

## Project Setup

 - Clone project with

  ```
  git clone git@github.com:elyte5star/go-api.git
  ```
  - Create file .env and set values for the environment variables below:

  ```
    MYAPP_DEBUG=false
    MYAPP_SERVICENAME="Elyte Realm API"
    MYAPP_READTIMEOUT="60"
    MYAPP_URL="http://localhost:8080"
    MYAPP_DOC="./docs/swagger.json"
    MYAPP_SERVICEPORT="8080"
    MYAPP_CLIENTORIGIN="http://localhost:3000"
    MYAPP_VERSION="v1.0.1"
    MYAPP_CORSORIGINS="https://demo-elyte.com,https://demo-elyte.com,http://*.demo-elyte.com,http://localhost,http://localhost:3000,http://localhost:9000"
    MYAPP_SMTPSERVER="smtp.gmail.com"
    MYAPP_SMTPUSERNAME="*******"
    MYAPP_SMTPPASSWORD="********"
    MYAPP_JWTSECRETKEY="**********"
    MYAPP_JWTEXPIREMINUTESCOUNT="3600"
    MYAPP_SMTPPORT="587"
    MYAPP_SUPPORTEMAIL="xxxxxx@you.com"
    MYSQL_HOST="localhost"
    MYSQL_USER="userExample"
    MYSQL_PORT="3306"
    MYSQL_DATABASE="elyteGO"
    MYSQL_ROOT_PASSWORD="54321"
    MYSQL_PASSWORD="54321"
  ```
  ```

- Use swag init to converts Go annotations to Swagger Documentation

```

```
- Expose the application to MYSQL Database at port 3306

```
```
- Run the Application

```
 go run *.go

 ```