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

<!--TODO-->
<h1>Todos</h1>

<ul>
	{#each todos as todo}
		<li>
			<input type="checkbox" id={`completed-${todo.id}`} {todo.completed} />
			<input type="text" id={`title-${todo.id}`} value={todo.title} />
			<button on:click={() => updateTodo(todo.id)}>Update</button>
			<button on:click={() => deleteTodo(todo.id)}>Delete</button>
		</li>
	{/each}
</ul>

<input type="text" id="title" />
<button on:click={createTodo}>Create</button>