# Go Micro

This project is an entry point to the world of Microservices.

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

- authentication-service
- logger-service
- mail-service

The communication between the broker and the 3 other services is through **RCP over TCP**.

---

## How to use

```bash
cd project
make up_build
```

You must have your containers running:

![Alt text](docs/docker_containers.png 'Docker containers')

Then you can send payloads to the broker-service.

```bash
curl -X POST http://localhost:8080/handle \
-H "Content-Type: application/json" \
-d '{"action":"auth","auth":{"email":"admin@example.com","password":"verysecret"}}'
```

![Alt text](docs/auth_request.png 'Authentication request')

---

```bash
curl -X POST http://localhost:8080/handle \
-H "Content-Type: application/json" \
-d '{"action":"log","log":{"name":"event","data":"Hello world !"}}'
```

![Alt text](docs/log_request.png 'Log request')

---

```bash
curl -X POST http://localhost:8080/handle \
-H "Content-Type: application/json" \
-d '{"action":"mail","mail":{"from": "me@example.com", "to": "you@example.com", "subject": "Test email", "message": "Hello world"}}'
```

![Alt text](docs/mail_request.png 'Mail request')
