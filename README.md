# KLIK DIGITAL API TEST


## Installation

### On Local machine (DATABASE)
 - Import <code>db_kdigital.sql</code> to mysql database . Give it name <code>'kDigitalTest'</code>;
 - Database configuration located on file <code>config.json</code>
 - <code>go build .</code> to build source 
 - Open the source code, open terminal and run <code>go run main.go</code>

### On Local machine (REDIS)
 - Download and install redis in your OS
 - Run redis server <code>redis_server</code>
 - To check whether your redis is running, open <code>redis-cli</code> and type command <code>ping</code>

## Available API's
BaseUrl : localhost:12345

### User Register API
 - Endpoint: [BaseUrl]/users/register
 - Method  : POST
 - Request Body :
  ```javascript
  { 
    "username": "<string>",
    "password": "<string>",
    "email":"<string>",
    "phone_number":"<string>"
  }
  ```
 - Password must be at least 6 character or more
 - Example:
 - - Request Body:
     ```javascript
        {
        	 "username": "buena123",
             "password": "buena123",
             "email":"buena134@gmail.com",
             "phone_number":"081234351231"
        }
     ```
 - - Response Body:    
     ```javascript
        {
         "data": {
                 "username": "buena123",
                 "email": "buena134@gmail.com",
                 "phone_number": "08123435",
                 "uuid": "02381104-aaee-4863-86f8-3d3688a1d6f0"
             },
        "message": "OK",
        "status_code": 200
        }
     ```


### User Login API

 - Endpoint  : [BaseUrl]/users/login
 - Method    : POST
 - Request Body:
 ```javascript
{
    "username":"<string>",
    "password":"<string>"
}
```
 - Example:
 - - Request Body:
   ```  javascript
         {
             "username":"buena123",
             "password":"buena123"    	
         }
     ```
 - - Response Body: 
      ```javascript
        {
            "data": {
                "uuid": "02381104-aaee-4863-86f8-3d3688a1d6f0"
            },
            "message": "OK",
            "status_code": 200
        }
        ```
 
 ### Create Transaction
 - Endpoint : [BaseUrl]/create/transactions
 - Method   : POST
 - Request Body:
  ```javascript
 {
    "method":"<string>",
    "uuid":"<string>"
}
  ```
 - Example:
 - - Request Body:
     ```javascript
        {
        	"method":"GPY",
            "uuid":"02381104-aaee-4863-86f8-3d3688a1d6f0"
        }
     ```
 - - Response Body
     ```javascript
        {
            "data": {
                "transaction_number": "GPY1603671712a1d6f0"
            },
            "message": "OK",
            "status_code": 200
        }
```
