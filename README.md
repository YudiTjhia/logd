#logd
logd is a server to do logging and access the log rows.

##logd functions
    - create log, 
    - find logs, 
    - find single log
      
##logd features     
    - multi domain/ip/ports
    - multi accounts
    - multi logins
    - app / serv / method tracking
    - tag enabled
    - differentiate log type: access, error, data
        
        
#logc 
logc is client sdk to connect to logd. 
current is built in go language
 
##logc functions

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
    
    
