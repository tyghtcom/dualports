# Dual Server HTTP Application

A Go HTTP application running two independent servers on separate ports, each with its own security profile.

## Architecture

**Public Server (Port 8080)**
- Exposed to external traffic
- Restricted single endpoint: `GET /` only
- All other requests return `403 Forbidden`

**Internal Server (Port 8081)**
- Internal application traffic only — good for an internal api
- Multiple mux routing endpoint support
- No access restrictions 

## Endpoints

### Public Server (:8080)
```
GET /    - Public endpoint response
*        - 403 Forbidden (all other paths/methods)
```

### Internal Server (:8081)
```
POST   /create     - Create operations
GET    /read/      - Read operations
PUT    /update/    - Update operations
DELETE /delete/    - Delete operations
```

## Features

- **Graceful Shutdown**: 30-second timeout for connection draining
- **Signal Handling**: Responds to SIGINT and SIGTERM
- **Concurrent Servers**: Both servers run simultaneously
- **Operational Logging**: Startup and shutdown events logged

## Requirements
- Go 1.24+
- Webserver (good: nginx, better: angie) using a gateway to serve internal port 8081,
or, point the public endpoint to resolve your domain, and add TLS support.
- Create your desired handlers for the endpoints. 
- Your imagination.