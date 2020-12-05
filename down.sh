#!/bin/bash
cd front-app && docker-compose down && cd ../
cd order-app && docker-compose down && cd ../
cd cart-app && docker-compose down && cd ../
cd product-app && docker-compose down && cd ../
cd user-app && docker-compose down && cd ../
cd message-app && docker-compose down && cd ../