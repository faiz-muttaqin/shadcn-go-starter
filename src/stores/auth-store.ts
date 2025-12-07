// DEPRECATED compatibility shim for legacy imports of `useAuthStore`.
// The app now uses `AuthProvider` + `useAuth()` context. Keep this file as a
// light shim so any late/leftover imports don't crash; prefer removing this
// entirely once the codebase is fully migrated.
import { clearClientAuth } from '../lib/auth-utils'

export const useAuthStore = {
  // mimic the old `getState().auth.reset()` usage
  getState: () => ({ auth: { reset: clearClientAuth } }),
}

export default useAuthStore
