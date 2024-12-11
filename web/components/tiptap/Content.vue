<script setup lang="ts">
// const { t } = useI18n();

type TiptapContentProps = {
  content: string;
  title?: string;
  class?: string;
  contentClass?: string;
};

const props = defineProps<TiptapContentProps>();
</script>

<template>
  <div :class="cn('flex flex-col gap-2 overflow-y-auto', $props.class)">
    <h1 v-if="title" class="font-medium text-lg">{{ title }}</h1>

    <div :class="cn('tiptap tiptap-content font-poetry space-y-1.5', contentClass)" v-html="content"></div>
  </div>
</template>

<style lang="scss">
.tiptap-content {

  // br {
  //   display: block;
  //   margin: 0.62rem 0 !important;
  //   content: "";
  // }
  p {
    font-size: 0.9rem;
  }
}
</style>
