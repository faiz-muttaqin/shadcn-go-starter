import { auth as firebaseAuth } from '@/lib/firebase'

export function clearClientAuth() {
  try {
    localStorage.removeItem('firebase_id_token')
    localStorage.removeItem('firebase_user')
  } catch {
    // ignore
  }

  try {
    // Attempt to sign out firebase if available
    void firebaseAuth.signOut()
  } catch {
    // ignore
  }
}
