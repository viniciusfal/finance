<script>
	import { onMount } from 'svelte';
	import { categories } from '$lib/stores/categories';
	import { Plus, Edit, Trash2 } from 'lucide-svelte';
	import CategoryForm from '$lib/components/CategoryForm.svelte';

	let showForm = false;
	let editingCategory = null;

	onMount(() => {
		categories.load();
	});

	function handleEdit(category) {
		editingCategory = category;
		showForm = true;
	}

	function handleDelete(category) {
		if (confirm('Tem certeza que deseja excluir esta categoria?')) {
			categories.delete(category.id);
		}
	}

	function handleFormClose() {
		showForm = false;
		editingCategory = null;
	}

	function handleFormSubmit() {
		categories.load();
		handleFormClose();
	}
</script>

<svelte:head>
	<title>Categorias - Financy</title>
</svelte:head>

<div class="space-y-6">
	<div class="flex items-center justify-between">
		<h1 class="text-3xl font-bold text-gray-900">Categorias</h1>
		<button
			on:click={() => {
				showForm = true;
				editingCategory = null;
			}}
			class="flex items-center gap-2 px-4 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700 transition-colors"
		>
			<Plus class="w-5 h-5" />
			Nova Categoria
		</button>
	</div>

	<!-- Grid de Categorias -->
	<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
		{#each $categories as category}
			<div class="bg-white rounded-lg shadow p-6 hover:shadow-md transition-shadow">
				<div class="flex items-start justify-between mb-4">
					<div
						class="w-16 h-16 rounded-full flex items-center justify-center text-2xl"
						style="background-color: {category.color}20"
					>
						<span>{category.icon}</span>
					</div>
					<div class="flex gap-2">
						<button
							on:click={() => handleEdit(category)}
							class="p-2 text-gray-600 hover:text-primary-600 hover:bg-primary-50 rounded transition-colors"
						>
							<Edit class="w-4 h-4" />
						</button>
						<button
							on:click={() => handleDelete(category)}
							class="p-2 text-gray-600 hover:text-red-600 hover:bg-red-50 rounded transition-colors"
						>
							<Trash2 class="w-4 h-4" />
						</button>
					</div>
				</div>
				<h3 class="text-lg font-semibold text-gray-900 mb-1">{category.name}</h3>
				{#if category.description}
					<p class="text-sm text-gray-600">{category.description}</p>
				{/if}
			</div>
		{:else}
			<div class="col-span-full p-12 text-center text-gray-500">
				Nenhuma categoria encontrada. Crie uma nova categoria para começar.
			</div>
		{/each}
	</div>
</div>

{#if showForm}
	<CategoryForm category={editingCategory} on:close={handleFormClose} on:submit={handleFormSubmit} />
{/if}

