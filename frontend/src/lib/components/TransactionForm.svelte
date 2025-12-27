<script>
	import { createEventDispatcher, onMount } from 'svelte';
	import { transactions } from '$lib/stores/transactions';
	import { categories } from '$lib/stores/categories';
	import { formatDateInput } from '$lib/utils/format.js';
	import { X } from 'lucide-svelte';

	export let transaction = null;

	const dispatch = createEventDispatcher();

	let formData = {
		title: '',
		description: '',
		amount_cents: 0,
		type: 'expense',
		category_id: null,
		due_date: '',
		is_recurring: false,
		is_installment: false,
		total_installments: 1
	};

	let loading = false;
	let error = '';

	onMount(() => {
		if (transaction) {
			formData = {
				title: transaction.title,
				description: transaction.description || '',
				amount_cents: transaction.amount_cents / 100, // Converter centavos para reais
				type: transaction.type,
				category_id: transaction.category_id,
				due_date: formatDateInput(transaction.due_date),
				is_recurring: transaction.is_recurring,
				is_installment: transaction.is_installment,
				total_installments: transaction.total_installments || 1
			};
		} else {
			// Data padrão: hoje
			const today = new Date();
			formData.due_date = formatDateInput(today);
		}
	});

	async function handleSubmit() {
		error = '';
		
		if (!formData.title) {
			error = 'Título é obrigatório';
			return;
		}

		if (formData.amount_cents <= 0) {
			error = 'Valor deve ser maior que zero';
			return;
		}

		if (!formData.due_date) {
			error = 'Data de vencimento é obrigatória';
			return;
		}

		if (formData.is_installment && formData.total_installments < 2) {
			error = 'Transação parcelada deve ter pelo menos 2 parcelas';
			return;
		}

		loading = true;

		try {
			const data = {
				...formData,
				amount_cents: Math.round(formData.amount_cents * 100), // Converter para centavos
				category_id: formData.category_id || null
			};

			if (transaction) {
				await transactions.update(transaction.id, data);
			} else {
				await transactions.add(data);
			}

			dispatch('submit');
		} catch (err) {
			error = err.message || 'Erro ao salvar transação';
		} finally {
			loading = false;
		}
	}

	function handleClose() {
		dispatch('close');
	}
</script>

<div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
	<div class="bg-white rounded-lg shadow-xl max-w-2xl w-full max-h-[90vh] overflow-y-auto">
		<div class="p-6 border-b border-gray-200 flex items-center justify-between">
			<h2 class="text-2xl font-bold text-gray-900">
				{transaction ? 'Editar Transação' : 'Nova Transação'}
			</h2>
			<button on:click={handleClose} class="text-gray-400 hover:text-gray-600">
				<X class="w-6 h-6" />
			</button>
		</div>

		<form on:submit|preventDefault={handleSubmit} class="p-6 space-y-4">
			{#if error}
				<div class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded">
					{error}
				</div>
			{/if}

			<div>
				<label class="block text-sm font-medium text-gray-700 mb-1">Título *</label>
				<input
					type="text"
					bind:value={formData.title}
					class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent"
					required
				/>
			</div>

			<div>
				<label class="block text-sm font-medium text-gray-700 mb-1">Descrição</label>
				<textarea
					bind:value={formData.description}
					class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent"
					rows="3"
				></textarea>
			</div>

			<div class="grid grid-cols-2 gap-4">
				<div>
					<label class="block text-sm font-medium text-gray-700 mb-1">Valor (R$) *</label>
					<input
						type="number"
						step="0.01"
						bind:value={formData.amount_cents}
						class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent"
						required
					/>
				</div>

				<div>
					<label class="block text-sm font-medium text-gray-700 mb-1">Tipo *</label>
					<select
						bind:value={formData.type}
						class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent"
					>
						<option value="expense">Despesa</option>
						<option value="income">Receita</option>
					</select>
				</div>
			</div>

			<div class="grid grid-cols-2 gap-4">
				<div>
					<label class="block text-sm font-medium text-gray-700 mb-1">Categoria</label>
					<select
						bind:value={formData.category_id}
						class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent"
					>
						<option value={null}>Sem categoria</option>
						{#each $categories as category}
							<option value={category.id}>{category.name}</option>
						{/each}
					</select>
				</div>

				<div>
					<label class="block text-sm font-medium text-gray-700 mb-1">Data de Vencimento *</label>
					<input
						type="date"
						bind:value={formData.due_date}
						class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent"
						required
					/>
				</div>
			</div>

			<div class="flex items-center gap-4">
				<label class="flex items-center gap-2">
					<input type="checkbox" bind:checked={formData.is_recurring} class="rounded" />
					<span class="text-sm text-gray-700">Transação Recorrente</span>
				</label>

				<label class="flex items-center gap-2">
					<input type="checkbox" bind:checked={formData.is_installment} class="rounded" />
					<span class="text-sm text-gray-700">Parcelada</span>
				</label>
			</div>

			{#if formData.is_installment}
				<div>
					<label class="block text-sm font-medium text-gray-700 mb-1">Número de Parcelas</label>
					<input
						type="number"
						min="2"
						bind:value={formData.total_installments}
						class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent"
					/>
					<p class="mt-1 text-xs text-gray-500">Cada parcela terá intervalo de 30 dias</p>
				</div>
			{/if}

			<div class="flex gap-3 pt-4">
				<button
					type="button"
					on:click={handleClose}
					class="flex-1 px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
				>
					Cancelar
				</button>
				<button
					type="submit"
					disabled={loading}
					class="flex-1 px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700 transition-colors disabled:opacity-50"
				>
					{loading ? 'Salvando...' : 'Salvar'}
				</button>
			</div>
		</form>
	</div>
</div>
