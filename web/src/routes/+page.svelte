<script lang="ts">
	import { writable } from 'svelte/store';
	import {
		SvelteFlow,
		Controls,
		Background,
		BackgroundVariant,
		MiniMap,
		type ColorMode,
		type Node,
		type Edge,
		type ProOptions
	} from '@xyflow/svelte';

	// 👇 this is important! You need to import the styles for Svelte Flow to work
	import '@xyflow/svelte/dist/style.css';
	import HostInfo from './host-info.svelte';
	import Toolbar from './toolbar.svelte';
	import type { PageData } from './$types';

	export let data: PageData;

	const initialNodes: Node[] = [];
	const initialEdges: Edge[] = [];

	let apps: App[] = data.response;

	apps.forEach((app, index) => {
		initialNodes.push({
			id: app.Id,
			type: 'default',
			data: { label: app.Name },
			style: 'border: 1px solid #8e42f7; padding: 10px;',
			position: { x: index * 300, y: 0 }
		});
	});

	let colorMode: ColorMode = 'dark';
	let proOptions: ProOptions = {
		hideAttribution: true
	};

	const nodes = writable<Node[]>(initialNodes);
	const edges = writable(initialEdges);

	let selectedTabId = writable<string>('');
	let isToolbarOpen = writable<boolean>(false);
</script>

<div style:height="100vh" class="flex flex-col">
	<SvelteFlow
		{nodes}
		{edges}
		{colorMode}
		fitView
		{proOptions}
		on:nodeclick={(event: { detail: { node: any } }) => {
			selectedTabId.set(event.detail.node.id);
			isToolbarOpen.set(true);
		}}
	>
		<Controls />
		<Background variant={BackgroundVariant.Dots} />
		<MiniMap />
	</SvelteFlow>
	<Toolbar {apps} selectedTabId={$selectedTabId} open={$isToolbarOpen} />
	<HostInfo />
</div>
