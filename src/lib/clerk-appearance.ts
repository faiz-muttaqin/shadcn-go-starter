/**
 * Clerk appearance configuration matching shadcn/ui theme
 * Uses CSS variables from theme.css for consistent styling
 */
export const clerkAppearance = {
  variables: {
    // Primary colors matching shadcn/ui theme
    colorPrimary: 'hsl(222.2 47.4% 11.2%)', // --primary
    colorBackground: 'hsl(228 0% 100%)', // --background
    colorText: 'hsl(222.2 47.4% 11.2%)', // --foreground
    colorTextSecondary: 'hsl(215.4 16.3% 46.9%)', // --muted-foreground
    colorDanger: 'hsl(0 72% 51%)', // --destructive
    colorSuccess: 'hsl(142 71% 45%)',
    colorInputBackground: 'hsl(0 0% 100%)',
    colorInputText: 'hsl(222.2 47.4% 11.2%)',

    // Border and spacing
    borderRadius: '0.5rem', // 8px

    // Typography
    fontFamily: 'inherit',
    fontSize: '0.875rem',
  },
  elements: {
    // Root card
    rootBox: 'w-full max-w-md',
    cardBox: '!shadow-none !border-none !rounded-none',
    card: '!shadow-none !border-none !rounded-none !m-0 !pt-0 !px-3',
    header: '!hidden ',
    // Logo
    logoBox: 'mb-6',
    logoImage: 'h-8 w-auto',

    // Header
    headerTitle: 'hidden text-2xl font-semibold tracking-tight text-foreground',
    headerSubtitle: 'hidden text-sm text-muted-foreground mt-2',


    // Form elements
    formFieldRow: 'space-y-2',
    formFieldLabel: 'text-sm font-medium leading-none text-foreground',
    formFieldLabelRow: 'flex items-center justify-between',
    formFieldInput:
      'flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50',
    formFieldInputShowPasswordButton:
      'text-muted-foreground hover:text-foreground',
    formFieldAction:
      'text-sm text-primary hover:text-primary/90 hover:underline',

    // Buttons
    formButtonPrimary:
      'inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2 w-full',
    formButtonReset:
      'inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2',

    // Footer
    footer: 'mt-6',
    footerAction: 'text-sm',
    footerActionText: 'text-muted-foreground',
    footerActionLink:
      'text-primary font-medium hover:text-primary/90 underline-offset-4 hover:underline',
    footerPages: 'mt-4',
    footerPagesLink:
      'text-sm text-muted-foreground hover:text-foreground underline-offset-4 hover:underline',

    // Divider
    dividerRow: '!py-0 !my-0 flex items-center gap-4',
    dividerLine: 'flex-1 bg-border h-px',
    dividerText: 'text-xs text-muted-foreground uppercase',

    // Social buttons
    socialButtons: 'space-y-2',
    socialButtonsBlockButton:
      'inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 px-4 py-2 w-full gap-2',
    socialButtonsBlockButtonText: 'font-normal',
    socialButtonsBlockButtonArrow: 'hidden',
    socialButtonsIconButton:
      'inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 border border-input bg-background hover:bg-accent hover:text-accent-foreground h-10 w-10',
    socialButtonsProviderIcon: 'w-4 h-4',

    // Alerts and messages
    alert: 'rounded-lg border border-border bg-background px-4 py-3',
    alertText: 'text-sm',

    // Identity preview
    identityPreviewText: 'text-sm text-muted-foreground',
    identityPreviewEditButton:
      'text-sm text-primary hover:text-primary/90 hover:underline',
    identityPreviewEditButtonIcon: 'w-3 h-3',

    // Form errors
    formFieldErrorText: 'text-sm font-medium text-destructive',
    formFieldSuccessText: 'text-sm font-medium text-success',
    formFieldWarningText: 'text-sm font-medium',
    formFieldInfoText: 'text-sm text-muted-foreground',
    formFieldHintText: 'text-xs text-muted-foreground',

    // Loading
    spinner: 'text-primary',

    // OTP input
    otpCodeFieldInput:
      'border border-input bg-background text-center text-lg font-medium rounded-md focus:ring-2 focus:ring-ring focus:ring-offset-2',

    // Internal card (for nested components)
    internal: 'space-y-4',
  },
  layout: {
    socialButtonsPlacement: 'bottom' as const,
    socialButtonsVariant: 'blockButton' as const,
    shimmer: true,
    logoPlacement: 'inside' as const,
  },
} as const

/**
 * Dark mode appearance configuration
 * Automatically applied when dark mode is active
 */
export const clerkDarkAppearance = {
  variables: {
    colorPrimary: 'hsl(210 40% 98%)', // Lighter for dark mode
    colorBackground: '#020919', // --background (dark)
    colorText: 'hsl(210 40% 98%)', // --foreground (dark)
    colorTextSecondary: 'hsl(215.4 16.3% 56.9%)', // --muted-foreground (dark)
    colorDanger: 'hsl(0 72% 51%)',
    colorSuccess: 'hsl(142 71% 45%)',
    colorInputBackground: 'hsl(217.2 32.6% 17.5%)', // --input (dark)
    colorInputText: 'hsl(210 40% 98%)',
    borderRadius: '0.5rem',
    fontFamily: 'inherit',
    fontSize: '0.875rem',
  },
  elements: {
    ...clerkAppearance.elements,
    card: 'bg-transparent shadow-none border-none rounded-lg !m-0 !pt-0 !px-3',

    // Override primary button for dark mode - white bg with dark text
    formButtonPrimary:
      'inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-white text-black [&>span]:text-black hover:bg-slate-100 h-10 px-4 py-2 w-full',
    buttonArrowIcon: 'text-black',
    // Override social buttons for dark mode - white bg with dark text
    socialButtons:
      'inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-white text-slate-900 border border-slate-200 hover:bg-slate-100 h-10 px-0 py-0 w-full gap-2',
    socialButtonsIconButton:
      'inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-white text-slate-900 border border-slate-200 hover:bg-slate-100 h-10 w-10',
    socialButtonsBlockButtonText: 'text-black font-normal',
  },
  layout: clerkAppearance.layout,
} as const
