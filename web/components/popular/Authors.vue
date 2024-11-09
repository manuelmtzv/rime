<script setup lang="ts">
import { userRepository } from "@/repositories/user.repository";

const { data } = await useAsyncData(
  "popular-authors",
  userRepository().getPopularUsers
);
</script>

<template>
  <Card v-if="data" class="flex flex-col gap-2">
    <div>Popular Writers</div>

    <ul class="space-y-2">
      <li v-for="user in data.data">
        <UserInlineEntry :popular-user="user" />
      </li>
    </ul>
  </Card>
</template>
