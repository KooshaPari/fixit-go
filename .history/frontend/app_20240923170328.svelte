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
		const title = document.getElementById(`title-${id}`).value;
		const response = await fetch(`/todos/${id}`, {
			method: 'PUT',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ title, completed: false })
		});
		const updatedTodo = await response.json();
		todos = todos.map(todo => todo.id === id ? updatedTodo : todo);
	}

	async function deleteTodo(id) {
		const response = await fetch(`/todos/${id}`, { method: 'DELETE' });
		todos = todos.filter(todo => todo.id !== id);
	}
</script>

<!--TODO-->