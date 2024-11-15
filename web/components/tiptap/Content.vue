import { important } from '../../.nuxt/types/tailwind.config';
<script setup lang="ts">
// const { t } = useI18n();

type TiptapContentProps = {
  content: string;
  title?: string;
  class?: string;
  contentClass?: string;
};

const props = defineProps<TiptapContentProps>();

const previewContent = computed(() =>
  props.content == "<p></p>" ? "Preview" : props.content
);
</script>

<template>
  <div :class="cn('flex flex-col gap-2 overflow-y-auto', $props.class)">
    <h1 v-if="title" class="font-medium text-lg">{{ title }}</h1>

    <div :class="cn('tiptap tiptap-content space-y-1 font-poetry', contentClass)" v-html="previewContent"></div>
  </div>
</template>

<style lang="postcss">
.tiptap-content {
  br {
    display: block;
    margin: 1.25rem 0 !important;
    content: "";
  }
}
</style>
