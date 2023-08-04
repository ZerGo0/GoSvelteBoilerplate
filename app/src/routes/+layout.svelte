<script lang="ts">
  import { onMount } from "svelte";
  import "../app.css";

  export let selected = 0;
  let showMenu = false;

  function toggleNavbar() {
    showMenu = !showMenu;
  }

  onMount(() => {
    // TODO: Is there a way to do this?
    const url = window.location.pathname;
    if (url == "/") {
      selected = 0;
    } else if (url == "/test") {
      selected = 1;
    }
  });
</script>

<div class="bg-slate-900">
  <nav class="container p-2 mx-auto md:flex md:justify-between md:items-center">
    <div>
      <a
        class="flex items-center justify-between gap-1 text-xl font-bold md:text-2xl hover:text-red-900"
        href="/"
        on:click={() => (selected = 0)}
      >
        PROJECTNAME
      </a>
    </div>

    <div on:click={toggleNavbar} class="flex md:hidden">
      <button
        type="button"
        class="text-white hover:text-red-800 focus:outline-none focus:text-gray-400"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="w-6 h-6"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
          />
        </svg>
      </button>
    </div>

    <!-- Mobile Menu open: "block", Menu closed: "hidden" -->
    <div
      class="flex-col mt-8 space-y-4 md:flex md:space-y-0 md:flex-row md:items-center md:space-x-10 md:mt-0 {showMenu
        ? 'flex'
        : 'hidden'}"
    >
      <a
        class="hover:text-red-800
          {selected == 0 ? 'text-red-800' : 'text-white'}"
        href="/"
        on:click={() => (selected = 0)}>Home</a
      >
      <a
        class=" hover:text-red-800
          {selected == 1 ? 'text-red-800' : 'text-white'}"
        href="/test"
        on:click={() => (selected = 1)}>Test</a
      >
    </div>
  </nav>
</div>

<div class="container h-full mx-auto">
  <div class="h-full flex flex-col items-center justify-center">
    <slot />
  </div>
</div>

<footer class="bg-slate-900 text-white py-2">
  <div class="container mx-auto text-center">
    <a
      href="https://github.com/zergo0/PROJECTNAME"
      class="hover:text-red-800"
      target="_blank"
      rel="noopener noreferrer"
    >
      <svg
        class="w-6 h-6 inline-block"
        fill="currentColor"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
      >
        <path
          d="M12 .5c-6.63 0-12 5.28-12 11.792 0 5.211 3.438 9.63 8.205 11.188.6.111.82-.254.82-.567 0-.28-.01-1.022-.01-2.01-3.338.711-4.042-1.582-4.042-1.582-.546-1.36-1.335-1.721-1.335-1.721-1.087-.727.084-.711.084-.711 1.205.082 1.838 1.215 1.838 1.215 1.07 1.796 2.809 1.276 3.495.976.108-.758.417-1.276.76-1.567-2.665-.296-5.466-1.31-5.466-5.827 0-1.287.465-2.338 1.235-3.164-.135-.298-.54-1.497.105-3.121 0 0 1.005-.316 3.3 1.209.96-.262 1.98-.392 3-.398 1.02.006 2.04.136 3 .398 2.28-1.525 3.285-1.209 3.285-1.209.645 1.624.24 2.823.12 3.121.765.826 1.23 1.877 1.23 3.164 0 4.53-2.805 5.527-5.475 5.817.42.354.81 1.077.81 2.182 0 1.578-.015 2.846-.015 3.236 0 .314.21.688.825.568C20.565 21.92 24 17.5 24 12.292 24 5.78 18.63.5 12 .5"
        />
      </svg>
    </a>
  </div>
</footer>
