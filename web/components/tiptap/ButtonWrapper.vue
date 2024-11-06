<script setup lang="ts">
type TiptapButtonWrapperProps = {
  isActive: boolean;
};

defineProps<TiptapButtonWrapperProps>();
</script>

<template>
  <button :class="cn(
    'p-1 flex text-gray-600 border border-transparent dark:text-gray-400 rounded-md transition-all hover:border-gray-100 h-fit',
    isActive &&
    'text-gray-900 bg-gray-50 dark:bg-gray-600 border-gray-200 dark:text-gray-200 dark:border-gray-500'
  )
    ">
    <slot />
  </button>
</template>
