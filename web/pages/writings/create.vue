<script setup lang="ts">
const form = reactive({
  title: "",
  content: "",
});
</script>

<template>
  <Page>
    <form class="form max-w-4xl mt-8 border-none shadow-none p-0">
      <ClientOnly>
        <TiptapEditor
          v-model="form.content"
          :preview="false"
          :tiptap-props="{ minHeight: '24rem' }"
        >
          <template #before-editor>
            <InputWrapper :label="$t('writing.editor.title')">
              <UInput
                v-model="form.title"
                :placeholder="$t('writing.editor.titlePlaceholder')"
                :class="['w-full mb-4']"
            /></InputWrapper>
          </template>
        </TiptapEditor>

        <template #fallback>
          <div class="flex flex-col gap-4 rounded-md min-h-[32.375rem]">
            <USkeleton class="w-full h-10" />

            <USkeleton class="w-full h-[3.75rem]" />

            <USkeleton class="w-full flex-1" />
          </div>
        </template>
      </ClientOnly>

      <SimpleAlert
        :message="$t('writing.editor.lineBreaksAdvice')"
        class="my-2"
      />

      <nav class="flex items-center justify-between">
        <UButton
          type="button"
          class="btn-secondary"
          variant="ghost"
          color="black"
          >{{ $t("writing.editor.cancel") }}
        </UButton>

        <div class="flex gap-4">
          <UButton
            type="submit"
            class="btn-primary"
            variant="soft"
            @click.prevent=""
            >{{ $t("writing.editor.save") }}
          </UButton>
          <UButton
            type="submit"
            class="btn-primary"
            color="black"
            @click.prevent=""
            >{{ $t("writing.editor.saveAndPublish") }}</UButton
          >
        </div>
      </nav>
    </form>
  </Page>
</template>
