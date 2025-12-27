<script>
	import '../app.css';
	import { onMount } from 'svelte';
	import { categories } from '$lib/stores/categories';
	import { transactions } from '$lib/stores/transactions';
	import { dashboard } from '$lib/stores/dashboard';
	import { Wallet, TrendingUp, TrendingDown, Tag } from 'lucide-svelte';

	let currentPath = '';

	onMount(() => {
		categories.load();
		transactions.load();
		dashboard.load();
	});

	$: currentPath = typeof window !== 'undefined' ? window.location.pathname : '';
</script>

<div class="min-h-screen bg-gray-50">
	<!-- Sidebar -->
	<aside class="fixed left-0 top-0 h-full w-64 bg-white shadow-lg">
		<div class="p-6">
			<div class="flex items-center gap-2 mb-8">
				<Wallet class="w-8 h-8 text-primary-600" />
				<h1 class="text-2xl font-bold text-gray-900">Financy</h1>
			</div>

			<nav class="space-y-2">
				<a
					href="/"
					class="flex items-center gap-3 px-4 py-3 rounded-lg transition-colors {currentPath === '/' ? 'bg-primary-50 text-primary-700' : 'text-gray-700 hover:bg-gray-100'}"
				>
					<TrendingUp class="w-5 h-5" />
					<span class="font-medium">Dashboard</span>
				</a>
				<a
					href="/transactions"
					class="flex items-center gap-3 px-4 py-3 rounded-lg transition-colors {currentPath.startsWith('/transactions') ? 'bg-primary-50 text-primary-700' : 'text-gray-700 hover:bg-gray-100'}"
				>
					<TrendingDown class="w-5 h-5" />
					<span class="font-medium">Transações</span>
				</a>
				<a
					href="/categories"
					class="flex items-center gap-3 px-4 py-3 rounded-lg transition-colors {currentPath.startsWith('/categories') ? 'bg-primary-50 text-primary-700' : 'text-gray-700 hover:bg-gray-100'}"
				>
					<Tag class="w-5 h-5" />
					<span class="font-medium">Categorias</span>
				</a>
			</nav>
		</div>
	</aside>

	<!-- Main Content -->
	<main class="ml-64 p-8">
		<slot />
	</main>
</div>

<style>
	:global(body) {
		margin: 0;
	}
</style>

