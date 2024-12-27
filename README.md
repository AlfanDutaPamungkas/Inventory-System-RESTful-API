# INVENTORY SYSTEM RESTful API
The Inventory System is a RESTful API built using Golang, designed to manage product inventory. The API is documented using OpenAPI 3.0 for easy reference and can be explored interactively.

## Features
- Authentication and authorization for secure API access (super admin and admin)
- CRUD products to the database
- Super admin can see the log activity

## Prerequisites
- [Golang](https://golang.org/doc/install) v1.18 or higher
- [MySQL](https://dev.mysql.com/downloads/mysql/) or any other SQL-based database
- [Cloudinary](https://cloudinary.com/) account to store product images in the cloud
- [Google Wire](https://github.com/google/wire) for dependency injection
- [Golang Migrate](https://github.com/golang-migrate/migrate) for database migration

## Instalation
1. Clone the repository:
    ```bash
    git clone https://github.com/AlfanDutaPamungkas/Inventory-System-RESTful-API.git
    ```
2. Navigate to the project directory:
    ```bash
    cd Inventory-System-RESTful-API
    ```
3. Install dependencies:
    ```bash
    go mod download
    ```
4. Set up your environment variables:
    Create a `.env` file in the project root and specify the following variables:
    ```env
    JWT_TOKEN_SECRET=your_jwt_secret_key
    JWT_EXPIRED_TIME_TOKEN=your_jwt_expired_time
    CLOUD_NAME=your_cloduinary_cloud_name
    CLOUDINARY_API_KEY=your_cloudinary_api_key
    CLOUDINARY_API_SECRET=your_cloudinary_api_secret
    DB_URL=your_db_url
    ```
5. Start the server:
    ```bash
    go run .
    ```
    The API will be running at `http://localhost:3000`.

## API Documentation (OpenAPI 3.0)

The API is fully documented using the OpenAPI 3.0 specification. You can view the  `apispec.json`

Contributing
Feel free to open issues or submit pull requests if you want to contribute to this project.
