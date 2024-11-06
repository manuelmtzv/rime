<script setup lang="ts">
import { writingRepository } from "@/repositories/writing.repository";

const { data, error } = await useAsyncData(
  "writtings",
  writingRepository().getWritings
);
</script>

<template>
  <Page>
    <pre v-if="!error">{{ data }}</pre>

    <pre v-else>
      {{ error }}
    </pre>
  </Page>
</template>
