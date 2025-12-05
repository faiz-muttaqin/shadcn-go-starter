import apiClient from './api-client'
import type { LoginResponse, User, TableSettingsMap, TableSettings } from '@/types/auth'

const APP_TABLE_KEY = 'app_table'

/**
 * Fetch user data and table settings from backend after Clerk auth
 */
export async function syncUserData(): Promise<User | null> {
  try {
    const response = await apiClient.get<LoginResponse>('/auth/login')

    if (response.data.success) {
      // Save table settings to localStorage
      if (response.data.table) {
        saveTableSettings(response.data.table)
      }

      return response.data.data
    }

    return null
  } catch (error) {
    if (import.meta.env.DEV) {
      // eslint-disable-next-line no-console
      console.error('Failed to sync user data:', error)
    }
    return null
  }
}

/**
 * Save table settings to localStorage
 * Merges with existing settings
 */
export function saveTableSettings(newSettings: TableSettingsMap): void {
  try {
    // Get existing settings
    const existing = getTableSettings()

    // Merge new settings with existing
    const merged: TableSettingsMap = {
      ...existing,
      ...newSettings,
    }

    // Save to localStorage
    localStorage.setItem(APP_TABLE_KEY, JSON.stringify(merged))

    if (import.meta.env.DEV) {
      // eslint-disable-next-line no-console
      console.log('✅ Table settings saved:', merged)
    }
  } catch (error) {
    if (import.meta.env.DEV) {
      // eslint-disable-next-line no-console
      console.error('Failed to save table settings:', error)
    }
  }
}

/**
 * Get table settings from localStorage
 */
export function getTableSettings(): TableSettingsMap {
  try {
    const stored = localStorage.getItem(APP_TABLE_KEY)
    return stored ? (JSON.parse(stored) as TableSettingsMap) : {}
  } catch (error) {
    if (import.meta.env.DEV) {
      // eslint-disable-next-line no-console
      console.error('Failed to parse table settings:', error)
    }
    return {}
  }
}

/**
 * Get table settings for specific table
 */
export function getTableSettingsByName(tableName: string): TableSettings | null {
  const settings = getTableSettings()
  return settings[tableName] || null
}

/**
 * Clear all table settings
 */
export function clearTableSettings(): void {
  localStorage.removeItem(APP_TABLE_KEY)
}

/**
 * Update specific table settings
 */
export function updateTableSettings(tableName: string, settings: TableSettings): void {
  const key = `t_${tableName}`
  const existing = getTableSettings()

  saveTableSettings({
    ...existing,
    [key]: settings,
  })
}

/**
 * Debug: Log all table settings
 */
export function debugTableSettings(): TableSettingsMap {
  const settings = getTableSettings()
  if (import.meta.env.DEV) {
    // eslint-disable-next-line no-console
    console.log('📊 Current table settings:', settings)
  }
  return settings
}
