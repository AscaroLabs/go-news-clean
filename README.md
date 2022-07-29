# API для сервиса просмотра новостей

## Структура проекта

```sh
├── cmd 
│   └── app
│       └── main.go # входная точка приложения 
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── google # зависимости для генерации кода
│   └── api
│       ├── annotations.proto
│       └── http.proto
├── go.sum
├── internal
│   ├── adapters
│   │   ├── handlers # контроллеры
│   │   │   └── v1
│   │   │       └── grpc
│   │   │           |   # проверка работы (*)
│   │   │           ├── contentCheckService.go
│   │   │           |   # контроллеры для 
│   │   │           |   # взаимодействия с новостями
│   │   │           ├── newsServiceServer.go
│   │   │           |   # контроллеры для 
│   │   │           |   # взаимодействия тегами
│   │   │           ├── tagServiceServer.go
│   │   │           |   # функия для 
│   │   │           |   # регистрирования сервис-серверов
│   │   │           └── registrar.go
│   │   └── postgres # postgres-реализации репозиториев
│   │       ├── access
│   │       │   └── repository.go
│   │       ├── files
│   │       │   └── repository.go
│   │       ├── news
│   │       │   └── repository.go
│   │       └── tags
│   │           └── repository.go
│   ├── app
│   │   └── app.go # инициализация необходимых частей
│   ├── domain
│   │   ├── entity
│   │   │   │   # сущности для работы с правами доступа
│   │   │   ├── access
│   │   │   │   ├── access.go
│   │   │   │   ├── dto.go
│   │   │   │   └── repository.go
│   │   │   │   # сущности для работы с файлами
│   │   │   ├── files
│   │   │   │   ├── dto.go
│   │   │   │   ├── files.go
│   │   │   │   └── repository.go
│   │   │   │   # сущности для работы с новостями
│   │   │   ├── news
│   │   │   │   ├── dto.go
│   │   │   │   ├── news.go
│   │   │   │   └── repository.go
│   │   │   │   # сущности для работы с тегами
│   │   │   └── tags
│   │   │       ├── repository.go
│   │   │       └── tags.go
│   │   └── usecase # определения сервисов и их методов
│   │       ├── access
│   │       │   └── access.go
│   │       ├── files
│   │       │   └── files.go
│   │       ├── news
│   │       │   └── news.go
│   │       └── tags
│   │           └── tags.go
│   ├── gateways
│   │   │   # grpc сервер
│   │   ├── grpc.go
│   │   │   # http сервер (обертка)
│   │   └── http.go
│   └── proto
│       │   # методы для преобразования protobuf
│       │   # структур в enitiy структуры 
│       ├── convertion.go
│       ├── news_grpc.pb.go
│       ├── news.pb.go
│       ├── news.pb.gw.go
│       ├── news.proto
│       └── news.swagger.json
├── Makefile # для удобной сборки
├── pkg
│   ├── auth # пакет отвечает за авторизацию
│   │   └── auth.go
│   ├── client # создание клиентов к сторонним приложениям
│   │   └── postgres.go
│   └── env # чтение переменных окружения
│       └── env.go
└── README.md
```

## Пример реализации сервиса

Сервисы реализуютя согласно чистой архитектуре

- [Clean Architecture with GO](https://manakuro.medium.com/clean-architecture-with-go-bce409427d31)
- [Go Advanced](https://www.youtube.com/playlist?list=PLP19RjSHH4aENxkai8lzF0ocA4EZyS0vn)

### Рассмотрим вариант реализации функционала работы с новостями

#### Начинаем с слоя сущностей:

```go
// /internal/domain/entity/news/news.go

import (
	".../internal/domain/entity/files"
	".../internal/domain/entity/tags"
)

type News struct {
	Id uuid.UUID
	Title string
	Author string
	Active bool
	ActiveFrom time.Time
	Text string
	TextJson string
	UserId uuid.UUID
	Tags []*tags.Tag
	FilesInfo []*files.File
	IsImportant bool
}
```

Сущность `News` позволит реализовать необходимую бизнес-логику.

Почему бы не использовать сгенерированную структуру `pb.NewsObject`?

*Используя структуру для работы с gRPC мы нарушаем порядок вложения зависимостей, что противоречит чистой архитектуре, а так же бизнес-логика нашего приложения начинает опираться на сгенерированные protobuf'ом структуры, из-за чего возникнут серьёзные проблемы при попытке перейти с gRPC на какую-нибудь другую технологию*

#### Далее определим репозиторий

```go
// /internal/domain/entity/news/repository.go

type NewsRepository interface {
	GetNews(GetDTO, uuid.UUID) ([]*News, error)
	GetOne(id uuid.UUID) (*News, error)
	GetOneBySlug(objectSlug string) (*News, error)
	Create(cdto CreateDTO) error
	Update(udto UpdateDTO) error
	Delete(ddto DeleteDTO) error
}
```

Репозиторий представляет собой интерфейс на методы которого мы будем оприраться при реализации бизнес-логики. Идея репозитория состоит в том, что в сервис, в котором будет реализована бизнес-логика (*см. ниже*), мы передадим интерфейс, а не конкретную реализацию, тем самым при реализации бизнес-логики мы сможем использовать эти методы не задумываясь, каким образом они реализованы (*дальше это станет понятнее*).

Методы репозитория принимают структуры-DTO они нужны для того, что бы не обращаться к структурам из внешних слоев и не нарушать правило зависимостей. Эти структуры определены в `/internal/domain/entity/news/dto.go`

Сущности вместе с репозиторием и DTO находятся в слое **Entities**

#### Подходим к реализации сервиса

Сервисы являютя структурой, в методах которой находится основная бизнес-логика (слой **Use Cases**)

```go
// /internal/domain/usecase/news/news.go
type NewsService struct {
	newsRepository news.NewsRepository // репозиторий, переданный по интерфейсу
	tagService     *tags.TagService // сервис тегов
	fileService    *files.FileService // сервис новостей
	accessService  *access.AccessService // сервис проверки достука
}

// Конструктор
// В app.go мы сзодадим объект структуры, которая будет
// реализовывать интерфейс репозитория новостей и передадим в конструктор 
// вместе с инстансами других сервисов
func NewNewsService(nr news.NewsRepository, ts *tags.TagService, as *access.AccessService, fs *files.FileService) *NewsService {
	return &NewsService{
		newsRepository: nr,
		tagService:     ts,
		accessService:  as,
		fileService:    fs,
	}
}
```

Теперь можем определять методы, реализующие бизнес логику, например `GetNews` и `Create`

```go
// /internal/domain/usecase/news/news.go
func (ns *NewsService) GetNews(gdto news.GetDTO, adto access_entity.AccessDTO) ([]*news.News, error) {

	// Здесь мы обращаемся к методу репозитория,
	// мы не знаем как он будет работать, но т.к. он реализует
	// интерфейс репозитория, то вернет то, что нам нужно
	news_list, err := ns.newsRepository.GetNews(gdto, adto.Id)
	if err != nil {
		return nil, err
	}

	// Здесь мы обращаемя в сторонним сервисам, а именно к сервису тегов
	// и к сервису файлов
	for _, news_element := range news_list {
		tags_list, _ := ns.tagService.GetByNews(news_element.Id)
		news_element.Tags = tags_list
	}
	for _, news_element := range news_list {
		files_list, _ := ns.fileService.GetByNews(news_element.Id)
		news_element.FilesInfo = files_list
	}
	return news_list, nil
}

func (ns *NewsService) Create(cdto news.CreateDTO, adto access_entity.AccessDTO) error {
	// Создавать новости могут не все пользователи, поэтому нужно обратится
	// к сервису проверки доступа через AccessDTO, который передается в
	// параметры функции
	err := ns.accessService.CanCreateNews(adto)
	if err != nil {
		return err
	}

	// И снова обращаемся к репозиторию через DTO
	err = ns.newsRepository.Create(cdto)
	if err != nil {
		return err
	}
	return nil
}
```

Таким образом, осталось просто реализовать репозиторий

### Репозиторий

У нас может быть несколько реализаций интерфейса репозитория, например на основе БД (например, `PostgresQL`), на основе кешей (например, через `Redis`).

Репозиторий файлов можер быть реализован также через БД, через S3 хранилище(`Minio`) или даже просто через память. При реализации бизнес-логики нам это не важно. 

Например реализации репозитория на основе PostgresQL будут лежать в папке `/internal/adapters/postgres/`