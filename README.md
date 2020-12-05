# EC micro-service with Go and

This repository is implemented electronic commerce application with micro service architecture

## Technologies

- backend
  - Go(echo, gorm)
  - mysql
  - Docker
  - RabbitMQ
- frontend
  - TypeScript
  - React
  - nginx

## Usage

Docker and docker-compose are required

```bash
./run.sh  # runnning application servers
./down.sh # down applications
```

## Directory Structure

```bash
.
├── cart-app    # Shopping Cart
├── front-app   # Frontend Application
├── message-app # Message Q for Connecting Applications
├── order-app   # Order Histories
├── product-app # Product List
└── user-app    # Management User and Authentication
```
