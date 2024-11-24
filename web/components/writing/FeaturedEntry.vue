<script setup lang="ts">
import type { Writing } from "@/types";

type Props = {
  writing: Writing;
};

const props = defineProps<Props>();
const localePath = useLocalePath();
const { user } = useUserState();

const author = computed(() => props.writing.author);

const likedWriting = computed(() => {
  return Boolean(
    props.writing.likes?.some((like) => like.author?.id === user.value?.id)
  );
});
</script>

<template>
  <article
    class="p-6 mx-auto flex flex-col gap-4 bg-gradient-to-r from-[#8B4513] via-[#6F4E37] to-[#4A3728] dark:from-[#3F4E4F] dark:via-[#364245] dark:to-[#2C3639] rounded-lg shadow-lg text-white"
  >
    <h1 class="underline">{{ $t("writing.featured.title") }}</h1>

    <NuxtLink :to="localePath(`/writings/${writing.id}`)">
      <h2 class="text-2xl font-bold font-poetry mb-4">{{ writing.title }}</h2>

      <TiptapContent class="font-poetry" :content="writing.text" />
    </NuxtLink>

    <div class="flex justify-between items-center gap-4 mt-2">
      <div v-if="author" class="space-x-2">
        <UAvatar :alt="`${author.name} ${author?.lastname}`" size="sm" />
        <span class="text-sm">{{ `${author.name} ${author?.lastname}` }}</span>
      </div>

      <nav class="flex gap-2">
        <LikeButton
          :liked="likedWriting"
          entity="writings"
          :entity-id="writing.id"
        />

        <button>
          <Icon name="heroicons:share" class="w-6 h-6" />
        </button>
      </nav>
    </div>
  </article>
</template>
