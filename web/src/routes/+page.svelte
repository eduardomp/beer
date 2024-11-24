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
	} from '@xyflow/svelte';
   
	// 👇 this is important! You need to import the styles for Svelte Flow to work
	import '@xyflow/svelte/dist/style.css';
   
	// We are using writables for the nodes and edges to sync them easily. When a user drags a node for example, Svelte Flow updates its position.
	const initialNodes : Node[] = [
	  {
		id: '1',
		type: 'input',
		data: { label: 'Input Node', body: "This is the input node" },
		style: 'border: 1px solid #8e42f7; padding: 10px;',
		position: { x: 0, y: 0 }
	  },
	  {
		id: '2',
		type: 'default',
		data: { label: 'Node' },
		style: 'border: 1px solid #8e42f7; padding: 10px;',
		position: { x: 0, y: 150 }
	  }
	];
   
	// same for edges
	const initialEdges: Edge[] = [
	  {
		id: '1-2',
		type: 'default',
		source: '1',
		target: '2',
		label: 'HTTP'
	  }
	];
   
	let colorMode: ColorMode = 'dark';

	const nodes = writable<Node[]>(initialNodes);
  	const edges = writable(initialEdges);
  </script>
   
  <!--
  👇 By default, the Svelte Flow container has a height of 100%.
  This means that the parent container needs a height to render the flow.
  -->
  <div style:height="90vh">
	<SvelteFlow
	  {nodes}
	  {edges}
	  {colorMode}
	  fitView
	  on:nodeclick={(event: { detail: { node: any; }; }) => console.log('on node click', event.detail.node)}
	>
	  <Controls />
	  <Background variant={BackgroundVariant.Dots} />
	  <MiniMap />
	</SvelteFlow>
  </div>