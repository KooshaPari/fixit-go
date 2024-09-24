<script>
	export let name;
	import { onMount } from 'svelte';
	let todo = {
		title: 'test',
		description: 'test2',
		dueDate: 'test3'
	}
	let todos = [todo];

	/*onMount(async () => {
		const response = await fetch('/todos');
		todos = await response.json();
	});*/

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

<div id="parent">
    <div id="header">
        <div id ="logoCont">
        <i><!--LOGO--></i>
        <h1>Fix It</h1></div>
        <div id="profileCont">
            <i><!--Profile Img--></i>
            <h2><!--Your Name--></h2></div>

        </div>

    <div id="body">
        <div id="content">
            <div id="fixItApp">
                <div id="todoList">
					<button class="addBtn">Add</button>
                    <ul id="todos">
                        <!--SCRIPT TO GEN TODO ITEMS-->
						{#each todos as todo}
						<li class="todoLstItm">
							<div class="todoLstItmCont">
								<div class="todoLstItmContLeft">
									<button class="completeBtn"></button>
								</div>
								<div class="todoLstItmContRight">
									<h3>{todo.title}</h3>
									<p>{todo.description}</p>
									<!-- due Date below-->
									<p>{todo.dueDate}</p>
									<div class="todoLstItmContRightBtns">
										<button class="deleteBtn"></button>
										<button class="editBtn"></button>
									</div>
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
	.todoLstItm{
		background-color: #ff3e00;

	}
	#parent{
		width:100%;

		display: block;
		background-color: #213F49
	}
</style>
