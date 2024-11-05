<script setup lang="ts"></script>

<template>
  <div class="flex flex-col min-h-[100svh]" id="app">
    <AppHeader />

    <main class="flex-1">
      <slot />
    </main>

    <AppFooter />
  </div>
</template>

<style>
.page-enter-active,
.page-leave-active {
  transition: all 350ms;
}
.page-enter-from,
.page-leave-to {
  opacity: 0;
  filter: blur(2px);
}
</style>
