<!-- frontend/App.svelte -->
<script>
	import { onMount } from 'svelte';

	let todos = [];

	onMount(async () => {
		const response = await fetch('/todos');
		todos = await response.json();
	});

	async function createTodo() {
		/*TODO*/
	}

	async function updateTodo(id) {
		/*TODO*/
	}

	async function deleteTodo(id) {
		/*TODO*/
	}
</script>

<!---->

<div id=""