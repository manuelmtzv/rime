<script setup lang="ts">
type Props = {
  contained?: boolean;
  class?: string;
};

withDefaults(defineProps<Props>(), {
  contained: true,
});
</script>

<template>
  <div :class="cn('px-4 py-4', contained && 'container', $props.class)">
    <slot />
  </div>
</template>
