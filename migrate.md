## マイグレーション up
```bash
migrate -database 'mysql://root:password@tcp(db:3306)/chat' -path migrations/ up
```