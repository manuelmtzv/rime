<script setup lang="ts">
import { userRepository } from "@/repositories/user.repository";

const { data } = await useAsyncData(
  "popular-authors",
  userRepository().getPopularUsers
);
</script>

<template>
  <Card class="flex flex-col gap-2">
    <template v-if="data">
      <div class="mb-1">
        <h3>{{ $t("user.popular.title") }}</h3>
      </div>

      <ul class="space-y-2">
        <li v-for="user in data.data">
          <UserInlineEntry :popular-user="user" />
        </li>
      </ul>
    </template>
  </Card>
</template>
