# Gestión de Encuentros - API REST

Este proyecto es una API REST desarrollada en Go para la gestión de encuentros deportivos, grupos, equipos y sus relaciones en el contexto de torneos. Forma parte del proyecto YoPLAY y permite realizar operaciones CRUD (Crear, Leer, Actualizar, Eliminar) sobre las entidades relacionadas con la gestión de encuentros.

## Características

- **API RESTful**: Implementada con Gorilla Mux para el enrutamiento.
- **Base de Datos**: PostgreSQL con esquema dedicado `gestion_encuentro`.
- **CORS Habilitado**: Permite solicitudes desde cualquier origen.
- **Entidades Gestionadas**:
  - Encuentros (partidos/matches)
  - Grupos
  - Relaciones Grupo-Encuentro
  - Relaciones Grupo-Equipo con estadísticas

## Requisitos

- Go 1.26.2 o superior
- PostgreSQL
- Base de datos `yoplay` con esquema `gestion_encuentro`
- Link https://drive.google.com/drive/folders/1-7PnRYW0MZ0_YdTG1E2mchiMi_XoF-dA?usp=drive_link (secciones script table y script insert)

## Instalación

1. Clona el repositorio:
   ```
   git clone <url-del-repositorio>
   cd gestion_encuentro
   ```

2. Instala las dependencias:
   ```
   go mod tidy
   ```

3. Configura la base de datos PostgreSQL:
   - Crea la base de datos `yoplay`
   - Crea el esquema `gestion_encuentro`
   - Asegúrate de que las tablas estén creadas según los modelos definidos

## Configuración

La configuración de la base de datos está en `config/db.go`. Por defecto:
- Host: localhost
- Puerto: 5432
- Usuario: postgres
- Contraseña: postgres
- Base de datos: yoplay
- Esquema: gestion_encuentro

Modifica estas variables si es necesario.

## Ejecución

Para ejecutar la aplicación:

```
go run main.go
```

El servidor se iniciará en el puerto 8090. Verás el mensaje: "Servidor operando en el puerto 8090"

## Endpoints de la API

### Encuentros

- `GET /encuentro` - Obtener todos los encuentros
- `GET /encuentro/{id}` - Obtener encuentro por ID
- `POST /encuentro` - Crear nuevo encuentro
- `PUT /encuentro/{id}` - Actualizar encuentro
- `DELETE /encuentro/{id}` - Eliminar encuentro

### Grupos

- `GET /grupo` - Obtener todos los grupos
- `GET /grupo/{id}` - Obtener grupo por ID
- `POST /grupo` - Crear nuevo grupo
- `PUT /grupo/{id}` - Actualizar grupo
- `DELETE /grupo/{id}` - Eliminar grupo

### Grupo-Encuentro

- `GET /grupoencuentro` - Obtener todas las relaciones grupo-encuentro
- `GET /grupoencuentro/{id}` - Obtener relación por ID
- `POST /grupoencuentro` - Crear nueva relación
- `PUT /grupoencuentro/{id}` - Actualizar relación
- `DELETE /grupoencuentro/{id}` - Eliminar relación

### Grupo-Equipo

- `GET /grupoequipo` - Obtener todas las relaciones grupo-equipo
- `GET /grupoequipo/{id}` - Obtener relación por ID
- `POST /grupoequipo` - Crear nueva relación
- `PUT /grupoequipo/{id}` - Actualizar relación
- `DELETE /grupoequipo/{id}` - Eliminar relación

## Estructura del Proyecto

```
gestion_encuentro/
├── main.go                 # Punto de entrada de la aplicación
├── go.mod                  # Módulo Go y dependencias
├── config/
│   └── db.go              # Configuración de conexión a PostgreSQL
├── controllers/
│   ├── encuentro_controllers.go
│   ├── grupo_controllers.go
│   ├── grupo_encuentro_controllers.go
│   └── grupo_equipo_controllers.go
├── models/
│   ├── encuentro.go
│   ├── grupo.go
│   ├── grupo_encuentro.go
│   └── grupo_equipo.go
└── routes/
    ├── encuentro_routes.go
    ├── grupo_routes.go
    ├── grupo_encuentro_routes.go
    └── grupo_equipo_routes.go
```

## Modelos de Datos

### Encuentro
- `id_encuentro`: Identificador único
- `fecha`: Fecha del encuentro
- `id_torneo`: ID del torneo
- `id_tipo_distribucion`: Tipo de distribución
- `fase_torneo`: Fase del torneo (Grupos, Octavos, Cuartos, Semifinal, Final)
- `id_equipo1`, `id_equipo2`: IDs de los equipos
- `resultado_equipo1`, `resultado_equipo2`: Resultados (Ganador, Perdedor, Empate)
- `activo`: Estado del encuentro
- `fecha_creacion`, `fecha_modificacion`: Timestamps

### Grupo
- `id_grupo`: Identificador único
- `grupo`: Nombre del grupo
- `activo`: Estado
- `fecha_creacion`, `fecha_modificacion`: Timestamps

### Grupo_Encuentro
- `id_grupo_encuentro`: Identificador único
- `id_grupo`: ID del grupo
- `id_encuentro`: ID del encuentro
- `activo`: Estado
- `fecha_creacion`, `fecha_modificacion`: Timestamps

### Grupo_Equipo
- `id_grupo_equipo`: Identificador único
- `id_grupo`: ID del grupo
- `id_equipo`: ID del equipo
- `partidos_jugados`, `empates`, `victorias`, `derrotas`: Estadísticas
- `activo`: Estado
- `fecha_creacion`, `fecha_modificacion`: Timestamps

## Dependencias

- `github.com/gorilla/mux v1.8.1`: Para enrutamiento HTTP
- `github.com/lib/pq v1.12.3`: Driver PostgreSQL

## Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/nueva-funcionalidad`)
3. Commit tus cambios (`git commit -am 'Agrega nueva funcionalidad'`)
4. Push a la rama (`git push origin feature/nueva-funcionalidad`)
5. Abre un Pull Request

## Licencia

Este proyecto es parte de YoPLAY. Consulta los términos de licencia aplicables.
