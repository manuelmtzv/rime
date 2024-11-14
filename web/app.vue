<script setup lang="ts">
const { finalizePendingLocaleChange } = useI18n();

const onBeforeEnter = async () => {
  await finalizePendingLocaleChange();
};
</script>

<template>
  <div>
    <NuxtLayout>
      <NuxtPage
        :transition="{
          name: 'page',
          mode: 'out-in',
          onBeforeEnter,
        }"
      />
    </NuxtLayout>

    <UNotifications />
    <AuthRequiredModal />
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
