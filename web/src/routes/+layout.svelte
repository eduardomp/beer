<script lang="ts">
	import '../app.css';

	let { children } = $props();
</script>

<div class="app">

	<main>
		{@render children()}
	</main>

	<footer>
		<p>
			BEER - BackEnd Engineers Rocks
		</p>
	</footer>
</div>

