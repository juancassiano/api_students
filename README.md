# api_students
Api in Golang

Routes:
- GET /students - List all students
- GET /students?active=True/False - List all students Active/Inactive
- POST /students - Create students
- GET /students/:id - Find students by Id
- PUT /students/:id - Update student
- DELETE /students/:id - Delete student

Struct Student:
- Name
- CPF
- Email
- Age
- Active