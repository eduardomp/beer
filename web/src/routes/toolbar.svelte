<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Collapsible from '$lib/components/ui/collapsible/index.js';
	import * as Tabs from '$lib/components/ui/tabs';
	import { ChevronsUpDown } from 'lucide-svelte';

	export let apps: App[];

	// TODO: When click on a node, selects the correspondent tab on the toolbar opening the content
</script>

<Collapsible.Root class="v-[10vh] bg-[#2C2C2C] text-white">
	<Tabs.Root value="account" class="w-[400px]">
		
		<Collapsible.Trigger asChild let:builder class="bg-[#2C2C2C]">
			<Button builders={[builder]} variant="ghost" size="sm" class="w-9 p-0 hover:bg-[#2C2C2C] hover:text-white ">
				<ChevronsUpDown size="16" stroke="3"/>	
			</Button>
		</Collapsible.Trigger>
		
		<Tabs.List class="bg-[#2C2C2C]">

			{#each apps as app}
				<Tabs.Trigger
					class="data-[state=active]:border-b-2 data-[state=active]:border-[#8D42F7] data-[state=active]:bg-[#2C2C2C] data-[state=active]:text-white"
					value={app.Name}>{app.Name}</Tabs.Trigger
				>
			{/each}

		</Tabs.List>

		<Collapsible.Content class="h-40 space-y-2">

			{#each apps as app}
				<Tabs.Content value={app.Name}>{app.Name}</Tabs.Content>
			{/each}

		</Collapsible.Content>
	</Tabs.Root>
</Collapsible.Root>
