# Waitlist API

A simple Go backend API for managing a waitlist. This API provides endpoints to add users to a waitlist and retrieve the list of waitlist entries.

## Features

- Add users to the waitlist
- Retrieve the list of waitlist entries
- SQLite database for data persistence
- Input validation
- Error handling

## Prerequisites

- Go 1.22 or later
- SQLite3

## Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```

## Running the API

Start the server:
```bash
go run main.go
```

The server will start on port 8080.

## API Endpoints

### Add to Waitlist
- **URL**: `/api/waitlist`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "email": "user@example.com",
    "name": "John Doe"
  }
  ```
- **Success Response**: 
  - Code: 201 Created
  - Body: `{"message": "Successfully added to waitlist"}`

### Get Waitlist
- **URL**: `/api/waitlist/list`
- **Method**: `GET`
- **Success Response**: 
  - Code: 200 OK
  - Body: Array of waitlist entries
    ```json
    [
      {
        "id": 1,
        "email": "user@example.com",
        "name": "John Doe",
        "created_at": "2024-04-10T12:00:00Z"
      }
    ]
    ```

## Error Responses

- **400 Bad Request**: Invalid request body or missing required fields
- **409 Conflict**: Email already exists in the waitlist
- **500 Internal Server Error**: Database errors

## Frontend Integration

To integrate this API with your frontend:

1. **Adding to Waitlist**:
   ```javascript
   async function addToWaitlist(email, name) {
     const response = await fetch('http://localhost:8080/api/waitlist', {
       method: 'POST',
       headers: {
         'Content-Type': 'application/json',
       },
       body: JSON.stringify({ email, name }),
     });
     
     if (!response.ok) {
       throw new Error('Failed to add to waitlist');
     }
     
     return await response.json();
   }
   ```

2. **Getting Waitlist**:
   ```javascript
   async function getWaitlist() {
     const response = await fetch('http://localhost:8080/api/waitlist/list');
     
     if (!response.ok) {
       throw new Error('Failed to get waitlist');
     }
     
     return await response.json();
   }
   ```

## Security Considerations

- The API is currently running on localhost for development
- For production, you should:
  - Add rate limiting
  - Implement proper authentication
  - Use HTTPS
  - Add input sanitization
  - Consider using a more robust database solution

## License

MIT 