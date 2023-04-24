# Go Micro

This project is an entry point to the world of Microservices.

It's a result of an udemy's training and my point of view.

## Architecture

![Alt text](docs/architecture_diagram.png 'Architecture Diagram')

The entry point is the API **broker-service**.

## How to use

```bash
curl -X POST http://localhost:8080/handle \
-H "Content-Type: application/json" \
-d '{"action":"auth","auth":{"email":"admin@example.com","password":"verysecret"}}'
```

```bash
curl -X POST http://localhost:8080/handle \
-H "Content-Type: application/json" \
-d '{"action":"log","log":{"name":"event","data":"Hello world !"}}'
```

```bash
curl -X POST http://localhost:8080/handle \
-H "Content-Type: application/json" \
-d '{"action":"mail","mail":{"from": "me@example.com", "to": "you@example.com", "subject": "Test email", "message": "Hello world"}}'
```
