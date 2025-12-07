import { auth } from '@/lib/firebase';

const BASE_URL = import.meta.env.VITE_BASE_URL || '/api';

export interface ApiResponse<T = unknown> {
	success: boolean;
	message: string;
    draw?: number;
	recordsTotal?: number;
    recordsFiltered?: number;
    data?: T;
	table?: Record<string, unknown>;
	error?: string;
}

export class ApiClient {
	private baseURL: string;

	constructor(baseURL: string = BASE_URL) {
		this.baseURL = baseURL;
	}

    private async getAuthToken(): Promise<string | null> {
        try {
            const user = auth.currentUser;
            if (user) {
                // Try to get token from Firebase (normal path)
                const token = await user.getIdToken();
                if (!token) {
                    localStorage.getItem('firebase_id_token');
                }else {
                    localStorage.setItem('firebase_id_token', token);
                }
                return token;
            }

            // No firebase user available — fallback to stored token
            const stored = localStorage.getItem('firebase_id_token');
            return stored ?? null;
        } catch (error) {
            console.error('Failed to get auth token from Firebase, falling back to localStorage:', error);
            const stored = localStorage.getItem('firebase_id_token');
            return stored ?? null;
        }
    }

	private async request<T = unknown>(
		endpoint: string,
		options: RequestInit = {}
	): Promise<ApiResponse<T>> {
		const token = await this.getAuthToken();

		const headers: Record<string, string> = {};

		// Merge existing headers
		if (options.headers) {
			const existingHeaders = new Headers(options.headers);
			existingHeaders.forEach((value, key) => {
				headers[key] = value;
			});
		}

		// If body is NOT FormData, default to JSON content-type unless caller provided a header
		if (!(options.body instanceof FormData) && !headers['Content-Type']) {
			headers['Content-Type'] = 'application/json';
		}

		if (token) {
			headers['Authorization'] = `Bearer ${token}`;
		}
		try {
			const response = await fetch(`${this.baseURL}${endpoint}`, {
				...options,
				headers,
			});

			const data: ApiResponse<T> = await response.json();

			// Save table data to localStorage if present
			if (data.table) {
				try {
					localStorage.setItem('app_table', JSON.stringify(data.table));
				} catch (e) {
					console.error('Failed to save app_table to localStorage:', e);
				}
			}

			if (!response.ok) {
				throw new Error(data.message || data.error || 'Request failed');
			}

			return data;
		} catch (error) {
			console.error('API request failed:', error);
			throw error;
		}
	}

	async get<T = unknown>(endpoint: string, options?: RequestInit): Promise<ApiResponse<T>> {
		return this.request<T>(endpoint, { ...options, method: 'GET' });
	}

	async post<T = unknown>(endpoint: string, body?: unknown, options?: RequestInit): Promise<ApiResponse<T>> {
		let bodyToSend: BodyInit | undefined = undefined;
		if (body instanceof FormData) {
			bodyToSend = body;
		} else if (body !== undefined) {
			bodyToSend = JSON.stringify(body);
		}
		return this.request<T>(endpoint, {
			...options,
			method: 'POST',
			body: bodyToSend,
		});
	}

	async put<T = unknown>(endpoint: string, body?: unknown, options?: RequestInit): Promise<ApiResponse<T>> {
		let bodyToSend: BodyInit | undefined = undefined;
		if (body instanceof FormData) {
			bodyToSend = body;
		} else if (body !== undefined) {
			bodyToSend = JSON.stringify(body);
		}
		return this.request<T>(endpoint, {
			...options,
			method: 'PUT',
			body: bodyToSend,
		});
	}

	async delete<T = unknown>(endpoint: string, options?: RequestInit): Promise<ApiResponse<T>> {
		return this.request<T>(endpoint, { ...options, method: 'DELETE' });
	}

	async patch<T = unknown>(endpoint: string, body?: unknown, options?: RequestInit): Promise<ApiResponse<T>> {
		return this.request<T>(endpoint, {
			...options,
			method: 'PATCH',
			body: body ? JSON.stringify(body) : undefined,
		});
	}
}

export const apiClient = new ApiClient();
