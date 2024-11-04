<script setup lang="ts">
type Props = {
    to: string;
    class?: string;
}

defineProps<Props>();
</script>

<template>
    <NuxtLink class="px-2 py-0.5 opacity-70 hover:opacity-100 transition-all duration-300 ease-in-out" :to="to"
        active-class="opacity-100">
        <slot />
    </NuxtLink>
</template>