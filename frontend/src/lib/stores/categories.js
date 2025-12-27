import { writable } from 'svelte/store';
import { api } from '../api/api.js';

function createCategoriesStore() {
	const { subscribe, set, update } = writable([]);
	let loading = false;

	return {
		subscribe,
		load: async () => {
			if (loading) return;
			loading = true;
			try {
				const categories = await api.categories.getAll();
				set(categories);
			} catch (error) {
				console.error('Failed to load categories:', error);
			} finally {
				loading = false;
			}
		},
		add: async (category) => {
			try {
				const newCategory = await api.categories.create(category);
				update((categories) => [...categories, newCategory]);
				return newCategory;
			} catch (error) {
				console.error('Failed to create category:', error);
				throw error;
			}
		},
		update: async (id, category) => {
			try {
				const updated = await api.categories.update(id, category);
				update((categories) =>
					categories.map((c) => (c.id === id ? updated : c))
				);
				return updated;
			} catch (error) {
				console.error('Failed to update category:', error);
				throw error;
			}
		},
		delete: async (id) => {
			try {
				await api.categories.delete(id);
				update((categories) => categories.filter((c) => c.id !== id));
			} catch (error) {
				console.error('Failed to delete category:', error);
				throw error;
			}
		}
	};
}

export const categories = createCategoriesStore();
