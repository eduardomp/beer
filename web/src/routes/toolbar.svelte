<script lang="ts">
	import * as Collapsible from '$lib/components/ui/collapsible/index.js';
	import * as Tabs from '$lib/components/ui/tabs';
</script>


// TODO: Turn the tabs gray and not rounded
// TODO: When click on any tab, open the content, if click again in another tab, just change the content without close
// TODO: Improve the border of the host-info component
// TODO: Receive props from the parent component
// TODO: List the tabs related to the apps
// TODO: Shows content related to the apps
// TODO: When click on a node, selects the correspondent tab on the toolbar opening the content

<Collapsible.Root class="v-[10vh] bg-[#2C2C2C] text-white">
	<Tabs.Root value="account" class="w-[400px]">
        <Collapsible.Trigger>
		    <Tabs.List>
				<Tabs.Trigger value="account">Account</Tabs.Trigger>
				<Tabs.Trigger value="password">Password</Tabs.Trigger>
            </Tabs.List>
        </Collapsible.Trigger>

		<Collapsible.Content class="space-y-2 h-40">
			<Tabs.Content value="account">Make changes to your account here.</Tabs.Content>
			<Tabs.Content value="password">Change your password here.</Tabs.Content>
		</Collapsible.Content>
	</Tabs.Root>
</Collapsible.Root>
