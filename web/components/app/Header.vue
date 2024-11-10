<script setup lang="ts">
const localePath = useLocalePath();
</script>

<template>
  <header
    class="pt-7 pb-5 shadow-sm shadow-gray-200 border-b border-transparent dark:border-gray-50 dark:shadow-none"
  >
    <div class="container flex items-center justify-between">
      <NuxtLink :to="localePath('/')">
        <AppLogo />
      </NuxtLink>

      <div class="hidden md:flex gap-4 items-center min-h-8">
        <AppNavigation />
        <AppThemeToggle />
      </div>

      <AppMobileNavigation class="block md:hidden" />
    </div>
  </header>
</template>
