# CRUD App Documentation

## Introduction

This document provides detailed information about the CRUD app built using Golang's Gin framework for managing "Person" resources.

## Endpoints

### 1. Create a Person

**Endpoint:** `POST /api`

**Request:**

```json
{
  "name": "John Doe",
  "location": 30,
  "title": "Software Engineer"
}
```

**Sample success response:**

```json
{
  "status": 201,
  "message": "success",
  "data": {
    "data": {
      "InsertedID": "65015a404eb4081132a94aab"
    }
  }
}
```

**Sample error response:**

```json
{
  "status": 400,
  "message": "error",
  "data": {
    "error": "Key: 'Person.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'Person.Location' Error:Field validation for 'Location' failed on the 'required' tag\nKey: 'Person.Title' Error:Field validation for 'Title' failed on the 'required' tag"
  }
}
```

### 2. Fetch all persons

**Endpoint:** `GET /api`

**Request:**

```json
{}
```

**Sample success response:**

```json
{
  "status": 200,
  "message": "success",
  "data": {
    "data": [
      {
        "id": "65009c65e2d7c9b8d5aecb32",
        "name": "Johnson Awah-Alfred",
        "location": "Enugu",
        "title": "Chief Blockchain Gopher"
      },
      {
        "id": "65009c8ce2d7c9b8d5aecb34",
        "name": "Chidera Anichebe",
        "location": "Lagos, Nigeria",
        "title": "Badass Software Engineer"
      }
    ]
  }
}
```

**Sample error response:**

```json
{
  "status": 500,
  "message": "error",
  "data": {
    "error": "Error fetching persons: mongo: no documents in result"
  }
}
```

### 3. Update a person

**Endpoint:** `PUT /api/:userId`

**Request:**

```json
{
  "name": "Johnson Awah-Alfred",
  "location": "Enugu",
  "title": "Chief Blockchain Gopher"
}
```

**Sample success response:**

```json
{
  "status": 200,
  "message": "success",
  "data": {
    "data": {
      "id": "65009c65e2d7c9b8d5aecb32",
      "name": "Johnson Awah-Alfred",
      "location": "Enugu",
      "title": "Chief Blockchain Gopher"
    }
  }
}
```

**Sample error response:**

```json
{
  "status": 500,
  "message": "error",
  "data": {
    "data": "User with specified ID not found!"
  }
}
```

### 4. Delete a person

**Endpoint:** `DELETE /api/:userId`

**Request:**

```json
{}
```

**Sample success response:**

```json
{
  "status": 200,
  "message": "success",
  "data": {
    "data": "User successfully deleted!"
  }
}
```

**Sample error response:**

```json
{
  "status": 404,
  "message": "error",
  "data": {
    "data": "User with specified ID not found!"
  }
}
```

### 5. Fetch single person

**Endpoint:** `GET /api/:userId`

**Request:**

```json
{}
```

**Sample success response:**

```json
{
  "status": 200,
  "message": "success",
  "data": {
    "data": {
      "id": "65009c65e2d7c9b8d5aecb32",
      "name": "Johnson Awah-Alfred",
      "location": "Enugu",
      "title": "Chief Blockchain Gopher"
    }
  }
}
```

**Sample error response:**

```json
{
  "status": 500,
  "message": "error",
  "data": {
    "data": "mongo: no documents in result"
  }
}
```

## Known Limitations/Assumptions

This app assumes that two different individuals can have the same name. Also, only basic input validation is implemented. More robust validation may be required in a production environment.

## Setup

### 1. Clone the repository

```bash
git clone https://github.com/starlingvibes/HNGX-BE-Task_2.git
```

### 2. Install dependencies

```bash
go get
```

### 3. Run the app

```bash
go build main.go && ./main
```
