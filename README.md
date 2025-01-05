# GO Web API REST - Digital House

[![GoDoc](https://godoc.org/github.com/qiangxue/go-rest-api?status.png)](http://godoc.org/github.com/qiangxue/go-rest-api)

## Descripción
API REST desarrollada en Go que implementa los principios [SOLID](https://en.wikipedia.org/wiki/SOLID)
y [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html). El proyecto está estructurado en capas bien definidas (Repository, Service, Controller) para garantizar la separación de responsabilidades y facilitar el mantenimiento. Se enfoca en buenas prácticas como uso de contexts, manejo de errores HTTP estandarizados, middleware para autenticación y autorización.

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
├── docker/                  # Configuraciones de Docker
│   └── mysql/
│       ├── Dockerfile      # Dockerfile para MySQL
│       └── init.sql        # Script inicial de la base de datos
│
├── internal/               # Código interno de la aplicación
│   ├── domain/            # Definición de modelos y entidades
│   │   ├── message.go
│   │   └── user.go
│   └── user/              # Módulo de usuarios
│       ├── errors.go      # Manejo de errores específicos
│       ├── user.controller.go    # Controladores
│       ├── user.repository.go    # Repositorios
│       └── user.service.go       # Servicios
│
├── pkg/                   # Paquetes compartidos
│   ├── bootstrap/         # Inicialización de la aplicación
│   │   └── bootstrap.go
│   ├── handler/          # Manejadores HTTP
│   │   └── user.go
│   └── transport/        # Capa de transporte
│       └── http.go
│
├── .env.example          # Ejemplo de variables de entorno
├── go.mod               # Dependencias del proyecto
├── go.sum               # Checksums de dependencias
└── README.md            # Documentación del proyecto
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

<br>
<br>

---
---
> [!NOTE]
> Este proyecto fue desarrollado gracias a los conocimientos adquiridos en el curso "Fundamentos de GO" de Digital House, donde se aprendieron conceptos fundamentales de Go y el desarrollo de APIs REST.
