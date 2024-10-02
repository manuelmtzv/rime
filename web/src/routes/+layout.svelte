<script>
  import "../assets/css/app.css";
  
	import Header from "$lib/components/base/header.svelte";
  import Footer from "$lib/components/base/footer.svelte";
</script>

<div class="min-h-[100svh] flex flex-col">
  <Header />

  <main class="flex-1">
    <slot />
  </main>

  <Footer />
</div>
