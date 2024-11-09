<script setup lang="ts">
import { writingRepository } from "@/repositories/writing.repository";
import type { ListResponse, Writing } from "@/types";

const { data } = await useAsyncData<ListResponse<Writing>>(
  "writings",
  writingRepository().getWritings
);
</script>

<template>
  <Page class="mt-6 mb-12">
    <div class="flex gap-8 mx-auto w-full max-w-7xl">
      <aside class="space-y-8 h-fit w-full max-w-60">
        <PopularTags />

        <PopularAuthors />
      </aside>

      <section class="space-y-8 flex-1">
        <AppSearchBar />

        <template v-if="data">
          <WritingFeaturedEntry :writing="data.data[0]" />

          <WritingEntry
            v-for="writing in data.data"
            :key="writing.id"
            :writing="writing"
          />
        </template>
      </section>

      <aside class="space-y-4 h-fit w-full max-w-56">
        <Card> Hola </Card>
        <Card> Hola </Card>

        <AppLocaleSelect />
      </aside>
    </div>
  </Page>
</template>
