<script>
  import "@fontsource/days-one";
  import logo from "/public/fliphousechill-logo.png";
  import guy from "/public/chillhouse-guy.png";
</script>

<main class="min-h-screen bg-[#e3cba4] flex flex-col items-center px-4 pb-8">
  <header class="flex flex-col items-center mt-8">
    <div class="font-bold text-[1.5rem] mb-2 tracking-wide">
      <span class="text-[#a05a3b]">Flip</span><span class="text-[#444]"
        >House</span
      ><span class="text-[#444]">Chill</span>
    </div>
    <img src={logo} alt="FlipHouseChill Logo" class="w-[70px] mb-2 spin-180-hover" />
  </header>

  <section class="flex items-start justify-center my-8 w-full max-w-[700px]">
    <img src={guy} alt="Chillhouse Guy" class="w-14 h-[71px] mr-4" />
    <div class="flex-1 text-left">
      <h1 class="text-[1.5rem] font-bold text-[#222] leading-tight m-0">
        Chillhouse <span class="text-[#a05a3b] font-bold">flipped</span> the
        <span class="text-[#a05a3b] font-bold">opening weekend</span> for the Emoji
        Movie
      </h1>
    </div>
  </section>

  <section class="flex flex-col gap-4 w-full max-w-[600px]">
    <div
      class="flex items-center rounded-xl px-6 py-4 text-[1.1rem] font-semibold shadow-sm bg-[#f6e7c1] text-[#222] text-left"
    >
      <img src={guy} alt="Chillhouse Guy" class="w-9 h-[46px] mr-3" />
      <span
        >Chillhouse <span class="text-[#a05a3b] font-bold">flipped</span> the
        <span class="text-[#a05a3b] font-bold">opening weekend</span> for the Emoji
        Movie</span
      >
    </div>
    <div
      class="flex items-center rounded-xl px-6 py-4 text-[1.1rem] font-semibold shadow-sm bg-[#ede0c2] text-[#222] text-left"
    >
      <img src={guy} alt="Chillhouse Guy" class="w-9 h-[46px] mr-3" />
      <span
        >Chillhouse <span class="text-[#a05a3b] font-bold">flipped</span> the
        <span class="text-[#a05a3b] font-bold">opening weekend</span> for the Emoji
        Movie</span
      >
    </div>
  </section>
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
