TodoApiRevel
============

Todo rest api using Revel Framework and PostgreSQL

Install
-------

You need GO installed, and this package :
``` Bash
go get github.com/revel/revel
go get github.com/lib/pq
go get github.com/jinzhu/gorm
go get github.com/Unknwon/com
go get github.com/donnpebe/TodoApiRevel
```


Add the following to the end of your .bashrc/.zshrc
``` Bash
export MGT_APP_SECRET='<long secret>'
export MGT_DB_DEV_SPEC='user=your_username password=your_password dbname=tododbdev sslmode=disable'
export MGT_DB_PROD_SPEC='user=your_username password=your_password dbname=tododbprod sslmode=require'
```

Now reload the changes:
```
source ~/.bashrc
or
source ~/.zshrc
```

Generate Secret
---------------

You can generate secret with following command :
``` Bash
cd $GOPATH/src/github.com/donnpebe/todoapirevel/tools
go run secret.go
```

License
-------

This project is under the MIT License.
