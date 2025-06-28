<script>
  export let cache = [];
  export let chillhouseCap = null;
  export let formatCap = (val) => val;
  import guyNotFlipped from "/public/chillhouse-guy-not-flipped.png";
  import guy from "/public/chillhouse-guy.png";
  import { onMount } from 'svelte';
  import { register } from 'swiper/element/bundle';
  import 'swiper/css';
  import 'swiper/css/effect-cards';

  function highlightText(text) {
    return text
      .replace(/(flipped)/gi, '<span class="highlight">$1</span>')
      .replace(/(opening weekend)/gi, '<span class="highlight">$1</span>');
  }

  onMount(() => {
    register();
  });
</script>

<div class="w-full h-full min-h-[320px]">
  <swiper-container
    effect="cards"
    grab-cursor="true"
    class="w-full max-w-xs sm:w-[340px] sm:h-[420px] mx-auto"
    style="height: 360px;"
  >
    {#if chillhouseCap}
      {#each cache.filter(entry => entry.item.market_cap > chillhouseCap) as entry, i}
        <swiper-slide class="bg-[#f6e7c1] rounded-2xl shadow-md px-4 sm:px-8 pt-6 sm:pt-10 pb-4 sm:pb-6 flex flex-col justify-end items-start min-h-[240px] sm:min-h-[340px] relative overflow-hidden"
          style="background-image: url('/bg-image.png'); background-repeat: no-repeat; background-size: cover;">
          <div class="w-full z-10 text-left">
            <div class="text-left break-words text-xl sm:text-3xl font-bold text-[#111] leading-tight mb-2">
              {@html highlightText(entry.item.label || 'Chillhouse flipped the opening weekend for the Emoji Movie')}
            </div>
            <div class="text-[#a05a3b] font-bold text-lg sm:text-2xl mt-2">
              {formatCap(entry.item.market_cap)}
            </div>
            <div class="text-zinc-600 font-medium text-base sm:text-lg mt-2">
              ChillHouse is <span class="highlight">{(entry.item.market_cap / chillhouseCap).toFixed(2)}x</span> away
            </div>
          </div>
          <img src={guyNotFlipped} alt="Chillhouse Guy" class="w-10 sm:w-12 absolute left-4 sm:left-6 top-4 sm:top-5 z-0" />
        </swiper-slide>
      {/each}
      {#each cache.filter(entry => entry.item.market_cap <= chillhouseCap) as entry, i}
        <swiper-slide class="bg-[#f6e7c1] rounded-2xl shadow-md px-4 sm:px-8 pt-6 sm:pt-10 pb-4 sm:pb-6 flex flex-col justify-end items-start min-h-[240px] sm:min-h-[340px] relative overflow-hidden"
          style="background-image: url('/bg-image.png'); background-repeat: no-repeat; background-size: cover;">
          <div class="w-full z-10 text-left">
            <div class="text-left break-words text-xl sm:text-3xl font-bold text-[#111] leading-tight mb-2">
              {@html highlightText(entry.item.label || 'Chillhouse flipped the opening weekend for the Emoji Movie')}
            </div>
            <div class="text-[#a05a3b] font-bold text-lg sm:text-2xl mt-2">
              {formatCap(entry.item.market_cap)}
            </div>
            <div class="text-green-700 font-medium text-base sm:text-lg mt-2">
              Chillhouse Flipped
            </div>
          </div>
          <img src={guy} alt="Chillhouse Guy Flipped" class="w-10 sm:w-12 absolute left-4 sm:left-6 top-4 sm:top-5 z-0" />
        </swiper-slide>
      {/each}
    {:else}
      {#each cache.slice(0, 7) as entry, i}
        <swiper-slide class="bg-[#f6e7c1] rounded-2xl shadow-md px-4 sm:px-8 pt-6 sm:pt-10 pb-4 sm:pb-6 flex flex-col justify-end items-start min-h-[240px] sm:min-h-[340px] relative overflow-hidden"
          style="background-image: url('/bg-image.png'); background-repeat: no-repeat; background-size: cover;">
          <div class="w-full z-10 text-left">
            <div class="text-left break-words text-xl sm:text-3xl font-bold text-[#111] leading-tight mb-2">
              {@html highlightText(entry.item.label || 'Chillhouse flipped the opening weekend for the Emoji Movie')}
            </div>
            <div class="text-[#a05a3b] font-bold text-lg sm:text-2xl mt-2">
              {formatCap(entry.item.market_cap)}
            </div>
          </div>
          <img src={guyNotFlipped} alt="Chillhouse Guy" class="w-10 sm:w-12 absolute left-4 sm:left-6 top-4 sm:top-5 z-0" />
        </swiper-slide>
      {/each}
    {/if}
  </swiper-container>
</div>

<style>
  .highlight {
    color: #a05a3b;
    font-weight: bold;
  }
</style> 