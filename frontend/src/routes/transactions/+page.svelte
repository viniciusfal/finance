<script>
	import { onMount } from 'svelte';
	import { transactions } from '$lib/stores/transactions';
	import { categories } from '$lib/stores/categories';
	import { formatCurrency, formatDate } from '$lib/utils/format.js';
	import { Plus, Edit, Trash2 } from 'lucide-svelte';
	import TransactionForm from '$lib/components/TransactionForm.svelte';

	let showForm = false;
	let editingTransaction = null;
	let filterType = 'all';

	onMount(() => {
		transactions.load();
		categories.load();
	});

	$: filteredTransactions = $transactions.filter((t) => {
		if (filterType === 'all') return true;
		return t.type === filterType;
	});

	function handleEdit(transaction) {
		editingTransaction = transaction;
		showForm = true;
	}

	function handleDelete(transaction) {
		if (confirm('Tem certeza que deseja excluir esta transação?')) {
			transactions.delete(transaction.id);
		}
	}

	function handleFormClose() {
		showForm = false;
		editingTransaction = null;
	}

	function handleFormSubmit() {
		transactions.load();
		handleFormClose();
	}
</script>

<svelte:head>
	<title>Transações - Financy</title>
</svelte:head>

<div class="space-y-6">
	<div class="flex items-center justify-between">
		<h1 class="text-3xl font-bold text-gray-900">Transações</h1>
		<button
			on:click={() => {
				showForm = true;
				editingTransaction = null;
			}}
			class="flex items-center gap-2 px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700 transition-colors"
		>
			<Plus class="w-5 h-5" />
			Nova Transação
		</button>
	</div>

	<!-- Filtros -->
	<div class="flex gap-2">
		<button
			on:click={() => (filterType = 'all')}
			class="px-4 py-2 rounded-lg transition-colors {filterType === 'all' ? 'bg-primary-600 text-white' : 'bg-white text-gray-700 hover:bg-gray-100'}"
		>
			Todas
		</button>
		<button
			on:click={() => (filterType = 'income')}
			class="px-4 py-2 rounded-lg transition-colors {filterType === 'income' ? 'bg-primary-600 text-white' : 'bg-white text-gray-700 hover:bg-gray-100'}"
		>
			Receitas
		</button>
		<button
			on:click={() => (filterType = 'expense')}
			class="px-4 py-2 rounded-lg transition-colors {filterType === 'expense' ? 'bg-primary-600 text-white' : 'bg-white text-gray-700 hover:bg-gray-100'}"
		>
			Despesas
		</button>
	</div>

	<!-- Lista de Transações -->
	<div class="bg-white rounded-lg shadow">
		<div class="divide-y divide-gray-200">
			{#each filteredTransactions as transaction}
				<div class="p-6 hover:bg-gray-50">
					<div class="flex items-center justify-between">
						<div class="flex items-center gap-4 flex-1">
							<div
								class="w-12 h-12 rounded-full flex items-center justify-center"
								style="background-color: {transaction.category?.color || '#6366F1'}20"
							>
								<span class="text-lg">{transaction.category?.icon || '💰'}</span>
							</div>
							<div class="flex-1">
								<div class="flex items-center gap-2">
									<h3 class="font-semibold text-gray-900">{transaction.title}</h3>
									{#if transaction.is_installment}
										<span class="px-2 py-1 text-xs bg-blue-100 text-blue-700 rounded">
											{transaction.total_installments}x
										</span>
									{/if}
								</div>
								<p class="text-sm text-gray-500">
									{transaction.category?.name || 'Sem categoria'} • {formatDate(transaction.due_date)}
								</p>
								{#if transaction.description}
									<p class="text-sm text-gray-600 mt-1">{transaction.description}</p>
								{/if}
								{#if transaction.is_installment && transaction.installments}
									<div class="mt-2 space-y-1">
										{#each transaction.installments as installment}
											<div class="flex items-center gap-2 text-xs">
												<span class="text-gray-500">Parcela {installment.installment_number}:</span>
												<span class="font-medium">{formatCurrency(installment.amount_cents)}</span>
												<span class="text-gray-500">- {formatDate(installment.due_date)}</span>
												<span
													class="px-2 py-0.5 rounded {installment.status === 'paid' ? 'bg-green-100 text-green-700' : 'bg-yellow-100 text-yellow-700'}"
												>
													{installment.status === 'paid' ? 'Pago' : 'Pendente'}
												</span>
											</div>
										{/each}
									</div>
								{/if}
							</div>
						</div>
						<div class="flex items-center gap-4">
							<div class="text-right">
								<p
									class="text-lg font-semibold {transaction.type === 'income' ? 'text-green-600' : 'text-red-600'}"
								>
									{transaction.type === 'income' ? '+' : '-'} {formatCurrency(transaction.amount_cents)}
								</p>
								<p
									class="text-sm px-2 py-1 rounded {transaction.status === 'paid' ? 'bg-green-100 text-green-700' : transaction.status === 'overdue' ? 'bg-red-100 text-red-700' : 'bg-yellow-100 text-yellow-700'}"
								>
									{transaction.status === 'paid' ? 'Pago' : transaction.status === 'overdue' ? 'Atrasado' : 'Pendente'}
								</p>
							</div>
							<div class="flex gap-2">
								<button
									on:click={() => handleEdit(transaction)}
									class="p-2 text-gray-600 hover:text-primary-600 hover:bg-primary-50 rounded transition-colors"
								>
									<Edit class="w-5 h-5" />
								</button>
								<button
									on:click={() => handleDelete(transaction)}
									class="p-2 text-gray-600 hover:text-red-600 hover:bg-red-50 rounded transition-colors"
								>
									<Trash2 class="w-5 h-5" />
								</button>
							</div>
						</div>
					</div>
				</div>
			{:else}
				<div class="p-12 text-center text-gray-500">
					Nenhuma transação encontrada
				</div>
			{/each}
		</div>
	</div>
</div>

{#if showForm}
	<TransactionForm
		transaction={editingTransaction}
		on:close={handleFormClose}
		on:submit={handleFormSubmit}
	/>
{/if}

