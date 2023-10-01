# Go Backend App with Gin, Firebase Authentication, and MongoDB

This is a simple Go backend application built with the Gin framework. It uses Firebase authentication and MongoDB as the database. The application allows users to create and fetch notes, with all operations protected by Firebase authentication.

## Prerequisites

Before running the application, make sure you have the following:

1. **Firebase Project with Private Configuration:**
   - Use your Firebase private configuration and save it in a file named `firebase.json`.

2. **MongoDB Connection String:**
   - Create a `.env` file in the project root directory and add your MongoDB connection string in it.

## Running the Application

To run or debug the main application, follow these steps:

1. Open the project in a remote container environment.

2. Install the required Go packages.

3. Run the app in normal or debug mode .vscode/launch.json helps on that
## Routes

The application provides the following routes:

### Register a New User

- **Route:** `/register`

### Create a New Note

- **Route:** `/create-note`

### Get All Notes

- **Route:** `/notes`


### Get Note by ID

- **Route:** `/notes/:id`


