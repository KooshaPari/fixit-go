<script>
	import { onMount, afterUpdate } from "svelte";
	import Modal from "./Modal.svelte";
	import { prevent_default } from "svelte/internal";
	import { fly, fade } from "svelte/transition";
	let todos = [];
	let tempid = -1;
	onMount(async () => {
		try {
			const response = await fetch("http://localhost:8081/api/todos");
			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`);
			}
			console.log("GRABBING TODOS");
			todos = await response.json();
		} catch (error) {
			console.error("Error fetching todos:", error);
		}
	});
	let showCreateModal = false;
	let showEditModal = false;
	async function loadEditPopup(id) {
		showEditModal = true;
		const response = await fetch(`http://localhost:8081/api/todos/${id}`, {
			method: "GET",
			headers: { "Content-Type": "application/json" },
		});
		const todo = await response.json();
		// Update the title, description, and dueDate variables
		title = todo.title;
		description = todo.description;
		dueDate = todo.dueDate;
		tempid = todo.id;
	}
	async function editTodo(title, description, dueDate) {
		const id = tempid;
		const newTodo = { id, title, description, dueDate };
		tempid = -1;
		try {
			const response = await fetch(`http://localhost:8081/api/todos/${id}`, {
				method: "PUT",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify(newTodo),
			});
			const updatedTodo = await response.json();
			todos = todos.map((todo) => (todo.id === id ? updatedTodo : todo));
		} catch (error) {
			console.error("Error editing todo:", error);
		}
	}

	async function createTodo(title, description, dueDate) {
		/**Generate ToDo Object, add to local array, then create fetch object to server to add to remote array*/

		const newTodo = { title, description, dueDate };
		try {
			const response = await fetch("http://localhost:8081/api/todos", {
				method: "POST",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify(newTodo),
			});
			const createdTodo = await response.json();
			todos = [...todos, createdTodo];
		} catch (error) {
			console.error("Error creating todo:", error);
		}
		showCreateModal = false;
	}
	async function deleteTodo(id) {
		console.log("DELETING: " + id);
		if (id !== undefined) {
			todos = todos.filter((todo) => todo.id !== id);
			const response = await fetch(`http://localhost:8081/api/todos/${id}`, {
				method: "DELETE",
				headers: { "Content-Type": "application/json" },
			});
		} else {
			console.error("Cannot delete todo item: id is undefined");
		}
	}
	let title;
	let description;
	let dueDate;
</script>

<main>
	<!---->

	<div id="header">
		<div id="logoCont">
			<h1>Fix It</h1>
		</div>
	</div>
	<div id="body">
		<div id="fixItApp" style="overflow: hidden;">
			<Modal bind:showModal={showEditModal}>
				<h2>Edit Todo Item</h2>
				<form>
					<input
						type="text"
						id="todo-title"
						bind:value={title}
						placeholder="Todo title"
					/>
					<input
						type="text"
						id="todo-description"
						bind:value={description}
						placeholder="Todo description"
					/>
					<input
						type="text"
						id="todo-dueDate"
						bind:value={dueDate}
						placeholder="Todo dueDate"
					/>
					<button
						on:click={(event) => {
							editTodo(title, description, dueDate);
						}}>Edit</button
					>
				</form>
			</Modal>
			<Modal bind:showModal={showCreateModal}>
				<h2>Create Todo Item</h2>
				<form>
					<input
						type="text"
						id="todo-title"
						bind:value={title}
						placeholder="Todo title"
					/>
					<input
						type="text"
						id="todo-description"
						bind:value={description}
						placeholder="Todo description"
					/>
					<input
						type="text"
						id="todo-dueDate"
						bind:value={dueDate}
						placeholder="Todo dueDate"
					/>
					<button on:click={() => createTodo(title, description, dueDate)}
						>Create</button
					>
				</form>
			</Modal>
			<div id="todoList">
				<button
					on:click={() => {
						showCreateModal = true;
					}}
					class="addBtn">Add</button
				>

				<ul id="todos">
					{#each todos as todo, index}
						<li
							class="todoLstItm"
							in:fly={{ y: 200, duration: 2000, delay: index * 100 }}
							out:fade={{ duration: 2000 }}
						>
							<div class="todoLstItmContLeft">
								<button
									on:click={() => deleteTodo(todo.id)}
									class="completeBtn todoBtn">&#10003;</button
								>
							</div>
							<div class="todoLstItmConRight">
								<h3 style="padding-inline: 1em; font-size: 0.75em; width: 40%;">
									{todo.title}
								</h3>
								<div style="display:flex; flex-direction:column; width: 40% ">
									<p class="todotxt">{todo.description}</p>
									<!-- due Date below-->
									<p class="todotxt">{todo.dueDate}</p>
								</div>
								<div class="todoLstItmConRightBtns">
									<button
										on:click={() => loadEditPopup(todo.id)}
										class="deleteBtn todoBtn">&#9998;</button
									>
									<button
										on:click={() => deleteTodo(todo.id)}
										class="editBtn todoBtn">&#10006;</button
									>
								</div>
							</div>
						</li>
					{/each}
				</ul>
			</div>
		</div>
	</div>
	<div id="footer">
		<h3>Developed By Koosha Paridehpour</h3>
	</div>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}
	#fixItApp {
		border: 2px solid #7ebab5;
		padding: 1em;
		border-radius: 1em;
		width: 75%;
		max-height: 65vh;
	}
	#fixItApp button:hover {
		background-color: #0e4a45;
	}
	#fixItApp button:active {
		background-color: #0a4440;
	}
	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
	.todoLstItmContLeft {
		background-color: #314d59;
		width: 25%;
		border-radius: 0.625em 0 0 0.625em;
		display: flex;
		align-items: center;
		justify-content: center;
	}
	.todoLstItmConRight {
		width: 75%;
		justify-content: center;
		align-items: center;
		overflow: hidden;
	}
	.todoLstItmConRight * {
		display: flex;
		justify-content: flex-start;
		align-items: flex-start;
	}
	.todoLstItmConRightBtns {
		display: flex;
		flex-direction: column;
		width: 20%;
		padding: 0.5em;
	}
	.todoLstItm {
		background-color: #415d69;
		margin: 0;
		border-radius: 0.625em;
		box-shadow: #213f49 ;
		display: flex;
		flex-direction: row;
	}
	.todotxt {
		font-size: 0.25em;
		text-align: left;
	}
	.todoLstItmContLeft,
	.todoLstItmConRight {
		display: flex;
		padding: 0.25em;
	}
	ul {
		list-style-type: none;
	}
	#body {
		align-items: flex-start;
		justify-content: center;
		display: flex;
		width: 100%;
		padding: 0.125em;
	}
	#header,
	#footer {
		background-color: #213f49;
		width: 100%;
		margin: 0;
	}
	.todoBtn {
		background-color: #314d59;
		border-radius: 0.5em;
		color: #f6f5f5;
		width: 100%;
		height: 25%;
		padding: 0.25em;
		display: flex;
		align-items: center;
		justify-content: center;
	}
	.completeBtn {
		width: 2.5em;
		height: 2.5em;
		padding: 0.25em;
		border-radius: 2.5em;
		background-color: #7eafb5;
		font-size: 100%;
	}
	#popupCont {
		background-color: #213f49;
		width: 100%;
		height: 100%;
		display: none;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}
	.popup.open {
		display: block !important;
	}
	#todos {
		padding-inline-start: 0 !important;
	}
	#todos > li {
		margin: 0.5em;
	}
	.addBtn {
		border-radius: 2em;
		background-color: #7eafb5;
		color: #e6e6e6;
	}
</style>
