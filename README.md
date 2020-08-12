# logd
logd is a server to do logging and access the log rows.

logging is very important when we already deploy our application
into client environments. we are limited to access the server. 
Thus, we need to define protocols that open our access either 
it is dev staging, uat staging or even production staging.

I try to create simple logging server and the client sdk to access.
the code is still baby. the future is seen.

## logd functions
    - create log, 
    - find logs, 
    - find single log
      
## logd features     
    - multi domain/ip/ports
    - multi accounts
    - multi logins
    - app / serv / method tracking
    - tag enabled
    - differentiate log type: access, error, data

## libraries
- please check related libraries that used.       
        
# logc 
logc is client sdk to connect to logd. 
current is built in go language
 
## logc functions

### NewLogc()
- to initialize logc
 
### Accs(...)
- to do logging with log_typ = "access"

### Err(...)
- to do logging with log_typ = "error"

### Data(...)
- to do logging with log_typ = "data"

### FindLogs(...)
- to query rows of logging

### FindLog(...)
- to query single rows of logging
    
    
