<script lang="ts">
	import '../app.css';

	let { children } = $props();

	export const prerender = false;
	export const ssr = false;
</script>

<div class="app">

	<main>
		{@render children()}
	</main>

</div>

