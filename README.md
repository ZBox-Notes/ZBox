# ZBox

## Setup Instructions

### Backend Setup

Create a `.env` file in the `backend` directory and copy the contents of `backend/.env.example` to the file, replacing the sample values with your desired values. If defaults are set in the `.env.example` file, you can leave them as is.

```bash
cd backend
docker-compose up
```

### Frontend Setup (Node.js 18+ required)

Create a `.env` file in the `frontend` directory and copy the contents of `frontend/.env.example` to the file, replacing the sample values with your desired values

Start the app using the following ocmmands

```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
npm install

# Start the development server
npm run dev
```

The application will be available at:

- Frontend: http://localhost:5173
- Backend: https://localhost:3000
