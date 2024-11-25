<script setup lang="ts">
import { writingRepository } from "@/repositories/writing.repository";
import type { ListResponse, Writing } from "@/types";

const { data, status, error } = await useLazyAsyncData<ListResponse<Writing>>(
  "feed-writings",
  writingRepository().getWritings
);
</script>

<template>
  <Page class="mt-6 mb-12">
    <div class="flex gap-8 mx-auto w-full max-w-7xl relative">
      <aside class="hidden lg:block space-y-8 h-fit w-full max-w-60">
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

        <template v-else-if="status == 'pending'">
          <WritingEntryPlaceholder v-for="i in 8" :key="i" />
        </template>

        <div v-else-if="status == 'error'">
          <pre>{{ error }}</pre>
        </div>
      </section>

      <aside
        class="hidden md:block space-y-8 h-fit w-full max-w-60 sticky top-8"
      >
        <UserCard />

        <WritingQuickActions />

        <AppLocaleSelect class="hidden lg:block" />
      </aside>
    </div>
  </Page>
</template>
