<script setup lang="ts">
import type { Writing } from "@/types";

type Props = {
  writing: Writing;
  showCreatedAt?: boolean;
};

const toast = useToast();
const props = withDefaults(defineProps<Props>(), {
  showCreatedAt: true,
});
const { t } = useI18n();
const authAction = useAuthAction();

const author = computed(() => props.writing.author);
const localePath = useLocalePath();

async function likeHandler() {
  authAction.run(() => {
    console.log("Like");
  });
}

function shareHandler() {
  const url = useRequestURL().host + `/writings/${props.writing.id}`;
  copyToClipboard(url);

  toast.add({
    title: t("app.copiedLink"),
  });
}
</script>

<template>
  <article class="px-2 mx-auto py-6 flex flex-col gap-4 border-b">
    <NuxtLink :to="localePath(`/writings/${writing.id}`)">
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
            {{ $d(new Date(writing.createdAt)) }}
          </span>
        </div>
      </div>

      <nav class="flex gap-2">
        <button @click="likeHandler">
          <Icon name="heroicons:heart" class="w-6 h-6" />
        </button>

        <button @click="shareHandler">
          <Icon name="heroicons:share" class="w-6 h-6" />
        </button>
      </nav>
    </div>
  </article>
</template>
