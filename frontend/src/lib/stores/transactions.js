import { writable } from 'svelte/store';
import { api } from '../api/api.js';

function createTransactionsStore() {
	const { subscribe, set, update } = writable([]);
	let loading = false;

	return {
		subscribe,
		load: async () => {
			if (loading) return;
			loading = true;
			try {
				const transactions = await api.transactions.getAll();
				set(transactions);
			} catch (error) {
				console.error('Failed to load transactions:', error);
			} finally {
				loading = false;
			}
		},
		add: async (transaction) => {
			try {
				const newTransaction = await api.transactions.create(transaction);
				update((transactions) => [newTransaction, ...transactions]);
				return newTransaction;
			} catch (error) {
				console.error('Failed to create transaction:', error);
				throw error;
			}
		},
		update: async (id, transaction) => {
			try {
				const updated = await api.transactions.update(id, transaction);
				update((transactions) =>
					transactions.map((t) => (t.id === id ? updated : t))
				);
				return updated;
			} catch (error) {
				console.error('Failed to update transaction:', error);
				throw error;
			}
		},
		delete: async (id) => {
			try {
				await api.transactions.delete(id);
				update((transactions) => transactions.filter((t) => t.id !== id));
			} catch (error) {
				console.error('Failed to delete transaction:', error);
				throw error;
			}
		},
		payInstallment: async (transactionId, installmentNumber) => {
			try {
				await api.transactions.payInstallment(transactionId, installmentNumber);
				await this.load();
			} catch (error) {
				console.error('Failed to pay installment:', error);
				throw error;
			}
		}
	};
}

export const transactions = createTransactionsStore();

