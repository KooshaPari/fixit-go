<script>
	import { onMount } from "svelte";
	import Modal from "./Modal.svelte";
	let todos = [];

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
	}
	async function editTodo(title, description, dueDate) {
		const newTodo = {title, descripton, dueDate}
		try{
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
		<div id="content">
			<div id="fixItApp">
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
						<button on:click={() => createTodo(title, description, dueDate)}
							>Create</button
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
						{#each todos as todo}
							<li class="todoLstItm">
								<div class="todoLstItmContLeft">
									<button
										on:click={() => deleteTodo(todo.id)}
										class="completeBtn todoBtn">Done</button
									>
								</div>
								<div class="todoLstItmConRight">
									<h3>{todo.title}</h3>
									<div style="display:flex; flex-direction:column; width: 30%">
										<p class="todotxt">{todo.description}</p>
										<!-- due Date below-->
										<p class="todotxt">{todo.dueDate}</p>
									</div>
									<div class="todoLstItmConRightBtns">
										<button
											on:click={() => editTodo(todo.id)}
											class="deleteBtn todoBtn">Edit</button
										>
										<button
											on:click={() => deleteTodo(todo.id)}
											class="editBtn todoBtn">X</button
										>
									</div>
								</div>
							</li>
						{/each}
					</ul>
				</div>
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
	.todoLstItm {
		background-color: #415d69;
		margin: 0;
		border-radius: 0.625em;
		display: flex;
		flex-direction: row;
	}
	.todotxt {
		font-size: 0.5em;
		text-align: left;
	}
	.todoLstItmContLeft,
	.todoLstItmConRight {
		display: flex;
	}
	ul {
		list-style-type: none;
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
		width: 1.5em;
		height: 0.5em;
	}
	.completeBtn {
		width: 2.5em;
		height: 2.5em;
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
</style>
