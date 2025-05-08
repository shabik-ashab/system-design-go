# 🌀 Round-Robin Reverse Proxy in Go

[![Go](https://img.shields.io/badge/Built%20with-Go-blue?logo=go)](https://golang.org)
[![Status](https://img.shields.io/badge/status-Learning%20Project-yellow)]()

A beginner-friendly, minimal reverse proxy built in Go using round-robin load balancing.  
Great for learning **system design**, **HTTP routing**, and **Go concurrency**.

---

## 📖 Table of Contents

- [🌐 What is a Reverse Proxy?](#-what-is-a-reverse-proxy)
- [🔁 Why Round-Robin?](#-why-round-robin)
- [🏗 Architecture](#-architecture)
- [🧱 Project Structure](#-project-structure)
- [📘 Learning Notes](#-learning-notes)

---

## 🌐 What is a Reverse Proxy?

A **reverse proxy** is a server that sits between clients and backend services.  
Clients send requests to the proxy, which forwards them to appropriate backend servers.

✅ Benefits:
- Load balancing  
- Centralized logging  
- SSL termination  
- Rate limiting / throttling  
- Health checks & failover

---

## 🔁 Why Round-Robin?

**Round-robin** is a simple load balancing strategy. Requests are sent to servers in a fixed order.

Example:

Request 1 → Backend A
Request 2 → Backend B
Request 3 → Backend C
Request 4 → Backend A ...


✅ Why use it?
- Simple to implement  
- Evenly distributes traffic  
- No need for load tracking (in basic form)

---

## 🏗 Architecture

```text
               +-------------------+
    Client --> | Reverse Proxy     | --> Backend 1 (localhost:8081)
              ( Port 8000 )        | --> Backend 2 (localhost:8082)
               +-------------------+ --> Backend 3 (localhost:8083)
The proxy server:

Accepts requests on :8000

Uses round-robin to forward to one of three backend servers

Returns the backend's response to the client

🧱 Project Structure

round-robin-reverse-proxy/
├── main.go              # Entry point for the proxy server
├── proxy/
│   └── proxy.go         # Reverse proxy logic (round-robin)
├── backend/
│   ├── server1.go       # Dummy backend server on port 8081
│   ├── server2.go       # Dummy backend server on port 8082
│   └── server3.go       # Dummy backend server on port 8083
├── go.mod               # Go module file
└── README.md            # This file


📘 Learning Notes
This repo is intentionally simple so you can learn:

What a reverse proxy is and how it works

How to implement basic round-robin load balancing

Go's http.Handler, request cloning, and concurrency