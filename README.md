# masterplan-backend

 ### To compile the program you will need the _go get_ following packages
 - github.com/gin-gonic/gin
 - github.com/akshaybharambe14/go-jsonc
 - github.com/natefinch/lumberjack
 - github.com/sirupsen/logrus
 - github.com/go-sql-driver/mysql
 - github.com/dgrijalva/jwt-go

### To compile and run the program
```
go run main.go
```

### To compile the program to create and executeable
```
go build main.go
```


 If you are still unable to compile and run the code then you could use the executable provided with the name  `main`. The executable was compiled on Linux - Ubuntu 18.04 so it won't run on Windows OS.

 An SQL file is also provided which includes the structure as well as the data that is being used by this program.

Credentials needed to connect to the DB should be entered in `config.jsonc` file.
The file also includes other configurable parameters. Hopefully the key names of the json are self-explanatory.
Incase you change the port number over which the program will be running on please make changes into `target` in  `vue.config.js` in masterplan-frontend