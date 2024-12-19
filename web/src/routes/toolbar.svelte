<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Collapsible from '$lib/components/ui/collapsible/index.js';
	import * as Tabs from '$lib/components/ui/tabs';
	import { ChevronsUpDown } from 'lucide-svelte';
	import { writable } from 'svelte/store';
	import ToolbarActions from './toolbar-actions.svelte';

	export let apps: App[];
	export let selectedTabId = "";
	export let open = false;

</script>

<Collapsible.Root class="v-[10vh] bg-[#2C2C2C] text-white" {open}>
	<Tabs.Root value={selectedTabId} class="bg-[#2C2C2C]">
		
		<Collapsible.Trigger asChild let:builder class="bg-[#2C2C2C]">
			<Button builders={[builder]} variant="ghost" size="sm" class="w-9 p-0 hover:bg-[#2C2C2C] hover:text-white ">
				<ChevronsUpDown size="16" stroke="3"/>	
			</Button>
		</Collapsible.Trigger>
		
		<Tabs.List class="bg-[#2C2C2C]" >

			{#each apps as app}
				<Tabs.Trigger
					class="data-[state=active]:border-b-2 data-[state=active]:border-[#8D42F7] data-[state=active]:bg-[#2C2C2C] data-[state=active]:text-white"
					value={app.Id}>{app.Name}
				</Tabs.Trigger>
			{/each}

		</Tabs.List>
		
		<Collapsible.Content class="h-40 space-y-2 border-t-2 border-[#737373]">

			{#each apps as app}
				<Tabs.Content value={app.Id}>
					<ToolbarActions {app} />
				</Tabs.Content>
			{/each}

		</Collapsible.Content>
	</Tabs.Root>
</Collapsible.Root>
