<script setup lang="ts">
import { tagRepository } from "@/repositories/tag.repository";

const { data } = await useAsyncData("popular-tags", tagRepository().getTags);
</script>

<template>
  <Card class="flex flex-col gap-2">
    <template v-if="data">
      <div class="inline-flex gap-2 items-center">
        <Icon name="heroicons:arrow-trending-up" />
        <h3>{{ $t("tag.popular.title") }}</h3>
      </div>

      <ul class="space-y-1">
        <li v-for="tag in data.data">
          <a href="#" class="tag">#{{ tag.name }}</a>
        </li>
      </ul>
    </template>
  </Card>
</template>

<style scoped>
.tag {
  @apply text-sm;
}
</style>
