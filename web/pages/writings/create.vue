<script setup lang="ts">
import { writingRepository } from "@/repositories";
import useVuelidate from "@vuelidate/core";

const { required } = useValidation();
const { loading } = useFormState();
const { createWriting } = writingRepository();
const { user } = useUserState();
const toast = useToast();
const localePath = useLocalePath();

const form = reactive({
  title: "",
  content: "",
});

const rules = {
  title: { required },
  content: { required },
};

const v$ = useVuelidate(rules, form);

async function submitHandler() {
  await v$.value.$validate();
  if (v$.value.$error) return;

  try {
    loading.value = true;
    const response = await createWriting({
      ...form,
      type: "poetry",
    });

    await navigateTo(localePath(`/writings/${response.data.id}`));
  } catch (err) {
    toast.add({
      color: "yellow",
      title: getErrorMessage(err),
    });
  } finally {
    loading.value = false;
  }
}

const emptyContent = computed(() => getHtmlText(form.content).length === 0);
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
            <InputWrapper
              :label="$t('writing.editor.title')"
              :error="toValue(v$.title.$errors[0]?.$message)"
              class="mb-4"
            >
              <UInput
                v-model="form.title"
                :placeholder="$t('writing.editor.titlePlaceholder')"
                :class="['w-full']"
            /></InputWrapper>
          </template>

          <template #after-editor>
            <span
              class="text-sm text-red-500 mt-1"
              v-if="emptyContent && v$.content.$dirty"
            >
              {{ $t("writing.editor.emptyContent") }}
            </span>
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
            class="btn-primary"
            variant="soft"
            :disabled="loading"
            @click.prevent=""
            >{{ $t("writing.editor.save") }}
          </UButton>

          <UButton
            class="btn-primary"
            color="black"
            :disabled="loading"
            @click.prevent="submitHandler"
            >{{ $t("writing.editor.saveAndPublish") }}</UButton
          >
        </div>
      </nav>
    </form>
  </Page>
</template>
