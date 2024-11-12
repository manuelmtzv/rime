<script setup lang="ts">
import { tagRepository } from "@/repositories/tag.repository";

const { data, status, error } = await useAsyncData("popular-tags", tagRepository().getTags);
</script>

<template>
  <Card class="flex flex-col gap-2">
    <div class="inline-flex gap-2 items-center mb-1">
      <Icon name="heroicons:arrow-trending-up" />
      <h3>{{ $t("tag.popular.title") }}</h3>
    </div>

    <ul class="space-y-1">
      <template v-if="data">
        <li v-for="tag in data.data">
          <a href="#" class="tag">#{{ tag.name }}</a>
        </li>
      </template>

      <template v-if="status == 'pending'">
        <li v-for="i in 6">
          <USkeleton class="w-20 h-4" />
        </li>
      </template>
    </ul>

    <div v-if="status == 'error'">
      <pre>{{ error }}</pre>
    </div>
  </Card>
</template>

<style scoped lang="postcss">
.tag {
  @apply text-sm hover:underline;
}
</style>
