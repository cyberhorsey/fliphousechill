<script>
  import "@fontsource/days-one";
  import logo from "/public/fliphousechill-logo.png";
  import guy from "/public/chillhouse-guy.png";
  import guyNotFlipped from "/public/chillhouse-guy-not-flipped.png";
  import { onMount, tick } from 'svelte';
  import { register } from 'swiper/element/bundle';
  import 'swiper/css';
  import 'swiper/css/effect-cards';
  import ChillView from './components/ChillView.svelte';

  let cache = [];
  let loading = true;
  let chillhouseCap = null;
  let etcCap = null;
  let chillhouseToEtcX = null;
  let view = 'list'; // 'list' or 'chill'
  let chillhouseEntry = null;
  let chillAudio;
  let muted = false;
  $: if (chillAudio) chillAudio.muted = muted;

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

   // whenever loading goes false, scroll to our ChillHouse card
 $: if (!loading) {
   // wait for DOM to update
   tick().then(() => {
     const el = document.querySelector('[data-lower="chillhouse"]');
     if (el) {
       el.scrollIntoView({ behavior: 'smooth', block: 'center' });
     }
   });
 }

function formatCap(val) {
  if (val >= 1e9) {
    return (val / 1e9)
      .toFixed(1)
      .replace(/\.0$/, '') + 'b';
  }
  if (val >= 1e6) {
    return (val / 1e6)
      .toFixed(1)
      .replace(/\.0$/, '') + 'm';
  }
  if (val >= 1e3) {
    return (val / 1e3)
      .toFixed(1)
      .replace(/\.0$/, '') + 'k';
  }
  return val.toString();
}

function highlightText(text) {
  // Simple demo: highlight 'flipped' and 'opening weekend' in brown
  return text
    .replace(/(flipped)/gi, '<span class="highlight">$1</span>')
    .replace(/(opening weekend)/gi, '<span class="highlight">$1</span>');
}

onMount(() => {
  register();
});
</script>

{#if view === 'chill'}
  <div class="gradient-blob top-left"></div>
  <div class="gradient-blob top-right"></div>
  <div class="gradient-blob bottom-left"></div>
  <div class="gradient-blob bottom-right"></div>
  <audio src="/chillhouse-beats.mp3" autoplay loop style="display:none" bind:this={chillAudio} {muted} />
  <button
    class="fixed sound-button top-6 right-6 z-20 bg-white border border-[#a05a3b] rounded-xl shadow-md p-2 flex items-center justify-center transition hover:scale-110 outline-none focus:outline-none focus:ring-0 focus:shadow-none"
    on:click={() => muted = !muted}
    aria-label={muted ? 'Unmute music' : 'Mute music'}
    style="width: 54px; height: 54px; outline: none; box-shadow: none;"
  >
    {#if muted}
      <svg xmlns="http://www.w3.org/2000/svg" width="80" height="80" viewBox="0 0 24 24" fill="none" stroke="#a05a3b" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="1" y1="1" x2="23" y2="23"></line><path d="M9 9v3a3 3 0 0 0 5.12 2.12M15 9.34V4a3 3 0 0 0-5.94-.6"></path><path d="M17 16.95A7 7 0 0 1 5 12v-2m14 0v2a7 7 0 0 1-.11 1.23"></path><line x1="12" y1="19" x2="12" y2="23"></line><line x1="8" y1="23" x2="16" y2="23"></line></svg>
    {:else}
      <svg xmlns="http://www.w3.org/2000/svg" width="80" height="80" viewBox="0 0 24 24" fill="none" stroke="#a05a3b" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 1a3 3 0 0 0-3 3v8a3 3 0 0 0 6 0V4a3 3 0 0 0-3-3z"></path><path d="M19 10v2a7 7 0 0 1-14 0v-2"></path><line x1="12" y1="19" x2="12" y2="23"></line><line x1="8" y1="23" x2="16" y2="23"></line></svg>
    {/if}
  </button>
{/if}

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
    <!-- Chill View with Swiper -->
    <section class="flex flex-col items-center justify-center w-full max-w-[700px] my-4">
    
      {#if !loading && cache.length}
        <ChillView {cache} {chillhouseCap} {formatCap} />
      {:else if loading}
        <p class="text-[#444] mt-2">Loading Chill View...</p>
      {:else}
        <p class="text-[#444] mt-2">No entries to show.</p>
      {/if}
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
  .gradient-blob {
    position: fixed;
    width: 350px;
    height: 350px;
    border-radius: 34.4375rem;
    opacity: 0.38;
    filter: blur(46px);
    z-index: 0;
    pointer-events: none;
    animation: blob-move 11s ease-in-out infinite alternate;
    transition: opacity 0.3s;
  }
  .top-left {
    top: -180px;
    left: -180px;
    background: linear-gradient(153deg, #90462F 16.98%, rgba(223, 206, 169, 0.15) 98.14%);
    animation-delay: 0s;
  }
  .top-right {
    top: -180px;
    right: -180px;
    background: linear-gradient(153deg, #61902F 16.98%, rgba(223, 206, 169, 0.15) 98.14%);
    animation-delay: 4s;
  }
  .bottom-left {
    bottom: -180px;
    left: -180px;
    background: linear-gradient(153deg, #2F908E 16.98%, rgba(223, 206, 169, 0.15) 98.14%);
    animation-delay: 8s;
  }
  .bottom-right {
    bottom: -180px;
    right: -180px;
    background: linear-gradient(153deg, #902F51 16.98%, rgba(223, 206, 169, 0.15) 98.14%);
    animation-delay: 12s;
  }
  @keyframes blob-move {
    0% {
      transform: scale(1) translate(0, 0);
    }
    50% {
      transform: scale(1.18) translate(80px, 60px);
    }
    100% {
      transform: scale(1) translate(0, 0);
    }
  }
  .sound-button {
    background-color: #fff;
    outline: none;
    box-shadow: none;
  }

  @media (max-width: 900px) {
    .gradient-blob {
      width: 200px;
      height: 200px;
    }
    .top-left, .top-right, .bottom-left, .bottom-right {
      top: -100px !important;
      bottom: -100px !important;
      left: -100px !important;
      right: -100px !important;
    }
  }
  @media (max-width: 600px) {
    h1 {
      font-size: 1.4rem !important;
    }
    .gradient-blob {
      width: 120px;
      height: 120px;
    }
    .top-left, .top-right, .bottom-left, .bottom-right {
      top: -50px !important;
      bottom: -50px !important;
      left: -50px !important;
      right: -50px !important;
    }
  }
  .swiper-container {
    width: 340px;
    height: 420px;
    margin: 0 auto;
  }
  .chill-card {
    background: #f6e7c1;
    border-radius: 18px;
    box-shadow: 0 2px 8px rgba(160,90,59,0.08);
    padding: 2.5rem 2rem 1.5rem 2rem;
    display: flex;
    flex-direction: column;
    justify-content: flex-end;
    align-items: flex-start;
    min-height: 340px;
    position: relative;
    overflow: hidden;
  }
  .chill-card-content {
    width: 100%;
    z-index: 2;
  }
  .chill-card-text {
    text-align: left;
    word-break: break-word;
  }
  .chill-card-guy {
    width: 48px;
    position: absolute;
    left: 1.5rem;
    top: 1.2rem;
    z-index: 1;
  }
  .highlight {
    color: #a05a3b;
    font-weight: bold;
  }
</style>
