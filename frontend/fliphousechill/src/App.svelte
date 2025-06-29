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
  let totalSupply = 1_000_000_000; // 1 billion tokens
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

function calculateChillhousePriceFromCap(marketCap) {
  if (!marketCap) return null;
  return marketCap / totalSupply;
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
  <div class="fixed bottom-6 right-6 z-30 flex items-center gap-2">
    <button
      class="sound-button bg-white border border-[#a05a3b] rounded-xl shadow-md p-2 flex items-center justify-center transition hover:scale-110 outline-none focus:outline-none focus:ring-0 focus:shadow-none"
      on:click={() => muted = !muted}
      aria-label={muted ? 'Unmute music' : 'Mute music'}
      style="width: 54px; height: 54px; outline: none; box-shadow: none;"
    >
      {#if muted}
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M17.25 9.75L19.5 12M19.5 12L21.75 14.25M19.5 12L21.75 9.75M19.5 12L17.25 14.25M6.75 8.25L11.47 3.53C11.5749 3.42524 11.7085 3.35392 11.8539 3.32503C11.9993 3.29615 12.15 3.311 12.2869 3.36771C12.4239 3.42442 12.541 3.52045 12.6234 3.64367C12.7058 3.76688 12.7499 3.91176 12.75 4.06V19.94C12.7499 20.0882 12.7058 20.2331 12.6234 20.3563C12.541 20.4795 12.4239 20.5756 12.2869 20.6323C12.15 20.689 11.9993 20.7038 11.8539 20.675C11.7085 20.6461 11.5749 20.5748 11.47 20.47L6.75 15.75H4.51C3.63 15.75 2.806 15.243 2.572 14.396C2.35751 13.6154 2.2492 12.8095 2.25 12C2.25 11.17 2.362 10.367 2.572 9.604C2.806 8.756 3.63 8.25 4.51 8.25H6.75Z" stroke="#954E35" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
      {:else}
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M19.114 5.636C19.9497 6.47173 20.6127 7.46388 21.065 8.55582C21.5173 9.64776 21.7501 10.8181 21.7501 12C21.7501 13.1819 21.5173 14.3522 21.065 15.4442C20.6127 16.5361 19.9497 17.5283 19.114 18.364M16.463 8.288C17.4474 9.27254 18.0004 10.6078 18.0004 12C18.0004 13.3922 17.4474 14.7275 16.463 15.712M6.75 8.25L11.47 3.53C11.5749 3.42524 11.7085 3.35392 11.8539 3.32503C11.9993 3.29615 12.15 3.311 12.2869 3.36771C12.4239 3.42442 12.541 3.52045 12.6234 3.64367C12.7058 3.76688 12.7499 3.91176 12.75 4.06V19.94C12.7499 20.0882 12.7058 20.2331 12.6234 20.3563C12.541 20.4795 12.4239 20.5756 12.2869 20.6323C12.15 20.689 11.9993 20.7038 11.8539 20.675C11.7085 20.6461 11.5749 20.5748 11.47 20.47L6.75 15.75H4.51C3.63 15.75 2.806 15.243 2.572 14.396C2.35751 13.6154 2.2492 12.8095 2.25 12C2.25 11.17 2.362 10.367 2.572 9.604C2.806 8.756 3.63 8.25 4.51 8.25H6.75Z" stroke="#954E35" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
      {/if}
    </button>
    <img src="/chillhousevibe.gif" alt="Chillhouse Vibe" class="w-14 h-14 rounded-xl object-cover" style="margin-left: 0.5rem;" />
  </div>
{/if}

<main class="min-h-screen bg-[#e3cba4] flex flex-col items-center px-2 sm:px-4 pb-8 w-full sm:overflow-x-visible overflow-x-hidden">
  <header class="flex flex-col items-center mt-8 mb-2 w-full max-w-md sm:max-w-2xl">
    <div class="flex flex-row items-center justify-center mb-2 w-full">
      <img src={guyNotFlipped} alt="Chillhouse Guy Not Flipped" class="mr-1" style="max-width: 24px;" />
      <img src={guy} alt="Chillhouse Guy" class="ml-1" style="max-width: 24px;" />
      <span class="font-bold text-[2rem] tracking-wide flex flex-row items-center ml-3">
        <span class="text-[#a05a3b]">Flip</span><span class="text-[#222]">HouseChill</span>
      </span>
    </div>
    {#if chillhouseCap}
      <div class="flex flex-row flex-wrap items-center justify-center gap-2 sm:gap-4 text-sm font-medium text-[#444] mt-1 w-full">
        <span>MC: <span class="text-[#222]">{formatCap(chillhouseCap).toUpperCase()}</span></span>
        <span>Price: <span class="text-[#222]">${chillhouseEntry && chillhouseEntry.item.price ? chillhouseEntry.item.price : (chillhouseCap ? calculateChillhousePriceFromCap(chillhouseCap).toFixed(3) : '0.016')}</span></span>
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
  <div class="flex items-center justify-center my-6 w-full max-w-xs sm:max-w-md">
    <div class="flex rounded-full border border-[#8a4c2b] bg-[#e3cba4] overflow-hidden shadow-sm w-full" style="padding: 0.25rem 0.5rem;">
      <button
        class="flex-1 px-2 sm:px-10 py-3 font-bold text-[1.1rem] sm:text-[1.25rem] transition-colors duration-200 min-w-0"
        style="border-radius: 9999px; outline: none; box-shadow: none; border: none; background: {view === 'list' ? '#c7ad8e' : 'transparent'}; color: #111;"
        on:click={() => view = 'list'}
      >
        List View
      </button>
      <button
        class="flex-1 px-2 sm:px-10 py-3 font-bold text-[1.1rem] sm:text-[1.25rem] transition-colors duration-200 min-w-0"
        style="border-radius: 9999px; outline: none; box-shadow: none; border: none; background: {view === 'chill' ? '#c7ad8e' : 'transparent'}; color: #111;"
        on:click={() => view = 'chill'}
      >
        Chill View
      </button>
    </div>
  </div>

  {#if view === 'list'}
    {#if loading}
      <section class="flex items-start justify-center my-8 w-full max-w-md sm:max-w-[700px]">
        <img src={guy} alt="Chillhouse Guy" class="w-14 h-[71px] mr-4" />
        <div class="flex-1 text-left">
          <h1 class="text-[1.2rem] sm:text-[1.5rem] font-bold text-[#222] leading-tight m-0">
            Loading latest entry...
          </h1>
        </div>
      </section>
      <section class="flex flex-col gap-4 w-full max-w-sm sm:max-w-[600px]">
        <div class="flex items-center rounded-xl px-4 sm:px-6 py-4 text-[1rem] sm:text-[1.1rem] font-semibold shadow-sm bg-[#f6e7c1] text-[#222] text-left">
          <img src={guy} alt="Chillhouse Guy" class="w-9 h-[46px] mr-3" />
          <span>Loading entries...</span>
        </div>
      </section>
    {:else if cache.length}
      <!-- Latest entry as headline -->
      <section class="flex items-start justify-center my-8 w-full max-w-md sm:max-w-[700px]">
        <img src={cache[0].item.icon || guy} alt={cache[0].item.label} class="mr-1" style="max-width: 60px;" />
        <div class="flex-1 text-left">
          <h2 class="text-[1.2rem] sm:text-[1.5rem] font-bold text-[#222] leading-tight m-0">
            {cache[0].item.label} <span class="text-[#a05a3b] font-bold">market cap</span> is <span class="text-[#a05a3b] font-bold">{formatCap(cache[0].item.market_cap)}</span>
          </h2>
          {#if chillhouseToEtcX}
            <div class="mt-2 text-[1rem] sm:text-[1.1rem] text-[#444]">
              Chillhouse needs <span class="font-bold text-[#a05a3b]">{chillhouseToEtcX.toFixed(2)}x</span> to flip Ethereum Classic market cap.
            </div>
          {/if}
        </div>
      </section>

      <!-- The rest as cards -->
      <section class="flex flex-col gap-4 w-full max-w-sm sm:max-w-[600px]">
       {#each cache.slice(1) as entry}
        <div
          class="flex flex-wrap rounded-xl px-4 sm:px-6 py-4 text-[1rem] sm:text-[1.1rem] font-semibold shadow-sm text-[#222] text-left items-start"
          data-lower={entry.item.label.toLowerCase()}
        >
          <img
            src={chillhouseCap && entry.item.market_cap > chillhouseCap
              ? guyNotFlipped
              : (entry.item.icon || guy)}
            alt={entry.item.label}
            class="w-9 h-[46px] mr-3 flex-shrink-0"
          />
          <div class="flex-1 flex flex-col justify-between gap-1">
            <div class="flex flex-row justify-between items-center w-full">
              <span class="whitespace-normal break-words flex-1">
                {entry.item.label}
              </span>
              <span class="font-bold text-[#a05a3b] flex-shrink-0">
                {formatCap(entry.item.market_cap)}
              </span>
            </div>
            {#if chillhouseCap && entry.item.market_cap > chillhouseCap}
              <span class="text-xs text-[#a05a3b] font-bold block">Chillhouse is {(entry.item.market_cap / chillhouseCap).toFixed(2)}x away</span>
            {/if}
          </div>
        </div>
      {/each}

      </section>
    {:else}
      <section class="flex items-start justify-center my-8 w-full max-w-md sm:max-w-[700px]">
        <img src={guy} alt="Chillhouse Guy" class="w-14 h-[71px] mr-4" />
        <div class="flex-1 text-left">
          <h1 class="text-[1.2rem] sm:text-[1.5rem] font-bold text-[#222] leading-tight m-0">
            No entries found.
          </h1>
        </div>
      </section>
    {/if}
  {:else}
    <!-- Chill View with Swiper -->
    <section class="flex flex-col items-center justify-center w-full max-w-xs sm:max-w-[700px] my-4">
    
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
