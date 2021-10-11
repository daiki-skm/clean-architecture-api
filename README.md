# clean-architecture-api

## Description
This is an example of implemention of Clean Architecture in Golang projects.

This project has 4 layer :
- Infrastructure Layer (Frameworks & Drivers)
- Interface Layer (Interface Adapters)
- Usecase Layer (Application Business Rules)
- Domain Layer (Enterprise Business Rules)

## Run the Applications
This projects used a firestore. You must set the firestore.
The following [Golang入門 Firestore導入篇](https://rightcode.co.jp/blog/information-technology/golang-introduction-firestore)

```
# Setting
$ go mod tidy

# Prepare firestore account, and setting
$ path/to/serviceAccount.json

# Move to directory
$ cd application

# Run the application
$ go run main.go

# Execute the call
$ http://localhost:1323/users
```

## References
- @nrslib, [Qiita 実装クリーンアーキテクチャ](https://qiita.com/nrslib/items/a5f902c4defc83bd46b8)
- Robert C. Martin (Uncle Bob), [The Clean Code Blog](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Golang入門](https://rightcode.co.jp/blog/information-technology/golang-introduction-environment-1)
