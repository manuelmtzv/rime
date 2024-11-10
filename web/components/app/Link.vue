<script setup lang="ts">
type Props = {
  to: string;
  class?: string;
  mobile?: boolean;
};

defineProps<Props>();
</script>

<template>
  <NuxtLink
    :class="cn(mobile ? 'link-mobile' : 'link', $props.class)"
    :to="to"
    active-class="!opacity-100"
  >
    <slot />
  </NuxtLink>
</template>

<style scoped>
.link {
  @apply px-2 py-0.5 opacity-75 hover:opacity-100 transition-all duration-300 ease-in-out;
}

.link-mobile {
  @apply px-2 py-2 opacity-75 hover:opacity-100 border rounded-md transition-all duration-300 ease-in-out block;
}
</style>
