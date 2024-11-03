<script setup lang="ts">
const year = new Date().getFullYear();
</script>

<template>
  <footer class="py-6">
    <div class="container">
      <p class="text-center text-sm text-gray-500">
        &copy; {{ year }} Rime. All rights reserved.
      </p>
    </div>
  </footer>
</template>
