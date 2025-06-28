<script>
  export let cache = [];
  export let chillhouseCap = null;
  export let formatCap = (val) => val;
  import guyNotFlipped from "/public/chillhouse-guy-not-flipped.png";
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

<div class="w-full h-full min-h-[420px]">
  <swiper-container
    effect="cards"
    grab-cursor="true"
    class="w-[340px] h-[420px] mx-auto"
  >
    {#each cache.slice(0, 7) as entry, i}
      <swiper-slide class="bg-[#f6e7c1] rounded-2xl shadow-md px-8 pt-10 pb-6 flex flex-col justify-end items-start min-h-[340px] relative overflow-hidden" style="background-image: url('/bg-image.png'); background-repeat: none; background-size: 340px 420px;">
        <div class="w-full z-10 text-left">
          <div class="text-left break-words text-3xl font-bold text-[#111] leading-tight mb-2">
            {@html highlightText(entry.item.label || 'Chillhouse flipped the opening weekend for the Emoji Movie')}
          </div>
          <div class="text-[#a05a3b] font-bold text-2xl mt-2">
            {formatCap(entry.item.market_cap)}
          </div>
          {#if chillhouseCap && entry.item.market_cap > chillhouseCap}
            <div class="text-zinc-600 font-medium text-lg mt-2">
              ChillHouse is <span class="highlight">{(entry.item.market_cap / chillhouseCap).toFixed(2)}x</span> away
            </div>
          {/if}
        </div>
        <img src={guyNotFlipped} alt="Chillhouse Guy" class="w-12 absolute left-6 top-5 z-0" />
      </swiper-slide>
    {/each}
  </swiper-container>
</div>

<style>
  .highlight {
    color: #a05a3b;
    font-weight: bold;
  }
</style> 