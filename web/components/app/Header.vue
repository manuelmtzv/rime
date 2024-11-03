<script setup lang="ts"></script>

<template>
  <header class="pt-7 pb-5">
    <div class="container flex items-center justify-between">
      <NuxtLink to="/">
        <AppLogo />
      </NuxtLink>

      <div class="flex gap-4 items-center">
        <AppNavigation />
        <AppThemeToggle />
      </div>
    </div>
  </header>
</template>
