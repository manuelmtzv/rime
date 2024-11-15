<script setup lang="ts">
const year = new Date().getFullYear();
</script>

<template>
  <footer class="py-6">
    <div class="container flex items-center gap-4 justify-center  text-sm text-gray-500">
      <p class="text-center">
        &copy; {{ year }} Rime. {{ $t("app.footer.rights") }}
      </p>

      <span>-</span>

      <a href="https://www.paypal.com/donate/?hosted_button_id=WJ3U5GF2M25XL" target="_blank"
        class=" px-2 py-0.5 border dark:border-gray-400 rounded-md inline-flex items-center gap-1 hover:border-gray-300 hover:dark:border-gray-50 transition-colors duration-300">
        <Icon name="heroicons:heart" class="text-red-500" size="16" />
        Donar
      </a>
    </div>
  </footer>
</template>
