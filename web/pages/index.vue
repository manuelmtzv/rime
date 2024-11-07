<script setup lang="ts">
import { writingRepository } from "@/repositories/writing.repository";
import type { ListResponse, Writing } from "@/types";

const { data } = await useAsyncData<ListResponse<Writing>>(
  "writings",
  writingRepository().getWritings
);
</script>

<template>
  <Page>
    <section class="space-y-8 mb-12">
      <template v-if="data">
        <WritingEntry
          v-for="writing in data.data"
          :key="writing.id"
          :writing="writing"
        />
      </template>
    </section>
  </Page>
</template>
