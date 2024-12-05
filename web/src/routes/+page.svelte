<script lang="ts">
	import { writable } from 'svelte/store';
	import {
	  SvelteFlow,
	  Controls,
	  Background,
	  BackgroundVariant,
	  MiniMap,
	  Panel,
	  type ColorMode,
	  type Node,
	  type Edge,
	  type ProOptions,
	} from '@xyflow/svelte';
   
	// 👇 this is important! You need to import the styles for Svelte Flow to work
	import '@xyflow/svelte/dist/style.css';
	import HostInfo from './host-info.svelte';
	import type { PageData } from './$types';
	
	export let data: PageData;
	
	const initialNodes : Node[] = [];
	const initialEdges: Edge[] = [];
	
	let apps: App[] = data.response;

	apps.forEach((app, index) => {
		initialNodes.push({
			id: app.Name,
			type: 'default',
			data: { label: app.Name },
			style: 'border: 1px solid #8e42f7; padding: 10px;',
			position: { x: index * 300, y: 0 }
		});
	});

	let colorMode: ColorMode = 'dark';
	let proOptions: ProOptions = {
		hideAttribution: true,
	}

	const nodes = writable<Node[]>(initialNodes);
  	const edges = writable(initialEdges);
  </script>
   
  <!--
  👇 By default, the Svelte Flow container has a height of 100%.
  This means that the parent container needs a height to render the flow.
  -->
  <div style:height="100vh">
	<SvelteFlow
	  {nodes}
	  {edges}
	  {colorMode}
	  fitView
	  {proOptions}
	  on:nodeclick={(event: { detail: { node: any; }; }) => console.log('on node click', event.detail.node)}
	>
	  <Controls />
	  <Background variant={BackgroundVariant.Dots} />
	  <MiniMap />
	  <Panel position="top-center">
		<HostInfo />
	  </Panel>
	</SvelteFlow>
  </div>