## Пример реализации проекта с использованием temporal

Для поднятия контейнеров и запуска запустите команду: `make build`
Команда поднимет следующие поды:
- backend приложение с бизнес логикой запуска **schedules**
- 5 холодных воркеров, которые выполняют **Workflow** - **GetOnlyIP**
- 5 горячих воркеров, которые выполняют **Workflow** - **GetAddressFromIP**
- core backend temporal - шедулер, к которому подключаются воркеры и backend нашего приложения
- temporal ui
- temporal postgres db
- temporal admin backend

Для наполнения постановки задач на шедулинг:
- реализовано API: `POST localhost:8888/schedules -d '{"schedule_type": "hot" (или "cold")}'`
- Для тестовых запусков выполните команды `make cold` и/или `make hot`
- Зайдите в temporal UI по адресу `localhost:8080`, где сможете поиграться с содержимым.

Для остановки приложения: `make stop`
