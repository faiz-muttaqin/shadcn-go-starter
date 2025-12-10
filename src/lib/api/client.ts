/* eslint-disable @typescript-eslint/no-explicit-any */
import { auth } from '@/lib/firebase';

const BASE_URL = import.meta.env.VITE_BACKEND || '/api';

export interface ApiResponse<T = any> {
	success: boolean;
	message?: string;
	draw?: number;
	recordsTotal?: number;
	recordsFiltered?: number;
	data?: T;
	table?: Record<string, unknown>;
	error?: string;
	errors?: Record<string, string[] | string> | string[];
}

// Extend RequestInit to accept a `data` payload and `params` for query string
export type ApiRequestInit = RequestInit & {
	data?: unknown;
	params?: Record<string, string | number | boolean>;
};

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
					// fallback - don't overwrite localStorage when token is empty
					return localStorage.getItem('firebase_id_token');
				} else {
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

	private async request<T = any>(endpoint: string, options: ApiRequestInit = {}): Promise<ApiResponse<T>> {
		const token = await this.getAuthToken();

		// Extract our custom properties so we don't pass them to fetch
		const { data: dataPayload, params, ...rest } = options;

		const headers: Record<string, string> = {};

		// Merge existing headers
		if (rest.headers) {
			const existingHeaders = new Headers(rest.headers as HeadersInit);
			existingHeaders.forEach((value, key) => {
				headers[key] = value;
			});
		}

		// If body is NOT FormData, default to JSON content-type unless caller provided a header
		if (!(rest.body instanceof FormData) && !(dataPayload instanceof FormData) && !headers['Content-Type']) {
			headers['Content-Type'] = 'application/json';
		}

		if (token) {
			headers['Authorization'] = `Bearer ${token}`;
		}

		// Build URL with params if provided
		let url = `${this.baseURL}${endpoint}`;
		if (params && Object.keys(params).length > 0) {
			const searchParams = new URLSearchParams();
			Object.entries(params).forEach(([k, v]) => searchParams.append(k, String(v)));
			const sep = url.includes('?') ? '&' : '?';
			url = `${url}${sep}${searchParams.toString()}`;
		}

		// Determine body to send: priority -> explicit rest.body, then dataPayload
		let bodyToSend: BodyInit | undefined = undefined;
		if (rest.body instanceof FormData) {
			bodyToSend = rest.body as BodyInit;
		} else if (dataPayload instanceof FormData) {
			bodyToSend = dataPayload as BodyInit;
		} else if (rest.body !== undefined) {
			// caller explicitly set body (could be string already)
			bodyToSend = rest.body as BodyInit;
		} else if (dataPayload !== undefined) {
			// JSON-encode the provided data payload
			bodyToSend = JSON.stringify(dataPayload);
		}

		try {
			const fetchInit: RequestInit = {
				...rest,
				headers,
				body: bodyToSend,
			};

			// If method is GET/HEAD, browsers ignore body — ensure we don't send it
			const method = (fetchInit.method || 'GET').toUpperCase();
			if (method === 'GET' || method === 'HEAD') {
				delete (fetchInit as any).body;
			}

			const response = await fetch(url, fetchInit);

			// Attempt to parse JSON, but guard against empty responses
			let data: ApiResponse<T> = {} as ApiResponse<T>;
			try {
				data = (await response.json()) as ApiResponse<T>;
			} catch {
				// ignore parse errors — keep data as empty object
			}

			// If backend says unauthorized, clear stored firebase token so we
			// don't keep optimistically assuming auth on next load.
			if (response.status === 401) {
				try {
					localStorage.removeItem('firebase_id_token');
				} catch {
					// ignore
				}
				throw new Error((data && (data.message || data.error)) || 'Unauthorized');
			}

			// Save table data to localStorage if present
			if (data && data.table && typeof data.table === 'object') {
				try {
					localStorage.setItem('app_table', JSON.stringify(data.table));
				} catch (_e) {
					console.error('Failed to save app_table to localStorage:', _e);
				}
			}

			if (!response.ok) {
				throw new Error((data && (data.message || data.error)) || 'Request failed');
			}

			return data;
		} catch (error) {
			console.error('API request failed:', error);
			throw error;
		}
	}

	async get<T = any>(endpoint: string, options?: ApiRequestInit): Promise<ApiResponse<T>> {
		return this.request<T>(endpoint, { ...(options || {}), method: 'GET' });
	}

	async post<T = any>(endpoint: string, body?: unknown, options?: ApiRequestInit): Promise<ApiResponse<T>> {
		return this.request<T>(endpoint, { ...(options || {}), method: 'POST', data: body });
	}

	async put<T = any>(endpoint: string, body?: unknown, options?: ApiRequestInit): Promise<ApiResponse<T>> {
		return this.request<T>(endpoint, { ...(options || {}), method: 'PUT', data: body });
	}

	async delete<T = any>(endpoint: string, options?: ApiRequestInit): Promise<ApiResponse<T>> {
		return this.request<T>(endpoint, { ...(options || {}), method: 'DELETE' });
	}

	async patch<T = any>(endpoint: string, body?: unknown, options?: ApiRequestInit): Promise<ApiResponse<T>> {
		return this.request<T>(endpoint, { ...(options || {}), method: 'PATCH', data: body });
	}
}

export const apiClient = new ApiClient();
