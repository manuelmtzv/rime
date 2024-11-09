<script setup lang="ts">
import type { PopularUser } from "@/types";

type Props = {
  popularUser: PopularUser;
};

const props = defineProps<Props>();

const fullName = `${props.popularUser.name} ${props.popularUser.lastname}`;
</script>

<template>
  <div class="flex items-center gap-2">
    <NuxtLink :to="`/author/${popularUser.id}`">
      <UAvatar size="sm" :alt="fullName" />
    </NuxtLink>

    <div>
      <NuxtLink :to="`/author/${popularUser.id}`">
        <h3 class="font-medium text-sm">{{ fullName }}</h3>
      </NuxtLink>
      <p class="text-xs text-gray-600">
        <template v-if="popularUser.followers">
          {{ popularUser.followers }} {{ $t("user.popular.followers") }}
        </template>

        <template v-else> {{ $t("user.popular.newUser") }} </template>
      </p>
    </div>
  </div>
</template>
