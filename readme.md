#This is read me file

This project for initial base of a Go-Lang Project

##Requerments

install below dependency first

```console
$ go get -u github.com/alive2212/goose/cmd/goose
```

```console
$ go get -u github.com/pressly/goose/cmd/goose
```

##Usage

###Database
this section contain 
1. migration
2. model

####Migration
To create new migration just go to `/database/migrations` address and use following command:

```console
$ goose create add_some_column sql
```

After that put something like `00001_add_users_column.sql` and use command like following:

```console
$ goose mysql "gouser:gopassword@/GoTest?parseTime=true" up
$ goose mysql "gouser:gopassword@/GoTest?parseTime=true" down
$ goose mysql "gouser:gopassword@/GoTest?parseTime=true" status
```

for more information use following address :

[Pressly Goose Document Page](https://github.com/pressly/goose "pressly/goose Github Homepage")
