# Articles Management System Backend (AMS Backend)

🚀 A secure and scalable backend server for managing articles, user authentication, and authorization, built with Golang and PostgreSQL.

## Features
- User Authentication & Authorization (JWT-based)
- Secure Password Encryption (bcrypt)
- CRUD Operations for Articles
- PostgreSQL Database Integration
- RESTful API Endpoints
- Role-Based Access Control (RBAC)
- Input Validation & Security Best Practices

## Tech Stack
- **Backend: ** Golang (Gin) Framework
- **Database: ** PostgreSQL Database
- **Authentication & Session Managment : ** Json Web Token
- **Encryption: ** bcrypt
- **PGX_v5: ** For PostgreSQL connection

## Getting Started

### Clone the Repository
```bash
git clone https://github.com/omkero/House-of-Wisdom.git

cd House-of-Wisdom
```

### Install Dependencies
verifies the dependencies in your go.mod and go.sum files
```bash
go mod tidy
```

### Configure Environment Variables
make sure you create a `.env` file if it not exist and add the following:
```
DATABASE_URL=postgres://user:password@localhost:5432/articles_db?sslmode=disable
JWT_SECRET=your_secret_key
```

### Run the Server
```bash
go run main.go
```
Server will start at **http://localhost:8080** 🚀

## API Endpoints

| Method | Endpoint                  | Description           | Auth Required |
|--------|---------------------------|----------------------|--------------|
| GET    | `/api/v1/get_all_articles` | Get all articles     | ❌ No  |
| GET    | `/api/v1/get_article`      | Get an article by title | ❌ No  |
| POST   | `/api/v1/create_new_article` | Create a new article | ✅ Yes |
| DELETE | `/api/v1/delete_article`   | Delete an article   | ✅ Yes |
| PUT    | `/api/v1/edit_article`     | Edit an article     | ✅ Yes |
| POST   | `/api/v1/auth/login`       | User login          | ❌ No  |
| POST   | `/api/v1/auth/signup`      | User signup         | ❌ No  |

## License
This project is licensed under the **MIT License**.

---

💡 **Contributions are welcome!** Feel free to open an issue or submit a pull request. 🚀

