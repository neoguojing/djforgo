# djforgo
This is a web administrator framework for golang web development. 
Inspired by [django](https://github.com/django/django). 

## Overview

* Provide admin pages like django do. Use [Bootstrap](http://v3.bootcss.com/).
* Support a few middleware [session,request,common].
* Support ORM ,Form,Model programing like django.

## Dependency
* go get github.com/alecthomas/log4go
* go get github.com/flosch/pongo2
* go get github.com/gorilla/context
* go get github.com/gorilla/mux
* go get github.com/gorilla/schema
* go get github.com/gorilla/sessions
* go get github.com/jinzhu/gorm
* go get github.com/jinzhu/gorm/dialects/mysql
* go get github.com/deckarep/golang-set
* go get github.com/RangelReale/osin
* go get github.com/felipeweb/osin-mysql
* go get github.com/go-sql-driver/mysql
* go get github.com/prometheus/client_golang/prometheus
* go get golang.org/x/crypto/bcrypt

## Getting Started
### Configuation
  Edit the config.json in root directory:
  * Set listenning port
  * Set the Database address
  * Set the Session 
  * Set Admin email
### Init DataBase
```bash
cd ./bin
go build
./bin -init
```
### Run
In root directory
```shell
  go build
  ./djforgo
```

#TODO
* Add object register system.
* Finish common form programing module.
* Add mail send feature and password retrieve feature.
* Finish oauth feature.

# Author
**neo**


