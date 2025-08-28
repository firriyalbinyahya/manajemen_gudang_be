# Simple Warehouse Management System

This project is a simple warehouse management system with a full-stack application built using Vue.js for the front-end and Golang for the back-end, powered by a MySQL database. It is designed to help businesses manage their product inventory efficiently.

### Key Features
* **User Authentication**: A secure login and registration system to restrict access to product management endpoints. It uses **JWT** (JSON Web Tokens) for authentication and **Bcrypt** for secure password hashing.
* **Product Management**: Comprehensive CRUD (Create, Read, Update, Delete) functionality for managing products.
* **Flexible API**: The product API supports **pagination**, **filtering** (by product status and low stock), and **searching** (by product name, SKU, or location), providing a powerful way to query the product list.
* **3-Layer Architecture**: The back-end code is organized into a **Controller** -> **Service** -> **Repository** pattern for modularity, reusability, and easy maintenance.
* **Technology Stack**:
    * **Backend**: Golang with Gin and GORM
    * **Authentication**: JWT
    * **Database**: MySQL

---

### Project Setup

To get this project up and running, you'll need **Go** and a **MySQL** database installed on your system.

1.  **Clone the repository**:
    ```bash
    git clone [https://github.com/firriyalbinyahya/manajemen_gudang_be.git](https://github.com/firriyalbinyahya/manajemen_gudang_be.git)
    cd manajemen_gudang_be
    ```

2.  **Configure the database**:
    * Create a new MySQL database.
    * Open the `main.go` file and update the database connection string with your credentials:
        ```go
        dsn := "user:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
        ```

3.  **Install dependencies**:
    ```bash
    go mod tidy
    ```

4.  **Run the application**:
    ```bash
    go run main.go
    ```
    The server will start on `http://localhost:8080`. GORM will automatically create the `users` and `products` tables on the first run.

---

### API Documentation

All API endpoints are prefixed with `/api/v1`.

#### 1. Authentication

| Method | Endpoint | Description | Sample Request Body | Success Response |
| :--- | :--- | :--- | :--- | :--- |
| `POST` | `/api/v1/auth/register` | Creates a new user account. | `{"username": "testuser", "password": "password123"}` | `201 Created` with a success message. |
| `POST` | `/api/v1/auth/login` | Logs in and returns a JWT token. | `{"username": "testuser", "password": "password123"}` | `200 OK` with a token. |

#### 2. Products

> **Note**: All product endpoints require an `Authorization: Bearer <TOKEN>` header.

| Method | Endpoint | Description | Query Parameters / Request Body | Success Response |
| :--- | :--- | :--- | :--- | :--- |
| `POST` | `/api/v1/products` | Adds a new product to the inventory. | `{"productName": "Laptop XYZ", "sku": "LPT-XYZ-01", "quantity": 100, "location": "A-101", "status": "In Stock"}` | `201 Created` with a success message. |
| `GET` | `/api/v1/products` | Retrieves a list of products. | **Optional:** `?page=1`, `?per_page=10`, `?status=In%20Stock`, `?low_stock=true`, `?search=ASUS` | `200 OK` with a JSON array of products and pagination metadata. |
| `GET` | `/api/v1/products/:id` | Retrieves a single product by ID. | N/A | `200 OK` with the product object. |
| `PUT` | `/api/v1/products/:id` | Updates an existing product. | **Optional:** `{"quantity": 90, "status": "Low Stock"}` | `200 OK` with a success message. |
| `DELETE`| `/api/v1/products/:id` | Deletes a product by ID. | N/A | `200 OK` with a success message. |

### Screenshots or Demo Link
