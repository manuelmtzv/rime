<script setup lang="ts"></script>

<template>
  <header
    class="pt-7 pb-5 shadow-sm border-b border-transparent dark:border-gray-50"
  >
    <div class="container flex items-center justify-between">
      <NuxtLink to="/">
        <AppLogo />
      </NuxtLink>

      <div class="flex gap-4 items-center min-w-[16.477rem] min-h-8">
        <AppNavigation />
        <AppThemeToggle />
      </div>
    </div>
  </header>
</template>
