import type { ThemeStyles } from '@/types/theme'

/**
 * Applies theme styles to CSS variables on the document root element.
 * This function dynamically updates CSS custom properties based on the theme state.
 *
 * @param styles - The theme styles object containing both light and dark mode styles
 * @param mode - The current theme mode ('light' or 'dark')
 *
 * @example
 * ```ts
 * const themeStyles = {
 *   light: { primary: "oklch(0.5 0.2 250)", ... },
 *   dark: { primary: "oklch(0.8 0.1 250)", ... }
 * };
 * applyThemeStyles(themeStyles, 'light');
 * // Sets --primary, --background, etc. on document.documentElement
 * ```
 */
export function applyThemeStyles(styles: ThemeStyles, mode: 'light' | 'dark') {
  if (typeof window === 'undefined') return

  const root = document.documentElement
  const modeStyles = styles[mode]

  if (!modeStyles) {
    console.warn(`No styles found for mode: ${mode}`)
    return
  }

  // Apply each style property as a CSS variable
  Object.entries(modeStyles).forEach(([key, value]) => {
    // Convert property name to CSS variable format (e.g., "primary" -> "--primary")
    const cssVarName = `--${key}`
    root.style.setProperty(cssVarName, value)
  })
}

/**
 * Resets theme styles to default by removing custom properties
 * This allows the CSS file defaults to take over
 */
export function resetThemeStyles() {
  if (typeof window === 'undefined') return

  const root = document.documentElement

  // Get computed styles to find all custom properties
  const computedStyle = getComputedStyle(root)

  // Remove all custom properties that start with --
  Array.from(computedStyle).forEach((property) => {
    if (property.startsWith('--')) {
      root.style.removeProperty(property)
    }
  })
}
