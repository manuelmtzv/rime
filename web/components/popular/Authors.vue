<script setup lang="ts">
import { userRepository } from "@/repositories/user.repository";

const { data, status, error } = await useLazyAsyncData(
  "popular-authors",
  userRepository().getPopularUsers
);
</script>

<template>
  <Card class="flex flex-col gap-2">
    <div class="mb-1">
      <h3>{{ $t("user.popular.title") }}</h3>
    </div>

    <ul class="space-y-2">
      <template v-if="data">
        <li v-for="user in data.data">
          <UserInlineEntry :popular-user="user" />
        </li>
      </template>

      <template v-else-if="status == 'pending'">
        <UserInlineEntryPlaceholder v-for="i in 8" :key="i" />
      </template>
    </ul>

    <div v-if="status == 'error'">
      <pre>{{ error }}</pre>
    </div>
  </Card>
</template>
