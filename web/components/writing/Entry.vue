<script setup lang="ts">
import type { Writing } from "@/types";

type Props = {
  writing: Writing;
};

const props = defineProps<Props>();

const author = computed(() => props.writing.author);
</script>

<template>
  <article class="max-w-xl mx-auto py-6 flex flex-col gap-4 border-b">
    <NuxtLink :to="`/writings/${writing.id}`">
      <h2 class="text-2xl font-bold font-poetry mb-4">{{ writing.title }}</h2>

      <TiptapContent class="font-poetry" :content="writing.text" />
    </NuxtLink>

    <div class="flex justify-between items-center gap-4 mt-2">
      <div v-if="author" class="space-x-2">
        <UAvatar :alt="`${author.name} ${author?.lastname}`" size="sm" />
        <span class="text-sm">{{ `${author.name} ${author?.lastname}` }}</span>
      </div>

      <nav class="flex gap-2">
        <button>
          <Icon name="heroicons:heart" class="w-6 h-6" />
        </button>

        <button>
          <Icon name="heroicons:share" class="w-6 h-6" />
        </button>
      </nav>
    </div>
  </article>
</template>
