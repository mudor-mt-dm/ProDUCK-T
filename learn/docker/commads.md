Команды:

1. Команда копирования файла в контейнер который запущен
```
docker cp ./from_root_system/source.sql container_name:/where_in_container/copy_of.sql
docker cp ./learn/sql/demo-big-20170815.sql db:/data/demo.sql
```
2. Команда перехода в терминал контейнера, где ему можно отдавать команды:
```
docker exec -it ContName bash
docker exec -it db bash
```
3. Запуск команды постгре в терминале
```
psql -f demo.sql -U postgres
psql -f /data/demo.sql -U postgres
```