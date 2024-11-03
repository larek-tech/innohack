# Решение команды misis banach space | Система аналитики финансовых электронных документов RAG-инструментами

Система обработки и построения аналитики финансовых электронных документов содержит 2 микросервиса и реализует функционал в виде построения отчетных графиков и чат-бота.

## Инструкция по запуску

_Для запуска необходимы утилиты make и docker-compose._

## Архитектура

![Архитектура](./images/architecture.png)

## Работа с данными

## Аналитика

## api/логика

Сервис, реализующий бизнес-логику и взаимодействие с сервисом аналитики.
Предоставляет API управления сессиями пользователя в чате, для возможности удобного сохранения истории запросов и ответов.
Все запросы и ответы пользователя сохраняются в базе данных PostgreSQL.

Взаимодействие с сервисом аналитики происходит через протокол gRPC.
[Интерфейс](./proto/analytics/analytics.proto) реализует 2 метода:

- `GetCharts` - построение графиков по финансовым документов.
- `GetDescriptionStream` - получение ответов на запросы в чате, с генерацией по токенам.

## Стек технологий

### Backend

- Golang: использован для реализации сервиса с бизнес-логикой.
- PostgreSQL: хранение запросов и ответов пользователя в различных сессиях.
- Jaeger: трассировка запросов.
- Docker: контейнеризация сервисов.

### ML/Analytics/DE

- Python: использован для реализации сервиса с аналитикой и работы с данными.
- MongoDB: хранение расчитанных графиков и агрегатов по финасовым документам.
- S3: хранилище для исходных файлов, которые обрабатываются сервисом.
- QDrant: хранение и поиск по векторам.

### Frontend

- TypeScript, JavaScript
- React: фреймворк для создания пользовательского интерфейса.

## Источники данных

- [Годовая отчетность](https://moskva.mts.ru/about/investoram-i-akcioneram/korporativnoe-upravlenie/raskritie-informacii/godovaya-otchetnost)
- [Выпуск ценных бумаг](https://moskva.mts.ru/about/investoram-i-akcioneram/korporativnoe-upravlenie/raskritie-informacii/vipusk-cennih-bumag)
- [Сообщения](https://moskva.mts.ru/about/investoram-i-akcioneram/korporativnoe-upravlenie/raskritie-informacii/soobshheniya)
- [Выпуск CFA](https://moskva.mts.ru/about/investoram-i-akcioneram/korporativnoe-upravlenie/raskritie-informacii/vypusk-cfa)
- [Существенные факты](https://moskva.mts.ru/about/investoram-i-akcioneram/korporativnoe-upravlenie/raskritie-informacii/sushhestvennie-fakti)
- [Ежеквартальные отчеты](https://moskva.mts.ru/about/investoram-i-akcioneram/korporativnoe-upravlenie/raskritie-informacii/ezhekvartalnie-otcheti)
- [Отчеты эмитента эмиссионных ценных бумаг](https://moskva.mts.ru/about/investoram-i-akcioneram/korporativnoe-upravlenie/raskritie-informacii/otchety-emitenta-emissionnyh-cennyh-bumag)
- [Инсайдерская информация ПАО МТС](https://moskva.mts.ru/about/investoram-i-akcioneram/korporativnoe-upravlenie/raskritie-informacii/insajderskaya-informacii-pao-mts)
