<!-- frontend/App.svelte -->
<script>
	import { onMount } from 'svelte';

	let todos = [];

	onMount(async () => {
		const response = await fetch('/todos');
		todos = await response.json();
	});

	async function createTodo() {
		const title = document.getElementById('title').value;
		const response = await fetch('/todos', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ title, completed: false })
		});
		const newTodo = await response.json();
		todos.push(newTodo);
	}

	async function updateTodo(id) {
		/*TODO*/
	}

	async function deleteTodo(id) {
		/*TODO*/
	}
</script>

<!--TODO-->