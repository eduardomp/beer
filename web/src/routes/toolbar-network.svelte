<script lang="ts">
    import * as Table from "$lib/components/ui/table/index.js";
    import ScrollArea from "$lib/components/ui/scroll-area/scroll-area.svelte";
    import { onMount } from 'svelte';
	import { writable } from "svelte/store";

    let connections = writable<Connection[]>([]);

    onMount(() => {
        //TODO: Externalize the SSE endpoint
	    let evtSource = new EventSource('http://localhost:9999/sse');

        evtSource.addEventListener('connectionsInfo', function (event: MessageEvent) {
            var dataobj = JSON.parse(event.data);
            console.log(dataobj);
            connections.update((items) => {
                items = [...dataobj?.Connections, ...items];
                return items 
            });
        });
    })
   </script>
    
    <ScrollArea class="h-36" >
        <Table.Root>
         <Table.Caption>list of captured connections</Table.Caption>
         <Table.Header>
          <Table.Row>
            <Table.Head class="h-3 text-center">#</Table.Head>
            <Table.Head class="h-3 text-center w-[100px]">Source</Table.Head>
            <Table.Head class="h-3 text-center">Target</Table.Head>
            <Table.Head class="h-3 text-center">Protocol</Table.Head>
            <Table.Head class="h-3 text-center">Status</Table.Head>
            <Table.Head class="h-3 text-center">Duration</Table.Head>
           </Table.Row>
         </Table.Header>
         <Table.Body>
          {#each $connections as connection, i}
           <Table.Row class="even:bg-[#353535]">
            <Table.Cell class="h-3 p-0 border-[#737373] font-bold text-center text-[#21C55D]">{i}</Table.Cell>
            <Table.Cell class="h-3 p-0 border-[#737373] font-bold text-center text-[#21C55D]">{connection.SourceApp.Name}</Table.Cell>
            <Table.Cell class="h-3 p-0 border-[#737373] font-bold text-center text-[#21C55D]">{connection.TargetApp.Name}</Table.Cell>
            <Table.Cell class="h-3 p-0 border-[#737373] font-bold text-center text-[#21C55D]">{connection.Protocol}</Table.Cell>
            <Table.Cell class="h-3 p-0 border-[#737373] font-bold text-center text-[#21C55D]">{connection.Status}</Table.Cell>
            <Table.Cell class="h-3 p-0 border-[#737373] font-bold text-center text-[#21C55D]">{connection.Duration}</Table.Cell>
           </Table.Row>
          {/each}
         </Table.Body>
        </Table.Root>
    </ScrollArea>