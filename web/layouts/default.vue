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
