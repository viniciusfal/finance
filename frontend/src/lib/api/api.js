// Usa variável de ambiente se disponível, senão usa /api (proxy local)
const API_BASE = import.meta.env.PUBLIC_API_URL || '/api';

async function request(endpoint, options = {}) {
	const response = await fetch(`${API_BASE}${endpoint}`, {
		headers: {
			'Content-Type': 'application/json',
			...options.headers
		},
		...options
	});

	if (!response.ok) {
		const error = await response.json().catch(() => ({ error: 'Unknown error' }));
		throw new Error(error.error || 'Request failed');
	}

	return response.json();
}

export const api = {
	// Transactions
	transactions: {
		getAll: () => request('/transactions'),
		getById: (id) => request(`/transactions/${id}`),
		create: (data) => request('/transactions', { method: 'POST', body: JSON.stringify(data) }),
		update: (id, data) => request(`/transactions/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
		delete: (id) => request(`/transactions/${id}`, { method: 'DELETE' }),
		payInstallment: (transactionId, installmentNumber) =>
			request(`/transactions/${transactionId}/installments/${installmentNumber}/pay`, { method: 'POST' })
	},

	// Categories
	categories: {
		getAll: () => request('/categories'),
		getById: (id) => request(`/categories/${id}`),
		create: (data) => request('/categories', { method: 'POST', body: JSON.stringify(data) }),
		update: (id, data) => request(`/categories/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
		delete: (id) => request(`/categories/${id}`, { method: 'DELETE' })
	},

	// Dashboard
	dashboard: {
		getSummary: () => request('/dashboard/summary')
	}
};

