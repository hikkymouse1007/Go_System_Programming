# Go_System_Programming

https://www.lambdanote.com/products/go

## デバッガのセットアップ

https://qiita.com/gs1068/items/1d89d4a9bed070782298

デバッガを起動するためには、go mod でモジュールにリモートリポジトリを指定する
必要がある。
https://stackoverflow.com/questions/67306638/go-test-results-in-go-cannot-find-main-module-but-found-git-config-in-users

```
git init
git remote add origin https://github.com/hikkymouse1007/Go_System_Programming.git
go mod init　github.com/hikkymouse1007/Go_System_Programming
```

https://qiita.com/ymmy02/items/b4b38a57c024510c5210
