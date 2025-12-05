// Base API Response wrapper
export interface ApiResponse<T = unknown> {
  success: boolean
  data?: T
  error?: string
  message?: string
}

// Paginated Response
export interface PaginatedResponse<T> {
  success: boolean
  data: T[]
  pagination: {
    page: number
    pageSize: number
    total: number
    totalPages: number
  }
}

// User Types
export interface User {
  id: string
  email: string
  fullName: string
  username: string
  avatar?: string
  role: string
  status: 'active' | 'inactive' | 'invited' | 'suspended'
  createdAt: string
  updatedAt: string
  // Optional fields for compatibility with existing components
  firstName?: string
  lastName?: string
  phoneNumber?: string
}

export interface UserRole {
  label: string
  value: string
  icon: string
}

export interface CreateUserInput {
  email: string
  fullName: string
  username: string
  password: string
  roleId?: number
}

export interface UpdateUserInput {
  fullName?: string
  username?: string
  roleId?: number
  status?: 'active' | 'inactive' | 'suspended'
}

// Auth Types
export interface LoginInput {
  email: string
  password: string
}

export interface LoginResponse {
  success: boolean
  data: {
    userData: {
      id: string
      fullName: string
      username: string
      avatar?: string
      email: string
      role: string
    }
    accessToken: string
    userAbilityRules: Array<{
      action: string
      subject: string
    }>
  }
}

// Task Types (example)
export interface Task {
  id: string
  title: string
  status: 'todo' | 'in-progress' | 'done' | 'canceled'
  label: 'bug' | 'feature' | 'documentation'
  priority: 'low' | 'medium' | 'high'
}

// Error Types
export interface ApiError {
  error: string
  message?: string
  statusCode?: number
}
