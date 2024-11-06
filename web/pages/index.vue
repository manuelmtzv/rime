<script setup lang="ts">
import { writingRepository } from "@/repositories/writing.repository";

const { data, error } = await useAsyncData(
  "writtings",
  writingRepository().getWritings
);

const value = ref("");
</script>

<template>
  <Page>
    <TiptapEditor v-model="value" />
  </Page>
</template>
