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

<div id="parent">
    <div id="header">
        <div id ="logoCont">
        <i><!--LOGO--></i>
        <h1>Fix It</h1></div>
        <div id="menu">
            <ul>
                <li id ="profileCont" class="headerLstItm">
                    
                </li>
            </ul>
        </div>

    </div>
    <div id="body">
        <div id="content">
            <div id="fixItApp">
                <div id="todoList">
                    <ul id="todos">
                        <!--SCRIPT TO GEN TODO ITEMS-->
                    </ul>
                </div>
            </div>
        </div>
    </div>
    <div id="footer">

    </div>
</div>