# tech_specification

Запуск БД через Docker:
```bash
docker compose -f docker-compose.yaml --build db
```

Запуск приложения через Docker (лучше запускать целиком папку через build в IDE, запуск сильно быстрее):
```bash
docker compose -f docker-compose.yaml --build app
```

Выбрана подобная архитектура с перспективой расширения функционала приложения. Если такая необходимость возникнет - масштабировать приложение не составит труда.