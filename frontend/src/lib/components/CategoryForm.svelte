<script>
	import { createEventDispatcher, onMount } from 'svelte';
	import { categories } from '$lib/stores/categories';
	import { X } from 'lucide-svelte';

	export let category = null;

	const dispatch = createEventDispatcher();

	const colorOptions = [
		{ name: 'Verde', value: '#10B981' },
		{ name: 'Azul', value: '#3B82F6' },
		{ name: 'Roxo', value: '#8B5CF6' },
		{ name: 'Laranja', value: '#F59E0B' },
		{ name: 'Rosa', value: '#EC4899' },
		{ name: 'Amarelo', value: '#FBBF24' },
		{ name: 'Vermelho', value: '#EF4444' },
		{ name: 'Indigo', value: '#6366F1' }
	];

	const iconOptions = ['💰', '🍔', '🚗', '🛒', '🏠', '💊', '🎬', '📚', '🏋️', '✈️', '🎮', '💳', '🏦', '📱', '💻', '🎨'];

	let formData = {
		name: '',
		description: '',
		color: '#6366F1',
		icon: '💰'
	};

	let loading = false;
	let error = '';

	onMount(() => {
		if (category) {
			formData = {
				name: category.name,
				description: category.description || '',
				color: category.color,
				icon: category.icon
			};
		}
	});

	async function handleSubmit() {
		error = '';

		if (!formData.name) {
			error = 'Nome é obrigatório';
			return;
		}

		loading = true;

		try {
			const data = {
				...formData,
				description: formData.description || null
			};

			if (category) {
				await categories.update(category.id, data);
			} else {
				await categories.add(data);
			}

			dispatch('submit');
		} catch (err) {
			error = err.message || 'Erro ao salvar categoria';
		} finally {
			loading = false;
		}
	}

	function handleClose() {
		dispatch('close');
	}
</script>

<div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
	<div class="bg-white rounded-lg shadow-xl max-w-lg w-full">
		<div class="p-6 border-b border-gray-200 flex items-center justify-between">
			<h2 class="text-2xl font-bold text-gray-900">
				{category ? 'Editar Categoria' : 'Nova Categoria'}
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
				<label class="block text-sm font-medium text-gray-700 mb-1">Nome *</label>
				<input
					type="text"
					bind:value={formData.name}
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

			<div>
				<label class="block text-sm font-medium text-gray-700 mb-1">Cor</label>
				<div class="grid grid-cols-4 gap-2">
					{#each colorOptions as color}
						<button
							type="button"
							on:click={() => (formData.color = color.value)}
							class="h-12 rounded-lg border-2 transition-all {formData.color === color.value ? 'border-gray-900 scale-110' : 'border-gray-300'}"
							style="background-color: {color.value}"
							title={color.name}
						></button>
					{/each}
				</div>
			</div>

			<div>
				<label class="block text-sm font-medium text-gray-700 mb-1">Ícone</label>
				<div class="grid grid-cols-8 gap-2">
					{#each iconOptions as icon}
						<button
							type="button"
							on:click={() => (formData.icon = icon)}
							class="h-12 rounded-lg border-2 text-2xl transition-all {formData.icon === icon ? 'border-primary-600 bg-primary-50' : 'border-gray-300 hover:border-gray-400'}"
						>
							{icon}
						</button>
					{/each}
				</div>
			</div>

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

