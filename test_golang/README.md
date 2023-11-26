```
cd $name_folder
go mod init $name_folder
```

Install mux package

```
go get -u github.com/gorilla/mux

```

The Difference betwenn localhost and 127.0.0.1

[Link Here](https://www.tutorialspoint.com/difference-between-localhost-and-127-0-0-1#:~:text=The%20most%20significant%20difference%20between,look%20up%20a%20table%20somewhere.)

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

```
ioutil.ReadAll -> io.ReadAll
ioutil.ReadFile -> os.ReadFile
ioutil.ReadDir -> os.ReadDir
// others
ioutil.NopCloser -> io.NopCloser
ioutil.TempDir -> os.MkdirTemp
ioutil.TempFile -> os.CreateTemp
ioutil.WriteFile -> os.WriteFile
```




