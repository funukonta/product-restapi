# Product REST API
Project Test Erajaya - Product Rest API

## External library used :
- mux
- pq

## Architecture:
Saya mennggunakan pendekatan Domain-Driven Design (DDD) dalam mengembangkan Project ini.

## How To Set Up:
 1. Clone the Repo `git clone https://github.com/funukonta/product-restapi`
 2. Create Docker Network `docker network create network_product`
 3. Run Docker compose up `docker compose up -d`
 6. Server ready to serve

## End Points
1. `/product`
method yang digunakan `POST` dengan format JSON Body, contoh :
`
{
    "name":"monitor",
    "price":100000,
    "desc":"monitor gaming",
    "qty":15
}
`

2. `/product/sort/{sortby}/{type}`
method yang digunakan `GET` dengan query url, contoh : `/product/sort/price/asc`
- sortby query : (price,name,createdat)
- type query : (asc,desc)