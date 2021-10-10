

## マイグレーション

### ファイル作成
```bash
migrate create -ext sql -dir sql/migrations -seq create_users_table
```

### 実行
```bash
migrate -database 'mysql://root:password@tcp(db:3306)/chat' -path sql/migrations up
```


## シーディング

### ファイル作成
```bash
migrate create -ext sql -dir sql/seeders -seq insert_to_users_table
```

### 実行
```bash
migrate -database 'mysql://root:password@tcp(db:3306)/chat' -path sql/seeders up
```