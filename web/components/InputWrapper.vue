<script lang="ts" setup>
type Props = {
  label: string;
  error?: string;
  class?: string;
  classNames?: {
    label?: string;
    error?: string;
  };
};

defineProps<Props>();
</script>

<template>
  <label :class="cn('flex flex-col gap-1', $props.class)">
    <span :class="cn('text-sm text-gray-500', classNames?.label)">
      {{ label }}
    </span>

    <slot></slot>

    <span :class="cn('text-sm text-red-500', classNames?.error)">
      {{ error }}
    </span>
  </label>
</template>
