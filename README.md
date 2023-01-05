## How to Run

- Setup Docker
  ```bash
  docker compose up
  ```
- Execute HTTP Binary
  ```bash
  go run ./cmd/grpc
  ```

## System Architecture

We use a clean architecture according to standards so that the design of our code is neater and more structured.
There are 3 main layers, namely controller, usecase/service, repository.
![cleanarc.png](documentation%2Fcleanarc.png)
We also use redis on the list-product and store it on the proxy repository layer, the goal is not to interfere with the business logic at the usecase/service layer.

![proxy-pattern.png](documentation%2Fproxy-pattern.png)

for more details about the proxy pattern can be seen here :
https://refactoring.guru/design-patterns/proxy

## API Documentation

- Add Product Endpoint : POST localhost:1323/add-product
- Sample Request
```json
{
  "id" : "005",
  "name": "xiomay",
  "price" : 900000,
  "description": "contoh des",
  "quantity": 100,
}
  ```

- Sample Response
```json
{
  "id" : "005",
  "name": "xiomay",
  "price" : 900000,
  "description": "contoh des",
  "quantity": 100
}
  ```
- List Product Endpoint : POST localhost:1323/list-product
- Sample Add Product Request
```json
{
  "sort_by" : "ProductNameDESC"
}
  ```

sort_by have 5 options

- NewProduct
- CheapestPrice
- ExpensivePrice
- ProductNameASC
- ProductNameDESC

- Sample Response
```json
[
  {
    "id": "005",
    "name": "xiomay",
    "price": 900000,
    "description": "contoh des",
    "quantity": 100,
    "created_at": "2023-01-05T04:02:20.038481Z",
    "updated_at": "2023-01-05T04:02:20.038481Z"
  },
  {
    "id": "002",
    "name": "samsung",
    "price": 2000000,
    "description": "contoh des",
    "quantity": 10,
    "created_at": "2023-01-05T03:59:55.064252Z",
    "updated_at": "2023-01-05T03:59:55.064252Z"
  },
  {
    "id": "003",
    "name": "oppo",
    "price": 1400000,
    "description": "contoh des",
    "quantity": 25,
    "created_at": "2023-01-05T04:00:15.330214Z",
    "updated_at": "2023-01-05T04:00:15.330214Z"
  },
  {
    "id": "001",
    "name": "iphone",
    "price": 900000,
    "description": "contoh des",
    "quantity": 16,
    "created_at": "2023-01-05T03:59:40.20027Z",
    "updated_at": "2023-01-05T03:59:40.20027Z"
  },
  {
    "id": "004",
    "name": "advan",
    "price": 1700000,
    "description": "contoh des",
    "quantity": 52,
    "created_at": "2023-01-05T04:01:13.111959Z",
    "updated_at": "2023-01-05T04:01:13.111959Z"
  }
]
  ```

## POC
- Add Product
![add-product.png](documentation%2Fadd-product.png)
- List Product
![list-product.png](documentation%2Flist-product.png)