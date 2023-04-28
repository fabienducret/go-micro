# Go Micro

This project is an entry point to the world of Microservices with Go.

It's a result of an udemy's training and my point of view.

---

## Architecture

![Alt text](docs/architecture_diagram.png 'Architecture Diagram')

The entry point is the API **broker-service**.
This endpoint is a HTTP API that can receive 3 kinds of actions :

- auth
- log
- mail

When the **broker-service** receives a payload, it communicates with 3 other services depending on the payload :

- **authentication-service**
- **logger-service**
- **mail-service**

The communication between the broker and the 3 other services is through **RCP over TCP**.

---

## How to use

### Build the containers

```bash
cd project
make up_build
```

You must have your containers running:

![Alt text](docs/docker_containers.png 'Docker containers')

### Broker calls

#### Hit request

Request

```bash
curl -X POST http://localhost:8080 \
-H "Content-Type: application/json" \
```

Response

```json
{
  "error": false,
  "message": "Hit the broker"
}
```

---

#### Authentication request

Request

```bash
curl -X POST http://localhost:8080/handle \
-H "Content-Type: application/json" \
-d '{"action":"auth","auth":{"email":"admin@example.com","password":"verysecret"}}'
```

Response with valid credentials

```json
{
  "error": false,
  "message": "Authenticated !",
  "data": {
    "email": "admin@example.com",
    "firstname": "Admin",
    "lastname": "User"
  }
}
```

Response with invalid credentials

```json
{
  "error": true,
  "message": "invalid credentials"
}
```

---

#### Log request

Request

```bash
curl -X POST http://localhost:8080/handle \
-H "Content-Type: application/json" \
-d '{"action":"log","log":{"name":"event","data":"Hello world !"}}'
```

Response

```json
{
  "error": false,
  "message": "Log handled for:event"
}
```

---

#### Mail request

Request

```bash
curl -X POST http://localhost:8080/handle \
-H "Content-Type: application/json" \
-d '{"action":"mail","mail":{"from": "me@example.com", "to": "you@example.com", "subject": "Test email", "message": "Hello world"}}'
```

Response

```json
{
  "error": false,
  "message": "Message sent to you@example.com"
}
```
