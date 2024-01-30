# FILE-PROCESSOR-UPLOAD

---
- Tabla de contenidos 
  - [Descripción](#descripción)
  - [Requerimientos de la aplicación](#requerimientos-de-la-aplicación)
  - [Inicializar la aplicación localmente](#inicializar-la-aplicación-localmente)
    - [Paquetes y Herramientas Instalados](#paquetes-y-herramientas-instalados)
---

## Descripción
API construida con el objetivo de recibir archivos, leer su contenido y procesar la información del documento.

---

## Requerimientos de la aplicación
1. Instalar y configurar git
- [Descargar Git](https://git-scm.com/downloads)
- [Pasos para la configuración](https://git-scm.com/book/en/v2/Getting-Started-First-Time-Git-Setup)

2. Instalar Make
- Opción 1: Descargar directamente de [GNUwin32](https://gnuwin32.sourceforge.net/packages/)
- Opción 2: Instalar [Chocolatey](https://chocolatey.org/install) y ejecutar el comando `choco install make`

3. Instalación de Go (Versión utilizada: 1.21.5)
- [Instrucciones de instalación de Go](https://go.dev/doc/install)
- Verificar la instalación con el comando `go version`  

---

## Inicializar la aplicación localmente
1. Clonar el repositorio
    ```bash
    git clone https://github.com/Pifiaa/File-Processor.git
    ```

3. Ejecutar el siguiente comando para descargar todas las depedencias
    ```bash
    go mod download
    ```

---

## Puesta en marcha

Existen varias formas de ejecutar el programa:

#### Go build

Go build es una herramienta de línea de comandos utilizada para compilar programas Go. Compila un archivo de código fuente Go y crea un archivo binario ejecutable.

La herramienta Go build determina automáticamente las dependencias del paquete Go y también las compila. La herramienta de compilación Go puede utilizarse para compilar un único paquete o un proyecto completo.

Se puede ejecutar usando:

```
go build main.go
```

También se pueden usar los comandos del makefile como:

```
make build: construirá un ejecutable para cada sistema operativo

make build-darwin: construye un ejecutable para Mac

make build-linux: construye un ejecutable para Linux

make build-windows: construye un ejecutable para Windows
```

#### Go run

La herramienta de ejecución Go es una utilidad integrada en la línea de comandos que permite compilar y ejecutar código Go de forma rápida y sencilla, sin la necesidad de crear un ejecutable.

Se puede ejecutar usando:

```bash
go run main.go
```

También se puede utilizar el siguiente comando del makefile:

```bash
make run
```

---

### Paquetes y Herramientas Instalados    
1. [viper](https://github.com/spf13/viper): Librería de configuración flexible que simplifica el manejo de la configuración de la aplicación con soporte para diversos formatos y fuentes.

2. [goose](https://github.com/pressly/goose): Herramienta que facilita la gestión de cambios estructurales en bases de datos, asegurando una evolución ordenada.
   
3. [gorm](https://gorm.io/): ORM (Object-Relational Mapping) para Go que simplifica la interacción con bases de datos relacionales a través de operaciones orientadas a objetos.
   
   GORM necesita controladores de bases de datos para conectarse a determinadas bases de datos. Dependiendo de la base de datos que vaya a utilizar, instale el controlador correspondiente. GORM es compatible con las bases de datos MySQL, PostgreSQL, SQLite, SQL Server y TiDB.

4. [Fiber](https://docs.gofiber.io): Framework web para Go que facilita la construcción de aplicaciones web y APIs de manera eficiente.

5. [jwt-go](https://github.com/golang-jwt/jwt):  La biblioteca jwt-go en Go facilita la creación, firma y verificación de JWT. 

---

## Endpoints
| Descripción                         | Metodo HTTP   | Ruta                       | Estatus  |
|-------------------------------------|---------------|----------------------------|----------|
| Subida de archivo de operaciones    | Post          | /api/upload                | &#x2705; | 
| Autenticación de usuario            | Post          | /api/login                 | &#x2705; | 