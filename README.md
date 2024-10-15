# **AI-Recruitment-backend**

## **Project Overview**

**AI-Recruitment-backend** is the backend system for an AI-powered recruitment platform. It allows companies to manage job listings, resumes, and streamline the hiring process with AI-driven candidate analysis and job matching.

This backend is built with **Go** and the **Gin** web framework, with **MySQL** as the database, and uses **JWT** for authentication. It integrates with external AI services to analyze resumes and recommend job candidates efficiently.

## **Key Features**

- **User Authentication**: JWT-based login and registration.
- **Job Management**: Create, update, delete, and search for job postings.
- **Resume Management**: Upload, analyze, and manage resumes.
- **AI Integration**: AI-powered resume screening and job matching.
- **Admin Features**: Manage users, job posts, and resumes.
- **RESTful API**: For smooth front-end and back-end interaction.

## **Project Structure**

```
├── controllers/        # Request handlers (job posts, users, resumes)
├── models/             # Database schema and ORM
├── routes/             # API route definitions
├── services/           # Core logic and AI service integration
├── middleware/         # Authentication middleware (JWT)
├── config/             # Configuration settings (database, AI, JWT)
├── utils/              # Utility functions
├── tests/              # Unit and integration tests
└── main.go             # Application entry point
```

## **Tech Stack**

- **Language**: Go
- **Web Framework**: Gin
- **Database**: MySQL
- **Authentication**: JWT
- **AI Services**: Integrated with external Python-based AI services

## **Requirements**

- **Go** 1.16+
- **MySQL** 5.7+
- **Python** 3.8+ (for AI services)
- **Docker** (optional, for deployment)

## **Installation Guide**

### **1. Clone the repository**

```bash
git clone https://github.com/your-repo/AI-Recruitment-backend.git
cd AI-Recruitment-backend
```

### **2. Install Go dependencies**

Ensure you have Go installed, then run:

```bash
go mod tidy
```

### **3. Configure Environment Variables**

Create a `.env` file in the root directory and add the following variables:

```bash
PORT=8080
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_mysql_password
DB_NAME=ai_recruitment
JWT_SECRET=your_jwt_secret_key
AI_SERVICE_URL=http://localhost:5000/analyze
```

### **4. Set up the MySQL Database**

Ensure MySQL is running and create the database:

```sql
CREATE DATABASE ai_recruitment;
```

### **5. Run Database Migrations**

Run any required migrations to set up the database schema.

### **6. Start the Application**

Launch the backend locally:

```bash
go run main.go
```

The server will run at `http://localhost:8080`.

### **7. Start the AI Service**

If you are using a separate AI service, navigate to the AI service directory and start it:

```bash
python ai_service.py
```

The AI service will run on `http://localhost:5000`.

## **Running Tests**

To run the unit tests for the backend, use:

```bash
go test ./...
```

## **Deployment**

### **Docker Deployment**

To deploy the backend using Docker:

1. **Build the Docker image:**

   ```bash
   docker build -t ai-recruitment-backend .
   ```

2. **Run the Docker container:**

   ```bash
   docker run -p 8080:8080 ai-recruitment-backend
   ```

### **Kubernetes Deployment** (Optional)

For large-scale deployment, you can use Kubernetes. Example deployment files can be included in the `k8s/` folder.

## **Contributing**

We welcome contributions! If you'd like to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-xyz`).
3. Make your changes and commit them (`git commit -m 'Add new feature'`).
4. Push your changes to your fork (`git push origin feature-xyz`).
5. Open a pull request to the main repository.

## **License**

This project is licensed under the **MIT License**. For more details, see the [LICENSE](LICENSE) file.
