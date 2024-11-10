<script setup lang="ts">
import { writingRepository } from "@/repositories/writing.repository";

const id = useRoute().params.id as string;
const { getWriting } = writingRepository();

const { data } = useAsyncData(`writing-${id}`, async () => {
  return await getWriting(id);
});

const writing = computed(() => data?.value?.data);
</script>

<template>
  <Page>
    <section v-if="writing" class="max-w-2xl mx-auto mt-8 mb-12 space-y-4">
      <WritingEntry :writing="writing" />
    </section>
  </Page>
</template>
