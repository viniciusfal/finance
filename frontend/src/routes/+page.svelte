<script>
	import { onMount } from 'svelte';
	import { dashboard } from '$lib/stores/dashboard';
	import { transactions } from '$lib/stores/transactions';
	import { categories } from '$lib/stores/categories';
	import { formatCurrency } from '$lib/utils/format.js';
	import { ArrowUp, ArrowDown, Wallet } from 'lucide-svelte';

	onMount(() => {
		dashboard.load();
		transactions.load();
		categories.load();
	});

	$: summary = $dashboard;
	$: recentTransactions = $transactions.slice(0, 5);
</script>

<svelte:head>
	<title>Dashboard - Financy</title>
</svelte:head>

<div class="space-y-8">
	<h1 class="text-3xl font-bold text-gray-900">Dashboard</h1>

	{#if summary}
		<!-- Cards de Resumo -->
		<div class="grid grid-cols-1 md:grid-cols-3 gap-6">
			<!-- Saldo Total -->
			<div class="bg-white rounded-lg shadow p-6">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm text-gray-600 mb-1">Saldo Total</p>
						<p class="text-2xl font-bold text-gray-900">{formatCurrency(summary.total_balance)}</p>
					</div>
					<Wallet class="w-12 h-12 text-primary-600" />
				</div>
			</div>

			<!-- Receitas do Mês -->
			<div class="bg-white rounded-lg shadow p-6">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm text-gray-600 mb-1">Receitas do Mês</p>
						<p class="text-2xl font-bold text-green-600">{formatCurrency(summary.monthly_income)}</p>
					</div>
					<ArrowUp class="w-12 h-12 text-green-600" />
				</div>
			</div>

			<!-- Despesas do Mês -->
			<div class="bg-white rounded-lg shadow p-6">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm text-gray-600 mb-1">Despesas do Mês</p>
						<p class="text-2xl font-bold text-red-600">{formatCurrency(summary.monthly_expense)}</p>
					</div>
					<ArrowDown class="w-12 h-12 text-red-600" />
				</div>
			</div>
		</div>

		<!-- Últimas Transações -->
		<div class="bg-white rounded-lg shadow">
			<div class="p-6 border-b border-gray-200">
				<h2 class="text-xl font-semibold text-gray-900">Últimas Transações</h2>
			</div>
			<div class="divide-y divide-gray-200">
				{#each recentTransactions as transaction}
					<div class="p-6 flex items-center justify-between hover:bg-gray-50">
						<div class="flex items-center gap-4">
							<div
								class="w-12 h-12 rounded-full flex items-center justify-center"
								style="background-color: {transaction.category?.color || '#6366F1'}20"
							>
								<span class="text-lg">{transaction.category?.icon || '💰'}</span>
							</div>
							<div>
								<p class="font-medium text-gray-900">{transaction.title}</p>
								<p class="text-sm text-gray-500">
									{transaction.category?.name || 'Sem categoria'}
								</p>
							</div>
						</div>
						<div class="text-right">
							<p
								class="font-semibold {transaction.type === 'income' ? 'text-green-600' : 'text-red-600'}"
							>
								{transaction.type === 'income' ? '+' : '-'} {formatCurrency(transaction.amount_cents)}
							</p>
							<p class="text-sm text-gray-500">{new Date(transaction.due_date).toLocaleDateString('pt-BR')}</p>
						</div>
					</div>
				{:else}
					<div class="p-6 text-center text-gray-500">
						Nenhuma transação encontrada
					</div>
				{/each}
			</div>
		</div>
	{:else}
		<div class="text-center py-12">
			<p class="text-gray-500">Carregando...</p>
		</div>
	{/if}
</div>

