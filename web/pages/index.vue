<script setup lang="ts">
import { writingRepository } from "@/repositories/writing.repository";

const { data, error } = await useAsyncData(
  "writings",
  writingRepository().getWritings
);

const value = ref("");
</script>

<template>
  <Page>
    <pre v-if="!error">
      {{ data }}
    </pre>

    <pre v-else>
      {{ error }}
    </pre>
  </Page>
</template>
