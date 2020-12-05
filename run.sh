#!/bin/bash
cd message-app && docker-compose up -d && cd ../
cd user-app && docker-compose up -d && cd ../
cd product-app && docker-compose up -d && cd ../
cd cart-app && docker-compose up -d && cd ../
cd order-app && docker-compose up -d && cd ../
cd front-app && docker-compose up -d && cd ../