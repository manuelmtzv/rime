<script setup lang="ts">
type Props = {
  class?: string;
};

defineProps<Props>();
</script>

<template>
  <div :class="cn('border p-6 rounded-md shadow-sm', $props.class)">
    <slot />
  </div>
</template>
