# **AI-Recruitment-backend**

[简体中文](./README_zhCN.md)

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
├── AI-service            # AI-related logic and functions
├── Dockerfile            # Docker setup
├── LICENSE               # License file
├── README.md             # Project overview and instructions
├── bin                   # Compiled binaries
├── cmd                   # App entry points
├── configs               # Config files
├── data                  # Data storage
├── docker-compose.yml    # Docker Compose config for multi-service setup
├── docs                  # Documentation (API docs, design notes)
├── internal              # Internal code, project-only
├── pkg                   # Reusable packages
└── scripts               # Scripts
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

### **3. Set up the MySQL Database**

Ensure MySQL is running and create the database:

```sql
CREATE DATABASE ai_recruitment;
```

### **4. Run Database Migrations**

Run any required migrations to set up the database schema.

### **5. Start the Application**

Launch the backend locally:

```bash
go run main.go
```

The server will run at `http://localhost:8001`.

### **7. Start the AI Service**

If you are using a separate AI service, navigate to the AI service directory and start it:

```bash
python AI-service/main.py
```

The AI service will run on `http://localhost:5000`.

## **Deployment**

### **Docker Deployment**

To deploy the backend using Docker:

1. **Build the Docker image:**

   ```bash
   docker build -t ai-recruitment-backend .
   ```

2. **Run the Docker container:**

   ```bash
   docker run -p 8001:8000 ai-recruitment-backend
   ```

### **Docker Compose Deployment**

```bash
docker compose up -d --build
```

## **Contributing**

We welcome contributions! If you'd like to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-xyz`).
3. Make your changes and commit them (`git commit -m 'Add new feature'`).
4. Push your changes to your fork (`git push origin feature-xyz`).
5. Open a pull request to the main repository.

## **License**

This project is licensed under the **MIT License**. For more details, see the [LICENSE](LICENSE) file.
