<script setup lang="ts">
import type { Writing } from '@/types';

const { t } = useI18n();

const form = reactive({
  title: t("writing.editor.newTitle"),
  content: "",
});

const showPreview = ref(false)
const { user } = useUserState()

function togglePreview() {
  showPreview.value = !showPreview.value
}

const newWriting = computed(() => ({
  id: 'new',
  title: form.title,
  text: form.content,
  author: user.value,
  createdAt: new Date().toISOString(),
  updatedAt: new Date().toISOString(),
  type: 'poetry',
} as Writing))
</script>

<template>
  <Page>
    <form class="form max-w-4xl mt-8 border-none shadow-none padding-0">
      <ClientOnly class="rounded-md min-h-[30.5rem]">
        <TiptapEditor v-model="form.content" :preview="false" :tiptap-props="{ minHeight: '24rem' }"
          :hide-editor="showPreview">
          <template #actions>
            <TiptapButtonWrapper :is-active="showPreview" @click.prevent="togglePreview"
              class="ml-auto rounded-lg px-2">
              <span class="text-xs"> Preview </span>
            </TiptapButtonWrapper>
          </template>

          <template #before-editor>
            <UInput v-if="!showPreview" v-model="form.title" label="Title" placeholder="Title"
              :class="['w-full', !showPreview && 'mb-4']" />
          </template>
        </TiptapEditor>
      </ClientOnly>

      <WritingPreviewEntry v-if="showPreview" :writing="newWriting" class="w-full border-b px-6" />

      <SimpleAlert :message="$t('writing.editor.lineBreaksAdvice')" class="my-2" />

      <nav class="flex items-center justify-between">
        <UButton type="button" class="btn-secondary" variant="ghost" color="black">{{ $t("writing.editor.cancel") }}
        </UButton>

        <div class="flex gap-4">
          <UButton type="submit" class="btn-primary" variant="soft" @click.prevent="">{{ $t("writing.editor.save") }}
          </UButton>
          <UButton type="submit" class="btn-primary" color="black" @click.prevent="">{{
            $t("writing.editor.saveAndPublish") }}</UButton>
        </div>
      </nav>
    </form>
  </Page>
</template>
