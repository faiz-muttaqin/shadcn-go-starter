// ========================================
// API Response Types
// ========================================

export interface ApiResponse<T = unknown> {
  success: boolean
  message: string
  data: T
  table?: TableSettingsMap
}

// ========================================
// User Types
// ========================================

export interface AbilityRule {
  id: number
  role_id: number
  created_at: string
  updated_at: string
  subject: string
  read: boolean
  update: boolean
  create: boolean
  delete: boolean
}

export interface Role {
  id: number
  created_at: string
  updated_at: string
  DeletedAt: string | null
  role_name: string
  alias: string
  created_by: number
  ability_rules: AbilityRule[]
}

export interface User {
  id: number
  external_id: string
  created_at: string
  updated_at: string
  deleted_at: string | null
  verification_status: string
  username: string
  first_name: string
  last_name: string
  phone_number: string
  avatar: string
  email: string
  status: string
  session: string
  last_login: string
  role_id: number
  role: Role
}

// ========================================
// Table Settings Types
// ========================================

export interface TableColumn {
  name: string
  data: string
  type: string
  visible: boolean
  visibility: boolean
  sortable: boolean
  filterable: boolean
  editable: boolean
  passwordable: boolean
  selection?: string // URL for select options, or empty string for text search
}

export interface TableSettings {
  checkable: boolean
  creatable: boolean
  editable: boolean
  deletable: boolean
  column: TableColumn[]
  row: number
  row_opt: number[]
  sort: string
  table_name: string
  url: string
}

export interface TableSettingsMap {
  [key: string]: TableSettings // e.g., "t_users": TableSettings
}

// ========================================
// Auth Response Types
// ========================================

export type LoginResponse = ApiResponse<User>
