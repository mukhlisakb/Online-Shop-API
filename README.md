# Online Shop for small scnale online shop or UMKM

Welcome to the Online-Shop-API repository! This API is designed for small-scale online shops, providing a comprehensive yet straightforward solution for managing products, orders, and customer information.

# Features

- Product Management: Create, update, delete, and retrieve product information.
- Order Management: Handle orders, including creation, updating, and retrieval of order details.
- Payment Processing: Integrate with payment gateways to handle transactions smoothly.
- Security: Implement robust security measures, including password hashing and secure data storage.

# Endpoints
- GET /api/v1/products - Retrieve a list of products.
- GET /api/v1/products/:id - Retrieve a specific product by ID
- GET /api/v1/products/id - 
- POST /api/v1/checkout - Process the checkout for a specific product
- POST /api/v1/order/:id/confirm - Confirm an order by ID
- POST /admin/products - Add a new product (admin only).
- PUT /admin/products/:id - Update a specific product by ID (admin only)
- DELETE  /admin/products/:id - Delete a specific product by ID (admin only)


# Technologies Used
- Go: The main programming language.
- Gin: A web framework for building APIs in Go.
- Bcrypt: For hashing passwords.
- SQLite/PostgreSQL/MySQL: Choose your preferred database.

# Getting Started

Follow these steps to set up and run the Online-Shop-API.

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/online-shop-api.git
   cd online-shop-api
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Run Docker for PostgreSQL Database:
   ```
   docker run --name onlineshopdb 
   -e POSTGRES_USER=admin 
   -e POSTGRES_PASSWORD=admin 
   -e POSTGRES_DB=database 
   -d -p 5432:5432 postgres:16
   ```

4. Export Required Environment Variables:
   ```
   export DB_URI=postgres://admin:admin@localhost:5432/database?sslmode=disable
   export ADMIN_SECRET=secret
   ```

5. Run the Program:
   ```
   go run .
   ```

# Contributing
We welcome contributions! Please fork the repository and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

# License
This project is licensed under the MIT License - see the LICENSE file for details.