<script lang="ts">
	import * as Tabs from '$lib/components/ui/tabs';

    import { LogsIcon, ContainerIcon } from 'lucide-svelte';

    export let app: App;

</script>
<Tabs.Root value="logs" class="bg-[#2C2C2C]">
        <div class="bg-[#2C2C2C] text-white flex flex-row">
            <div class="w-[10vw]">
                <Tabs.List class="bg-[#2C2C2C] flex-col justify-start w-full h-40 pt-0 pr-0" >
                    <Tabs.Trigger
                        class="data-[state=active]:border-r-2 data-[state=active]:border-[#8D42F7] data-[state=active]:bg-[#353535] data-[state=active]:text-white w-full items-start justify-start text-start"
                        value="logs">
                        <div class="grid grid-cols-4">
                            <LogsIcon size="16" stroke="3" color="#22c55e" class="mr-2"/>
                            <span class="text-white col-span-3">Logs</span>
                        </div>
                    </Tabs.Trigger>
                    <Tabs.Trigger
                        class="data-[state=active]:border-r-2 data-[state=active]:border-[#8D42F7] data-[state=active]:bg-[#353535] data-[state=active]:text-white w-full items-start justify-start text-start"
                        value="other">
                        <div class="grid grid-cols-4">
                            <ContainerIcon size="16" stroke="3" color="#3b82f6" class="mr-2"/>
                            <span class="text-white col-span-3">Container</span>
                        </div>
                    </Tabs.Trigger>
                </Tabs.List>
            </div>
            <div class="w-[90vw] bg-[#353535] p-4">
                    <Tabs.Content value="logs">Logs of {app.Name}</Tabs.Content>
                    <Tabs.Content value="other">other of {app.Name}</Tabs.Content>
            </div>
        </div>
</Tabs.Root>