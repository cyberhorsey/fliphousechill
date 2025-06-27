<script>
  import svelteLogo from './assets/svelte.svg'
  import viteLogo from '/vite.svg'

  // fetch https://chillhouse-api-zaydoupxua-uw.a.run.app/cache and store it
  // in a variable called cache
  let cache = [];
  async function fetchCache() {
    try {
      const response = await fetch('https://chillhouse-api-zaydoupxua-uw.a.run.app/cache', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      })
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }
      const data = await response.json()
      cache = data
      console.log('Cache fetched successfully:', cache)
    } catch (error) {
      console.error('Error fetching cache:', error)
    }
  }

  fetchCache()
  
   // format a float as "23.5m" or "1.2b"
  function formatCap(val) {
    if (val >= 1e9) {
      return (val / 1e9)
        .toFixed(1)
        .replace(/\.0$/, '') + 'b';
    }
    return (val / 1e6)
      .toFixed(1)
      .replace(/\.0$/, '') + 'm';
  }

</script>

<main>
  <div>
    <a href="https://vite.dev" target="_blank" rel="noreferrer">
      <img src={viteLogo} class="logo" alt="Vite Logo" />
    </a>
    <a href="https://svelte.dev" target="_blank" rel="noreferrer">
      <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
    </a>
  </div>
  <h1>Vite + Svelte</h1>


  <div class="card">
    {#if cache.length}
      <h2>Cache Data</h2>
      {#each cache as entry (entry.item.label)}
        <div class="cache-entry">
          {#if entry.item.icon}
            <img src={entry.item.icon} alt={entry.item.label} width="24" height="24" />
          {/if}
          {#if entry.item.label.toLowerCase() == "chillhouse"}
          <strong>{entry.item.label}:</strong>
          {:else}
          <span>{entry.item.label}:</span>
          {/if}
          <span>{formatCap(entry.item.market_cap)}</span>
        </div>
      {/each}
    {:else}
      <p>Loading cache...</p>
    {/if}
  </div>

  <p>
    Check out <a href="https://github.com/sveltejs/kit#readme" target="_blank" rel="noreferrer">SvelteKit</a>, the official Svelte app framework powered by Vite!
  </p>

  <p class="read-the-docs">
    Click on the Vite and Svelte logos to learn more
  </p>
</main>

<style>
  .logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 300ms;
  }
  .logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
  }
  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
  }
  .read-the-docs {
    color: #888;
  }
</style>
