<script setup lang="ts">
import type { Writing } from "@/types";

type Props = {
    writing: Writing;
};

const props = defineProps<Props>()
const author = computed(() => props.writing.author);
</script>

<template>
    <article class="px-2 mx-auto py-6 flex flex-col gap-4 border-b">
        <div>
            <h2 class="text-2xl font-bold font-poetry mb-4">{{ writing.title }}</h2>

            <TiptapContent class="font-poetry" :content="writing.text" />
        </div>

        <div class="flex justify-between items-center gap-4 mt-2">
            <div v-if="author" class="flex gap-2">
                <UAvatar :alt="`${author.name} ${author?.lastname}`" size="md" />

                <div class="flex flex-col">
                    <span class="text-sm">{{
                        `${author.name} ${author?.lastname}`
                        }}</span>
                    <span class="text-sm text-gray-500">
                        <ClientOnly fallback="...">
                            {{ $d(new Date(writing.createdAt)) }}
                        </ClientOnly>
                    </span>
                </div>
            </div>

            <nav class="flex gap-2 ml-auto">
                <Icon name="heroicons:heart" class="w-6 h-6" />
                <Icon name="heroicons:share" class="w-6 h-6" />
            </nav>
        </div>
    </article>
</template>
