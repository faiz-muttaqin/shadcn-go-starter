import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { toast } from 'sonner'
import { usersService } from '@/services'
import type { CreateUserInput, UpdateUserInput } from '@/types/api'

/**
 * Query key factory for users
 */
export const userKeys = {
  all: ['users'] as const,
  lists: () => [...userKeys.all, 'list'] as const,
  list: (params?: Record<string, unknown>) =>
    [...userKeys.lists(), params] as const,
  details: () => [...userKeys.all, 'detail'] as const,
  detail: (id: string) => [...userKeys.details(), id] as const,
}

/**
 * Hook to fetch paginated users list
 */
export function useUsers(params?: {
  page?: number
  pageSize?: number
  status?: string[]
  role?: string[]
  username?: string
}) {
  return useQuery({
    queryKey: userKeys.list(params),
    queryFn: () => usersService.getUsers(params),
    staleTime: 30 * 1000, // 30 seconds
    // Enable persistent cache
    meta: { persist: true },
  })
}

/**
 * Hook to fetch a single user by ID
 */
export function useUser(userId: string) {
  return useQuery({
    queryKey: userKeys.detail(userId),
    queryFn: () => usersService.getUser(userId),
    enabled: !!userId,
    staleTime: 60 * 1000, // 1 minute
  })
}

/**
 * Hook to create a new user
 */
export function useCreateUser() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (input: CreateUserInput) => usersService.createUser(input),
    onSuccess: (data) => {
      // Invalidate users list to refetch
      queryClient.invalidateQueries({ queryKey: userKeys.lists() })

      toast.success('User created successfully', {
        description: `${data.data?.email} has been added`,
      })
    },
    onError: (error: Error) => {
      toast.error('Failed to create user', {
        description: error.message,
      })
    },
  })
}

/**
 * Hook to update an existing user
 */
export function useUpdateUser() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({
      userId,
      input,
    }: {
      userId: string
      input: UpdateUserInput
    }) => usersService.updateUser(userId, input),
    onSuccess: (data, variables) => {
      // Invalidate both list and detail queries
      queryClient.invalidateQueries({ queryKey: userKeys.lists() })
      queryClient.invalidateQueries({
        queryKey: userKeys.detail(variables.userId),
      })

      toast.success('User updated successfully', {
        description: `${data.data?.email} has been updated`,
      })
    },
    onError: (error: Error) => {
      toast.error('Failed to update user', {
        description: error.message,
      })
    },
  })
}

/**
 * Hook to delete a user
 */
export function useDeleteUser() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (userId: string) => usersService.deleteUser(userId),
    onSuccess: () => {
      // Invalidate users list
      queryClient.invalidateQueries({ queryKey: userKeys.lists() })

      toast.success('User deleted successfully')
    },
    onError: (error: Error) => {
      toast.error('Failed to delete user', {
        description: error.message,
      })
    },
  })
}

/**
 * Hook to bulk delete users
 */
export function useDeleteUsers() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (userIds: string[]) => usersService.deleteUsers(userIds),
    onSuccess: (_, variables) => {
      // Invalidate users list
      queryClient.invalidateQueries({ queryKey: userKeys.lists() })

      toast.success('Users deleted successfully', {
        description: `${variables.length} user(s) have been deleted`,
      })
    },
    onError: (error: Error) => {
      toast.error('Failed to delete users', {
        description: error.message,
      })
    },
  })
}

/**
 * Hook to update user role
 */
export function useUpdateUserRole() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ userId, roleId }: { userId: string; roleId: number }) =>
      usersService.updateUserRole(userId, roleId),
    onSuccess: (data, variables) => {
      // Invalidate both list and detail queries
      queryClient.invalidateQueries({ queryKey: userKeys.lists() })
      queryClient.invalidateQueries({
        queryKey: userKeys.detail(variables.userId),
      })

      toast.success('User role updated successfully', {
        description: `Role changed to ${data.data?.role}`,
      })
    },
    onError: (error: Error) => {
      toast.error('Failed to update user role', {
        description: error.message,
      })
    },
  })
}
