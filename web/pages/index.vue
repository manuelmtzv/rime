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
    <div class="flex gap-8 mx-auto w-full max-w-6xl">
      <aside class="space-y-8 h-fit w-full max-w-52">
        <Card class="flex flex-col gap-2">
          <div class="inline-flex gap-2 items-center">
            <Icon name="heroicons:arrow-trending-up" />
            <span>Trending Tags</span>
          </div>

          <ul class="space-y-1">
            <li>
              <a href="#" class="text-blue-500 text-sm">#PoetryMonth</a>
            </li>
            <li>
              <a href="#" class="text-blue-500 text-sm">#Love</a>
            </li>
            <li>
              <a href="#" class="text-blue-500 text-sm">#Life</a>
            </li>
            <li>
              <a href="#" class="text-blue-500 text-sm">#Dreams</a>
            </li>
            <li>
              <a href="#" class="text-blue-500 text-sm">#Haiku</a>
            </li>
          </ul>
        </Card>

        <Card class="flex flex-col gap-2">
          <div>Popular Writers</div>

          <ul class="space-y-2">
            <li>
              <UserInlineEntry />
            </li>
            <li>
              <UserInlineEntry />
            </li>
            <li>
              <UserInlineEntry />
            </li>
            <li>
              <UserInlineEntry />
            </li>
            <li>
              <UserInlineEntry />
            </li>
            <li>
              <UserInlineEntry />
            </li>
            <li>
              <UserInlineEntry />
            </li>
            <li>
              <UserInlineEntry />
            </li>
          </ul>
        </Card>
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

      <aside class="space-y-4 h-fit w-full max-w-52">
        <Card> Hola </Card>
        <Card> Hola </Card>
      </aside>
    </div>
  </Page>
</template>
