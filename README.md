# GO Web API REST - Digital House

[![GoDoc](https://godoc.org/github.com/qiangxue/go-rest-api?status.png)](http://godoc.org/github.com/qiangxue/go-rest-api)

## Descripción
API REST desarrollada en Go que implementa los principios [SOLID](https://en.wikipedia.org/wiki/SOLID)
y [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). El proyecto está estructurado en capas bien definidas (Repository, Service, Controller) para garantizar la separación de responsabilidades y facilitar el mantenimiento. Se enfoca en buenas prácticas como uso de contexts, manejo de errores HTTP estandarizados, middleware para autenticación y autorización.

## Características Principales
- Principios SOLID para un diseño orientado a objetos mantenible y escalable
- Clean Architecture para separar las capas de la aplicación
- Diseño basado en interfaces para facilitar testing y desacoplamiento
- Patrón Repository para abstracción de la capa de datos
- Service Layer para encapsular la lógica de negocio

## Tecnologías y Patrones
- Go 1.13+
- Gin Gonic para routing y middleware
- PostgreSQL como base de datos
- JWT para autenticación
- Bootstrap Package para configuración inicial
- Transport Package para la capa HTTP
- Patrón Repository
- Contexts y manejo de errores

## Estructura del Proyecto
```
├── cmd/
│   └── main.go               # Punto de entrada de la aplicación
├── internal/                 # Código interno de la aplicación
├── pkg/
│   ├── bootstrap/           # Inicialización y configuración
│   ├── handler/             # Manejadores HTTP
│   │   └── user.go         # Handlers de usuarios
│   └── transport/          # Capa de transporte
│       ├── gin.go          # Configuración de Gin
│       └── http.go         # Configuración HTTP general
├── .env                     # Variables de entorno
├── .env.example             # Ejemplo de variables de entorno
├── .gitignore
├── docker-compose.yml       # Configuración de Docker
├── go.mod                   # Dependencias
└── go.sum                   # Checksums de dependencias
```


<!-- > [!IMPORTANT]
> Para comenzar, asegúrate de tener Go instalado en tu sistema. Puedes verificar tu versión con el comando `go version` en la terminal. Se recomienda Go 1.13 o superior para este proyecto. -->

## Documentación de la API

### Endpoints disponibles:

- `GET /healthcheck`: Verificación del estado del servicio
- `POST /v1/login`: Autenticación de usuarios
- `GET /v1/products`: Obtener lista de productos
- `GET /v1/products/:id`: Obtener producto por ID
- `POST /v1/products`: Crear nuevo producto
- `PATCH /v1/products/:id`: Actualizar producto existente
- `DELETE /v1/products/:id`: Eliminar producto

##  Tests
Para ejecutar los tests:

```bash
go test ./...
```
<br>
<br>

---
---
> [!NOTE]
> Este proyecto fue desarrollado gracias a los conocimientos adquiridos en el curso "Fundamentos de GO" de Digital House, donde se aprendieron conceptos fundamentales de Go y el desarrollo de APIs REST.
