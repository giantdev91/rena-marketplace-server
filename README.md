# Rena NFT Marketplace Server

See [API Spec](./api.md) (modified from [RealWorld API Spec](https://github.com/gothinkster/realworld/tree/master/api) to simplify)  

API Server technology stack is  

- Server code: `golang`
- REST Server: `gin`
- Database: `PostgreSQL` with `golang-migrate` to migrate  
- ORM: `gorm v2`  
- Dependency Injection: `fx`  
- Unit Testing: `go test` and `testify`
- Configuration management: `cobra`

---  

Commands to run the app
- make build
- make factory (Command to add initial nft items to database)
- make start
- make migrate
- make test
- make lint
