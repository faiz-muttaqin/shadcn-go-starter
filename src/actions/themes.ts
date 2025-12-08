import { z } from 'zod'
import { themeStylesSchema, type ThemeStyles, type Theme } from '@/types/theme'
import { apiClient } from '@/lib/api/client'

// Validation schemas
const createThemeSchema = z.object({
  name: z
    .string()
    .min(1, 'Theme name cannot be empty')
    .max(50, 'Theme name too long'),
  styles: themeStylesSchema,
})

const updateThemeSchema = z.object({
  id: z.string().min(1, 'Theme ID required'),
  name: z
    .string()
    .min(1, 'Theme name cannot be empty')
    .max(50, 'Theme name too long')
    .optional(),
  styles: themeStylesSchema.optional(),
})

// Custom error types
export class UnauthorizedError extends Error {
  constructor(message = 'Unauthorized') {
    super(message)
    this.name = 'UnauthorizedError'
  }
}

export class ValidationError extends Error {
  constructor(
    message = 'Validation failed',
    public details?: unknown
  ) {
    super(message)
    this.name = 'ValidationError'
  }
}

export class ThemeNotFoundError extends Error {
  constructor(message = 'Theme not found') {
    super(message)
    this.name = 'ThemeNotFoundError'
  }
}

export class ThemeLimitError extends Error {
  constructor(message = 'Theme limit reached') {
    super(message)
    this.name = 'ThemeLimitError'
  }
}

// Helper to handle API errors
function handleApiError(error: unknown): never {
  if (error instanceof Error) {
    // Check if it's an HTTP error with specific status
    if (
      error.message.includes('Unauthorized') ||
      error.message.includes('401')
    ) {
      throw new UnauthorizedError()
    }
    if (error.message.includes('not found') || error.message.includes('404')) {
      throw new ThemeNotFoundError()
    }
    if (error.message.includes('limit') || error.message.includes('403')) {
      throw new ThemeLimitError(error.message)
    }
  }
  throw error
}

// Get all themes for current user
export async function getThemes(): Promise<Theme[]> {
  try {
    const response = await apiClient.get<Theme[]>('/themes')

    if (!response.success || !response.data) {
      throw new Error(response.message || 'Failed to fetch themes')
    }

    return response.data
  } catch (error) {
    console.error('getThemes error:', error)
    handleApiError(error)
  }
}

// Get single theme by ID
export async function getTheme(themeId: string): Promise<Theme> {
  try {
    if (!themeId) {
      throw new ValidationError('Theme ID required')
    }

    const response = await apiClient.get<Theme>(`/themes/${themeId}`)

    if (!response.success || !response.data) {
      throw new ThemeNotFoundError()
    }

    return response.data
  } catch (error) {
    console.error('getTheme error:', error)
    handleApiError(error)
  }
}

// Create new theme
export async function createTheme(formData: {
  name: string
  styles: ThemeStyles
}): Promise<Theme> {
  try {
    const validation = createThemeSchema.safeParse(formData)
    if (!validation.success) {
      throw new ValidationError('Invalid input', validation.error.format())
    }

    const { name, styles } = validation.data

    const response = await apiClient.post<Theme>('/themes', {
      name,
      styles,
    })

    if (!response.success || !response.data) {
      throw new Error(response.message || 'Failed to create theme')
    }

    return response.data
  } catch (error) {
    console.error('createTheme error:', error)
    handleApiError(error)
  }
}

// Update existing theme
export async function updateTheme(formData: {
  id: string
  name?: string
  styles?: ThemeStyles
}): Promise<Theme> {
  try {
    const validation = updateThemeSchema.safeParse(formData)
    if (!validation.success) {
      throw new ValidationError('Invalid input', validation.error.format())
    }

    const { id, name, styles } = validation.data

    if (!name && !styles) {
      throw new ValidationError('No update data provided')
    }

    const updateData: { name?: string; styles?: ThemeStyles } = {}
    if (name) updateData.name = name
    if (styles) updateData.styles = styles

    const response = await apiClient.patch<Theme>(`/themes/${id}`, updateData)

    if (!response.success || !response.data) {
      throw new Error(response.message || 'Failed to update theme')
    }

    return response.data
  } catch (error) {
    console.error('updateTheme error:', error)
    handleApiError(error)
  }
}

// Delete theme
export async function deleteTheme(
  themeId: string
): Promise<{ id: string; name: string }> {
  try {
    if (!themeId) {
      throw new ValidationError('Theme ID required')
    }

    const response = await apiClient.delete<{ id: string; name: string }>(
      `/themes/${themeId}`
    )

    if (!response.success || !response.data) {
      throw new Error(response.message || 'Failed to delete theme')
    }

    return response.data
  } catch (error) {
    console.error('deleteTheme error:', error)
    handleApiError(error)
  }
}
