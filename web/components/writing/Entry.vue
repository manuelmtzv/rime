<script setup lang="ts">
import type { Writing } from "@/types";
import { useUserState } from "@/composables/useUserState";

type Props = {
  writing: Writing;
  link?: boolean;
  showCreatedAt?: boolean;
};

const toast = useToast();
const props = withDefaults(defineProps<Props>(), {
  showCreatedAt: true,
  link: true,
});
const { t } = useI18n();
const { user } = useUserState();

const localePath = useLocalePath();
const author = computed(() => props.writing.author);

function shareHandler() {
  const url = useRequestURL().host + `/writings/${props.writing.id}`;
  copyToClipboard(url);

  toast.add({
    title: t("app.copiedLink"),
  });
}

const to = localePath(`/writings/${props.writing.id}`);

const likedWriting = computed(() => {
  return Boolean(
    props.writing.likes?.some((like) => like.author?.id === user.value?.id)
  );
});
</script>

<template>
  <article class="px-2 mx-auto py-6 flex flex-col gap-4 border-b">
    <NuxtLink :to="link ? to : undefined">
      <h2 class="text-2xl font-bold font-poetry mb-4">{{ writing.title }}</h2>

      <TiptapContent class="font-poetry" :content="writing.text" />
    </NuxtLink>

    <div class="flex justify-between items-center gap-4 mt-2">
      <div v-if="author" class="flex gap-2">
        <UAvatar :alt="`${author.name} ${author?.lastname}`" size="md" />

        <div class="flex flex-col">
          <span class="text-sm">{{
            `${author.name} ${author?.lastname}`
          }}</span>
          <span v-if="showCreatedAt" class="text-sm text-gray-500">
            <ClientOnly fallback="...">
              {{ $d(new Date(writing.createdAt)) }}
            </ClientOnly>
          </span>
        </div>
      </div>

      <nav class="flex gap-2">
        <LikeButton :liked="likedWriting" entity="writings" :entity-id="writing.id" />

        <button @click="shareHandler">
          <Icon name="heroicons:share" class="w-6 h-6" />
        </button>
      </nav>
    </div>
  </article>
</template>
