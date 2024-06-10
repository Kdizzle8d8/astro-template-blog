<script lang="ts">
  import { Button } from "$lib/components/ui/button/index.js";

  type routeObj = {
    name: string;
    path: string;
  };
  const routes: routeObj[] = [
    { name: "Home", path: "/" },
    { name: "Posts", path: "/posts" },
    { name: "Drafts", path: "/draft" },
  ];
</script>

<div class="w-full flex items-center h-[3rem] border-b-2">
  <span class="flex items-end">
    <h3 class=" font-bold ml-2 self-center ">Template's Blog</h3>
    <span class="inline-flex  ml-3">
      {#each routes as route}
        <Button variant="link" href={route.path}  >{route.name}</Button>
      {/each}
    </span>
  </span>
</div>
