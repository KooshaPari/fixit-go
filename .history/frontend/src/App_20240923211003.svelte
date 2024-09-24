<script>
	export let name;
	import { onMount } from 'svelte';
	let todo = {
		title: 'Fix This App',
		description: 'Style the Description',
		dueDate: 'Tonight'
	}
	let todos = [todo];

	/*onMount(async () => {
		const response = await fetch('/todos');
		todos = await response.json();
	});

	async function editTodo() {
		/*TODO*/
	}
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

<main>
<!---->


    <div id="header">
        <div id ="logoCont">
        <h1>Fix It</h1></div>
       

    <div id="body">
        <div id="content">
            <div id="fixItApp">
                <div id="todoList">
					<button class="addBtn">Add</button>
                    <ul id="todos">
                       
						{#each todos as todo}
						<li class="todoLstItm">
					
								<div class="todoLstItmContLeft">
									<button class="completeBtn todoBtn">Done</button>
								</div>
								<div class="todoLstItmConRight">
									<h3>{todo.title}</h3>
									<div style="display:flex; flex-direction:column; width: 30%">
									<p class="todotxt">{todo.description}</p>
									<!-- due Date below-->
									<p class="todotxt">{todo.dueDate}</p></div>
									<div class="todoLstItmConRightBtns">
										<button class="deleteBtn todoBtn">Edit</button>
										<button class="editBtn todoBtn">X</button>
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
	.todoLstItmContLeft{
		background-color: #314D59;
		width: 25%;
		border-radius:0.625em 0 0 0.625em;
		display: flex;
		align-items: center;
		justify-content: center;
	}
	.todoLstItmConRight{
		width: 75%;
		justify-content: center;
		align-items: center;
		overflow:hidden;
	}
	.todoLstItm{
		background-color: #415D69;
		margin: 0;
		border-radius: 0.625em;
		display: flex;
		flex-direction: row;


	}
	.todotxt{
		font-size: 0.5em;
		text-align: left;
	}
	.todoLstItmContLeft, .todoLstItmConRight{
		display: flex;
	}
	ul {
    list-style-type: none;
  }

	#header, #footer{
		background-color: #213f49;
		width: 100%;
		margin:0;
	}
	.todoBtn{
		background-color: #314D59;
		border-radius: 0.5em;
		color: #f6f5f5;
		width: 1.5em;
		height: 0.5em;
	}
	.completeBtn{
		width: 2.5em;
		height: 2.5em;
		border-radius:2.5em;
		background-color: #7eafb5;
		font-size: 100%;
	}


	
</style>
