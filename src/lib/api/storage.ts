
export interface AppTable {
	user?: Record<string, unknown>;
	[key: string]: unknown;
}

export class StorageHelper {
	/**
	 * Get app_table from localStorage
	 */
	static getAppTable(): AppTable | null {
		try {
			const data = localStorage.getItem('app_table');
			return data ? JSON.parse(data) : null;
		} catch (e) {
			console.error('Failed to parse app_table:', e);
			return null;
		}
	}

	/**
	 * Save app_table to localStorage
	 */
	static saveAppTable(table: AppTable): void {
		try {
			localStorage.setItem('app_table', JSON.stringify(table));
		} catch (e) {
			console.error('Failed to save app_table:', e);
		}
	}

	/**
	 * Clear app_table from localStorage
	 */
	static clearAppTable(): void {
		localStorage.removeItem('app_table');
	}

	/**
	 * Get specific table data
	 */
	static getTableData(key: string): unknown {
		const table = this.getAppTable();
		return table ? table[key] : null;
	}
}
