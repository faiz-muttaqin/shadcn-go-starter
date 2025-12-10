# Shadcn Admin Go Starter Dashboard

A full-stack admin dashboard starter built with **React + Vite** frontend and **Go + Gin** backend. Features a beautiful, responsive UI crafted with Shadcn UI components, complete REST API, authentication, and database integration.

This project combines a modern React frontend with a powerful Go backend, providing a production-ready starter template for building admin dashboards and data-driven applications.

## Features

### Frontend
- 🎨 Light/dark mode with system preference detection
- 📱 Fully responsive design for mobile, tablet, and desktop
- ♿ Accessible components following ARIA guidelines
- 🔍 Global search command palette (Cmd/Ctrl + K)
- 🌐 RTL (Right-to-Left) language support
- 📊 Data tables with sorting, filtering, and pagination
- 📈 Charts and visualizations with Recharts
- 🎯 Type-safe routing with TanStack Router


### Backend
- ⚡ High-performance Go backend with Gin framework
- 🗄️ GORM for database operations with multiple DB support
- 🔒 JWT-based authentication and authorization
- 📝 Automatic OpenAPI/Swagger documentation generation
- 🔄 RESTful API with CRUD operations
- 📤 File upload handling (images, documents, media)
- 🎯 Generic CRUD handlers for rapid development
- 🔍 Advanced filtering, sorting, and search capabilities
- 📦 Batch operations support (bulk create, update, delete)
- 📊 Excel template generation for batch uploads

### Developer Experience
- 🚀 Fast development with Vite HMR
- 🛠️ TypeScript for type safety
- 🎨 TailwindCSS for styling
- 📦 Modular component architecture
- 🧪 ESLint + Prettier for code quality
- 🔥 Live reload for both frontend and backend
- 📚 Auto-generated API documentation

<details>
<summary>Customized Components (click to expand)</summary>

This project uses Shadcn UI components, but some have been slightly modified for better RTL (Right-to-Left) support and other improvements. These customized components differ from the original Shadcn UI versions.

If you want to update components using the Shadcn CLI (e.g., `npx shadcn@latest add <component>`), it's generally safe for non-customized components. For the listed customized ones, you may need to manually merge changes to preserve the project's modifications and avoid overwriting RTL support or other updates.

> If you don't require RTL support, you can safely update the 'RTL Updated Components' via the Shadcn CLI, as these changes are primarily for RTL compatibility. The 'Modified Components' may have other customizations to consider.

### Modified Components

- scroll-area
- sonner
- separator

### RTL Updated Components

- alert-dialog
- calendar
- command
- dialog
- dropdown-menu
- select
- table
- sheet
- sidebar
- switch

**Notes:**

- **Modified Components**: These have general updates, potentially including RTL adjustments.
- **RTL Updated Components**: These have specific changes for RTL language support (e.g., layout, positioning).
- For implementation details, check the source files in `src/components/ui/`.
- All other Shadcn UI components in the project are standard and can be safely updated via the CLI.

</details>

## Tech Stack

### Frontend
**UI Framework:** [React](https://react.dev/) 19 + [ShadcnUI](https://ui.shadcn.com) (TailwindCSS + RadixUI)

**Build Tool:** [Vite](https://vitejs.dev/) 7

**Routing:** [TanStack Router](https://tanstack.com/router/latest) with file-based routing

**State Management:** [Zustand](https://zustand-demo.pmnd.rs/) + [TanStack Query](https://tanstack.com/query/latest)

**Forms:** [React Hook Form](https://react-hook-form.com/) + [Zod](https://zod.dev/)

**Type Checking:** [TypeScript](https://www.typescriptlang.org/) 5.9

**Styling:** [TailwindCSS](https://tailwindcss.com/) 4

**Icons:** [Lucide Icons](https://lucide.dev/icons/), [Tabler Icons](https://tabler.io/icons)


**HTTP Client:** [Axios](https://axios-http.com/)

### Backend
**Language:** [Go](https://go.dev/) 1.23+

**Web Framework:** [Gin](https://gin-gonic.com/)

**ORM:** [GORM](https://gorm.io/) with support for MySQL, PostgreSQL, SQLite

**Validation:** [go-playground/validator](https://github.com/go-playground/validator)

**API Documentation:** Custom OpenAPI/Swagger generator

**File Processing:** [Excelize](https://github.com/qax-os/excelize) for Excel operations

**Logging:** [Logrus](https://github.com/sirupsen/logrus)

### Development Tools
**Linting/Formatting:** [ESLint](https://eslint.org/) & [Prettier](https://prettier.io/)

**Go Hot Reload:** [Air](https://github.com/cosmtrek/air)

**Task Runner:** [Concurrently](https://github.com/open-cli-tools/concurrently)

## Prerequisites

Before you begin, ensure you have the following installed on your system:

### Required
- **Node.js** >= 18.0.0 ([Download](https://nodejs.org/))
- **pnpm** >= 8.0.0 (Install: `npm install -g pnpm`)
- **Go** >= 1.23 ([Download](https://go.dev/dl/))
- **Git** ([Download](https://git-scm.com/downloads))

### Optional (for development)
- **Air** - Go hot reload tool (Install: `go install github.com/cosmtrek/air@latest`)
- **Make** - Build automation (usually pre-installed on Linux/Mac)

### Database
Choose one of the following databases:
- **MySQL** 8.0+ ([Download](https://dev.mysql.com/downloads/mysql/))
- **PostgreSQL** 14+ ([Download](https://www.postgresql.org/download/))
- **SQLite** 3+ (No installation needed, embedded database)

## Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/faiz-muttaqin/shadcn-admin-go-starter.git
cd shadcn-admin-go-starter
```

### 2. Install Frontend Dependencies

```bash
pnpm install
```

### 3. Install Backend Dependencies

```bash
go mod download
go mod tidy
```

### 4. Configure Environment Variables

Create a `.env` file in the root directory or set environment variables:

```bash
# Application Settings
APP_NAME="Shadcn Admin Go Starter"
APP_VERSION="2.2.786"
APP_DESCRIPTION="Admin Dashboard with Go Backend"
APP_ENV="development"  # development, staging, production

# Server Configuration
APP_PORT=8080
APP_HOST="localhost"
APP_PUBLIC_URL="http://localhost:8080"

# Frontend Base Path (for API routes)
VITE_BASE="/api"

# Database Configuration (choose one)
# SQLite (default - no setup needed)
DB_DRIVER="sqlite"
DB_DSN="./data/app.db"

# MySQL
# DB_DRIVER="mysql"
# DB_DSN="username:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

# PostgreSQL
# DB_DRIVER="postgres"
# DB_DSN="host=localhost user=username password=password dbname=dbname port=5432 sslmode=disable"

# API Documentation
APP_ENABLE_SWAGGER="true"
APP_DOCS_PATH="/docs"

# Security
APP_ENABLE_BEARER="true"
JWT_SECRET="your-secret-key-change-in-production"

# File Upload
MAX_UPLOAD_SIZE=10485760  # 10MB in bytes

# Logging
LOG_LEVEL="debug"  # debug, info, warn, error
LOG_OUTPUT="stdout"  # stdout, file, both
```

### 5. Initialize Database

The application will automatically create the database schema on first run. For manual migration:

```bash
go run main.go migrate
```

## Running the Application

### Option 1: Run Both Frontend and Backend Together (Recommended)

```bash
pnpm run dev:all
```

This will start:
- Frontend dev server on `http://localhost:5173`
- Backend API server on `http://localhost:8080`
- Auto-reload for both on file changes

### Option 2: Run Frontend and Backend Separately

**Terminal 1 - Frontend:**
```bash
pnpm run dev
```

**Terminal 2 - Backend:**
```bash
# With hot reload (requires Air)
air

# Or without hot reload
go run main.go
```

### Option 3: Run Production Build

```bash
# Build everything
pnpm run build:all

# Run the compiled binary
./bin/shadcn-admin-go-starter

# Or on Windows
.\bin\shadcn-admin-go-starter.exe
```

## Accessing the Application

Once running, open your browser and navigate to:

- **Frontend:** http://localhost:5173 (development) or http://localhost:8080 (production)
- **Backend API:** http://localhost:8080/api
- **API Documentation:** http://localhost:8080/docs
- **Health Check:** http://localhost:8080/api/health

## Run Locally

Clone the project

```bash
  git clone https://github.com/faiz-muttaqin/shadcn-admin-go-starter.git
```

Go to the project directory

```bash
  cd shadcn-admin-go-starter
```

Install dependencies

```bash
  pnpm install
  go mod download
```

Configure environment variables (see Configuration section above)

```bash
  cp .env.example .env
  # Edit .env with your settings
```

Start the development servers

```bash
  pnpm run dev:all
```

## Available Scripts

### Frontend Scripts

```bash
# Start frontend development server
pnpm run dev

# Build frontend for production
pnpm run build

# Preview production build
pnpm run preview

# Lint frontend code
pnpm run lint

# Format code with Prettier
pnpm run format

# Check code formatting
pnpm run format:check

# Generate authenticated routes
pnpm run gen:routes

# Watch and auto-generate routes
pnpm run gen:routes:watch

# Check for unused dependencies
pnpm run knip
```

### Backend Scripts

```bash
# Run backend with hot reload
air

# Run backend without hot reload
go run main.go

# Build backend binary
go build -o bin/app main.go

# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Format Go code
go fmt ./...

# Lint Go code (requires golangci-lint)
golangci-lint run

# Update dependencies
go get -u ./...
go mod tidy
```

### Combined Scripts

```bash
# Run both frontend and backend
pnpm run dev:all

# Build both frontend and backend
pnpm run build:all
```

## Project Structure

```
shadcn-admin-go-starter/
├── backend/                    # Go backend
│   ├── cmd/                    # Command-line applications
│   ├── internal/               # Private application code
│   │   ├── database/          # Database connection and setup
│   │   ├── handler/           # HTTP request handlers
│   │   ├── middleware/        # Gin middlewares
│   │   ├── model/             # Database models
│   │   ├── routes/            # API route definitions
│   │   └── user/              # User management
│   ├── pkg/                   # Public libraries
│   │   ├── docs/              # API documentation generator
│   │   ├── logger/            # Logging utilities
│   │   ├── types/             # Shared types
│   │   └── util/              # Helper functions
│   ├── backend.go             # Backend entry point
│   └── docs.json              # Generated OpenAPI spec
├── src/                       # React frontend
│   ├── assets/                # Static assets
│   ├── components/            # React components
│   │   ├── ui/               # Shadcn UI components
│   │   ├── layout/           # Layout components
│   │   └── data-table/       # Table components
│   ├── config/                # App configuration
│   ├── context/               # React contexts
│   ├── features/              # Feature modules
│   │   ├── auth/             # Authentication
│   │   ├── dashboard/        # Dashboard
│   │   ├── users/            # User management
│   │   └── tasks/            # Task management
│   ├── hooks/                 # Custom React hooks
│   ├── lib/                   # Utility libraries
│   ├── routes/                # Route components
│   ├── services/              # API services
│   ├── stores/                # Zustand stores
│   ├── styles/                # Global styles
│   └── types/                 # TypeScript types
├── public/                    # Public static files
├── scripts/                   # Build and utility scripts
├── .env                       # Environment variables
├── .env.example              # Environment template
├── go.mod                     # Go dependencies
├── main.go                    # Go application entry
├── package.json               # Node.js dependencies
├── tsconfig.json             # TypeScript config
├── vite.config.ts            # Vite config
└── README.md                 # This file
```

## API Documentation

The backend automatically generates OpenAPI/Swagger documentation for all API endpoints.

### Accessing Documentation

Once the backend is running, visit: http://localhost:8080/docs

The documentation is automatically updated when you:
- Add new routes
- Modify request/response structures
- Update handler functions

### Key Features

- Interactive API explorer
- Request/response examples
- Authentication testing
- Schema validation
- Auto-generated from code

## Development Guide

### Creating a New API Endpoint

1. **Define the Model** (`backend/internal/model/`)
```go
type Product struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Name        string    `gorm:"column:name" json:"name"`
    Price       float64   `gorm:"column:price" json:"price"`
    CreatedAt   time.Time `json:"created_at"`
}
```

2. **Use Generic Handlers** (recommended) or create custom handlers
```go
// In routes file
router.GET("/products", handler.GET_DEFAULT_TableDataHandler(db, &model.Product{}, nil))
router.POST("/products", handler.POST_DEFAULT_TableDataHandler(db, &model.Product{}, nil))
router.PATCH("/products", handler.PATCH_DEFAULT_TableDataHandler(db, &model.Product{}, nil))
router.DELETE("/products", handler.DELETE_DEFAULT_TableDataHandler(db, &model.Product{}))
```

3. **Frontend API Service** (`src/services/`)
```typescript
export const productApi = {
  getAll: () => apiClient.get('/products'),
  create: (data) => apiClient.post('/products', data),
  update: (id, data) => apiClient.patch(`/products?id=${id}`, data),
  delete: (id) => apiClient.delete(`/products?id=${id}`)
}
```

### Generic CRUD Features

The generic handlers provide:
- ✅ Automatic CRUD operations
- ✅ Filtering and search
- ✅ Sorting and pagination
- ✅ Batch operations
- ✅ File uploads
- ✅ Validation
- ✅ Excel template generation

### File Upload Handling

File uploads are automatically handled for fields with specific types:
- `avatar`, `image` - Image files
- `file`, `document` - Documents
- `video` - Video files
- `audio` - Audio files
- `archive` - Archive files

### Database Migrations

The application uses GORM's AutoMigrate feature. Models are automatically migrated on startup.

For manual control:
```go
db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})
```

## Customization

### Theme Customization

Edit `src/styles/index.css` to customize colors, fonts, and other design tokens:

```css
@layer base {
  :root {
    --background: 0 0% 100%;
    --foreground: 222.2 84% 4.9%;
    --primary: 221.2 83.2% 53.3%;
    /* ... */
  }
}
```

### Adding Shadcn Components

```bash
npx shadcn@latest add button
npx shadcn@latest add card
npx shadcn@latest add dialog
```

### RTL Support

To enable RTL layout:

1. Set direction in your app:
```typescript
import { DirectionProvider } from '@/context/direction-provider'

<DirectionProvider defaultDirection="rtl">
  <App />
</DirectionProvider>
```

2. RTL-compatible components are already configured

## Deployment

### Frontend Deployment (Vercel, Netlify, etc.)

```bash
# Build frontend
pnpm run build

# Output will be in dist/
# Deploy dist/ folder to your hosting
```

### Backend Deployment

```bash
# Build backend binary
go build -o bin/app main.go

# Deploy binary to your server
# Make sure to set environment variables
```

### Docker Deployment (Coming Soon)

```bash
docker build -t shadcn-admin .
docker run -p 8080:8080 shadcn-admin
```

## Troubleshooting

### Common Issues

**Problem:** Port already in use
```bash
# Find and kill process on port 8080 (backend)
# Windows
netstat -ano | findstr :8080
taskkill /PID <PID> /F

# Linux/Mac
lsof -ti:8080 | xargs kill -9
```

**Problem:** Database connection failed
- Check your database credentials in `.env`
- Ensure database server is running
- For SQLite, check file permissions

**Problem:** Go modules not found
```bash
go mod download
go mod tidy
```

**Problem:** Frontend build fails
```bash
# Clear cache and reinstall
rm -rf node_modules pnpm-lock.yaml
pnpm install
```

**Problem:** CORS errors
- Check `VITE_BASE` matches backend API prefix
- Ensure backend CORS middleware is configured
