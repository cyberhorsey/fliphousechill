<script>
  import "@fontsource/days-one";
  import logo from "/public/fliphousechill-logo.png";
  import guy from "/public/chillhouse-guy.png";
  import guyNotFlipped from "/public/chillhouse-guy-not-flipped.png";

  let cache = [];
  let loading = true;
  let chillhouseCap = null;
  let etcCap = null;
  let chillhouseToEtcX = null;
  let view = 'list'; // 'list' or 'chill'
  let chillhouseEntry = null;

  async function fetchCache() {
    try {
      const response = await fetch('https://chillhouse-api-zaydoupxua-uw.a.run.app/cache', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      cache = data;
      // Find Chillhouse and Ethereum Classic entries
      chillhouseEntry = cache.find(e => e.item.label && e.item.label.toLowerCase().includes('chillhouse'));
      const etcEntry = cache.find(e => e.item.label && e.item.label.toLowerCase().includes('ethereum classic'));
      chillhouseCap = chillhouseEntry ? chillhouseEntry.item.market_cap : null;
      etcCap = etcEntry ? etcEntry.item.market_cap : null;
      chillhouseToEtcX = (chillhouseCap && etcCap) ? (etcCap / chillhouseCap) : null;
    } catch (error) {
      console.error('Error fetching cache:', error);
    } finally {
      loading = false;
    }
  }

  fetchCache();

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

<main class="min-h-screen bg-[#e3cba4] flex flex-col items-center px-4 pb-8">
  <header class="flex flex-col items-center mt-8 mb-2">
    <div class="flex flex-row items-center justify-center mb-2">
      <img src={guyNotFlipped} alt="Chillhouse Guy Not Flipped" class="mr-1" style="max-width: 24px;" />
      <img src={guy} alt="Chillhouse Guy" class="ml-1" style="max-width: 24px;" />
      <span class="font-bold text-[2rem] tracking-wide flex flex-row items-center ml-3">
        <span class="text-[#a05a3b]">Flip</span><span class="text-[#222]">HouseChill</span>
      </span>
    </div>
    {#if chillhouseCap}
      <div class="flex flex-row items-center justify-center gap-4 text-sm font-medium text-[#444] mt-1">
        <span>MC: <span class="text-[#222]">{formatCap(chillhouseCap).toUpperCase()}</span></span>
        <span>Price: <span class="text-[#222]">${chillhouseEntry && chillhouseEntry.item.price ? chillhouseEntry.item.price : '0.016'}</span></span>
        <a
          href="https://dexscreener.com/solana/35tqqmeirwebk6fr5qipwastuaavo32vjnuljpxvsxuk"
          target="_blank"
          rel="noopener noreferrer"
          class="underline decoration-underline font-bold"
          style="color: #222; text-decoration: underline;"
        >
          Dexscreener
        </a>
      </div>
    {/if}
  </header>

  <!-- Tab Toggle UI -->
  <div class="flex items-center justify-center my-6">
    <div class="flex rounded-full border border-[#8a4c2b] bg-[#e3cba4] overflow-hidden shadow-sm" style="padding: 0.25rem 0.5rem;">
      <button
        class="px-10 py-3 font-bold text-[1.25rem] transition-colors duration-200"
        style="border-radius: 9999px; outline: none; box-shadow: none; border: none; background: {view === 'list' ? '#c7ad8e' : 'transparent'}; color: #111;"
        on:click={() => view = 'list'}
      >
        List View
      </button>
      <button
        class="px-10 py-3 font-bold text-[1.25rem] transition-colors duration-200"
        style="border-radius: 9999px; outline: none; box-shadow: none; border: none; background: {view === 'chill' ? '#c7ad8e' : 'transparent'}; color: #111;"
        on:click={() => view = 'chill'}
      >
        Chill View
      </button>
    </div>
  </div>

  {#if view === 'list'}
    {#if loading}
      <section class="flex items-start justify-center my-8 w-full max-w-[700px]">
        <img src={guy} alt="Chillhouse Guy" class="w-14 h-[71px] mr-4" />
        <div class="flex-1 text-left">
          <h1 class="text-[1.5rem] font-bold text-[#222] leading-tight m-0">
            Loading latest entry...
          </h1>
        </div>
      </section>
      <section class="flex flex-col gap-4 w-full max-w-[600px]">
        <div class="flex items-center rounded-xl px-6 py-4 text-[1.1rem] font-semibold shadow-sm bg-[#f6e7c1] text-[#222] text-left">
          <img src={guy} alt="Chillhouse Guy" class="w-9 h-[46px] mr-3" />
          <span>Loading entries...</span>
        </div>
      </section>
    {:else if cache.length}
      <!-- Latest entry as headline -->
      <section class="flex items-start justify-center my-8 w-full max-w-[700px]">
        <img src={cache[0].item.icon || guy} alt={cache[0].item.label} class="mr-1" style="max-width: 60px;" />
        <div class="flex-1 text-left">
          <h2 class="text-[1.5rem] font-bold text-[#222] leading-tight m-0">
            {cache[0].item.label} <span class="text-[#a05a3b] font-bold">market cap</span> is <span class="text-[#a05a3b] font-bold">{formatCap(cache[0].item.market_cap)}</span>
          </h2>
          {#if chillhouseToEtcX}
            <div class="mt-2 text-[1.1rem] text-[#444]">
              Chillhouse needs <span class="font-bold text-[#a05a3b]">{chillhouseToEtcX.toFixed(2)}x</span> to flip Ethereum Classic market cap.
            </div>
          {/if}
        </div>
      </section>

      <!-- The rest as cards -->
      <section class="flex flex-col gap-4 w-full max-w-[600px]">
        {#each cache.slice(1) as entry}
          <div
            class="flex rounded-xl px-6 py-4 text-[1.1rem] font-semibold shadow-sm text-[#222] text-left items-center"
            class:bg-[#decca7]={entry.item.label === 'Chillhouse'}
            class:border={entry.item.label === 'Chillhouse'}
            class:border-[#a05a3b]={entry.item.label === 'Chillhouse'}
            class:font-bold={entry.item.label === 'Chillhouse'}
            class:bg-[#f6e7c1]={entry.item.label !== 'Chillhouse'}
          >
            <img src={chillhouseCap && entry.item.market_cap > chillhouseCap ? guyNotFlipped : (entry.item.icon || guy)} alt={entry.item.label} class="w-9 h-[46px] mr-3 flex-shrink-0 self-center" />
            <div class="flex flex-col justify-center">
              <span>
                {entry.item.label}: <span class="text-[#a05a3b] font-bold">{formatCap(entry.item.market_cap)}</span>
              </span>
              {#if chillhouseCap && entry.item.market_cap > chillhouseCap}
                <div class="mt-1 text-[1rem] font-normal">
                  <span class="text-[#444]">ChillHouse is</span> <span class="text-[#a05a3b] font-bold">{(entry.item.market_cap / chillhouseCap).toFixed(2)}x</span> <span class="text-[#444]">away</span>
                </div>
              {/if}
            </div>
          </div>
        {/each}
      </section>
    {:else}
      <section class="flex items-start justify-center my-8 w-full max-w-[700px]">
        <img src={guy} alt="Chillhouse Guy" class="w-14 h-[71px] mr-4" />
        <div class="flex-1 text-left">
          <h1 class="text-[1.5rem] font-bold text-[#222] leading-tight m-0">
            No entries found.
          </h1>
        </div>
      </section>
    {/if}
  {:else}
    <!-- Chill View Placeholder -->
    <section class="flex flex-col items-center justify-center w-full max-w-[700px] my-8">
      <h2 class="text-[1.5rem] font-bold text-[#a05a3b]">Chill View</h2>
      <p class="text-[#444] mt-2">This is the Chill View. Add your custom content here!</p>
    </section>
  {/if}
</main>

<style>
  :global(.spin-180-hover:hover) {
    transition: transform 0.5s cubic-bezier(0.4,0,0.2,1);
    transform: rotate(180deg);
  }
  :global(.spin-180-hover) {
    transition: transform 0.5s cubic-bezier(0.4,0,0.2,1);
  }
  :global(body) {
    background-color: #e3cba4;
  }
  @media (max-width: 600px) {
    h1 {
      font-size: 1.4rem !important;
    }
  }
</style>
