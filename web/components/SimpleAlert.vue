<script setup lang="ts">
type Props = {
    title?: string;
    message?: string;
    class?: string
}

defineProps<Props>();
</script>

<template>
    <div :class="cn('flex flex-col gap-2 border-l-4 pl-4 dark:border-gray-300', $props.class)">
        <h1 v-if="title">{{ title }}</h1>
        <p v-if="message" class="text-sm text-gray-500 dark:text-gray-400"> {{ message }}</p>
    </div>
</template>