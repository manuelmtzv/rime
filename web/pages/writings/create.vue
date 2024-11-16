<script setup lang="ts">
const form = reactive({
  title: "",
  content: "",
});
</script>

<template>
  <Page>
    <form
      action=""
      class="form max-w-4xl mt-8 border-none shadow-none padding-0"
    >
      <h1 class="text-lg md:text-xl font-medium">
        {{ $t("writing.editor.newTitle") }}
      </h1>

      <ClientOnly>
        <TiptapEditor
          v-model="form.content"
          :tiptap-props="{ minHeight: '24rem' }"
        />
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
          >{{ $t("writing.editor.cancel") }}</UButton
        >

        <div class="flex gap-4">
          <UButton
            type="submit"
            class="btn-primary"
            variant="soft"
            @click.prevent=""
            >{{ $t("writing.editor.save") }}</UButton
          >
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
