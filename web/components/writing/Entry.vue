<script setup lang="ts">
import type { Writing } from "@/types";

type Props = {
  writing: Writing;
};

const props = defineProps<Props>();

const author = computed(() => props.writing.author);
</script>

<template>
  <article class="max-w-xl mx-auto py-6 space-y-4 border-b">
    <h2 class="text-2xl font-bold font-poetry">{{ writing.title }}</h2>

    <TiptapContent class="font-poetry" :content="writing.text" />

    <div>
      <div v-if="author" class="space-x-2">
        <UAvatar :alt="`${author.name} ${author?.lastname}`" size="sm" />
        <span class="text-sm">{{ `${author.name} ${author?.lastname}` }}</span>
      </div>
    </div>
  </article>
</template>
