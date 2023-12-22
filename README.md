# proof-of-work

![Proof of work](https://upload.wikimedia.org/wikipedia/commons/thumb/5/55/Proof_of_Work_challenge_response.svg/825px-Proof_of_Work_challenge_response.svg.png)

## Conntents:

- [Description](#Description)
- [Main Technologies](#Main-technologies)
- [Getting started](#Getting-started)
    - [Start with Docker](#Start-with-docker)



### Description
- A protected TCP server with the challenge-response.
- The choice of the POW algorithm should be explained.
-  After Proof Of Work verification, server send one of the quotes from “word of wisdom” book.


---

### Main Technologies

| **Backend**  |                                                              ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)                                                              |
|:------------:|:--------------------------------------------------------------------------------------------------------------------------:|
| **PaaS**  |        ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)        |

---

### Getting started
Run application.

#### Start with Docker

> Clone the repository

```
git clone https://github.com/Vaixle/proof-of-work
```

> Run application in Docker container
```
docker network create pow-network 
```
```
docker build -t pow-server -f ./docker/Dockerfile.server
```
```
docker run --network pow-network --name pow-server -p 8088:8088 -it pow-server
```
```
docker build -t pow-client -f ./docker/Dockerfile.client 
```
```
docker run --network pow-network -it pow-client
```