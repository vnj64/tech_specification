# tech_specification

Запуск БД через Docker:
```bash
docker compose -f docker-compose.yaml --build db
```

Выбрана подобная архитектура с перспективой расширения функционала приложения. Если такая необходимость возникнет - масштабировать приложение не составит труда.

Значения POSTGRES_USER и POSTGRES_PASSWORD необходимо зашифровать в main по аналогии с:
```go
fmt.Println(encryptor.Encrypt("string"))
panic(1)
```