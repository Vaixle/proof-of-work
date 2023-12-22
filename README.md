# proof-of-work
[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/inactive.svg)](https://www.repostatus.org/#inactive)
![GitHub License](https://img.shields.io/github/license/vaixle/proof-of-work)

![image](https://github.com/Vaixle/proof-of-work/assets/58184233/5e25f27c-a2a3-4a64-ac49-3318aa2bd84b)

Proof of work (PoW) is a form of cryptographic proof in which one party (the prover) proves to others (the verifiers) that a certain amount of a specific computational effort has been expended.[1] Verifiers can subsequently confirm this expenditure with minimal effort on their part. The concept was invented by Moni Naor and Cynthia Dwork in 1993 as a way to deter denial-of-service attacks and other service abuses such as spam on a network by requiring some work from a service requester, usually meaning processing time by a computer. The term "proof of work" was first coined and formalized in a 1999 paper by Markus Jakobsson and Ari Juels.

## Conntents:

- [Description](#Description)
- [Main Technologies](#Main-technologies)
- [Getting started](#Getting-started)
    - [Start with Docker](#Start-with-docker)



### Description
- A protected TCP server with the challenge-response.
- After Proof Of Work verification, server send one of the quotes from “word of wisdom” book.


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
docker build -t pow-server -f ./docker/Dockerfile.server .
```
```
docker run --network pow-network --name pow-server -p 8088:8088 -it pow-server
```
```
docker build -t pow-client -f ./docker/Dockerfile.client . 
```
```
docker run --network pow-network -it pow-client
```
