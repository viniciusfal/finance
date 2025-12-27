import { writable } from 'svelte/store';
import { api } from '../api/api.js';

function createDashboardStore() {
	const { subscribe, set } = writable(null);
	let loading = false;

	return {
		subscribe,
		load: async () => {
			if (loading) return;
			loading = true;
			try {
				const summary = await api.dashboard.getSummary();
				set(summary);
			} catch (error) {
				console.error('Failed to load dashboard:', error);
			} finally {
				loading = false;
			}
		}
	};
}

export const dashboard = createDashboardStore();
