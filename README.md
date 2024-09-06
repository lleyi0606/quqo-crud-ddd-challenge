# Products CRUD

### Software Design
**DDD (Domain-Driven Design)** 

By maintianing an onion structure with domain-application-infrastructure. 

### Data Source
CockroachDB
Redis 

### APIs
* GET `/products`: Retrieve all products.
* GET `/products/:id`: Retrieve a specific product.
* POST `/products`: Create a new product.
* PUT `/products/:id`: Update an existing product.
* DELETE `/products/:id`: Delete an existing product.
* GET `/products/search?name=**` Return a list of products matching the text string. (With AWS Opensearch)

![Screenshot 2024-09-06 at 4 07 34 PM](https://github.com/user-attachments/assets/fe2108b5-db6b-4648-a8b4-d02cb90ee577)
