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
    <UAvatar size="sm" :alt="fullName" />

    <div>
      <h3 class="font-medium text-sm">{{ fullName }}</h3>
      <p class="text-xs">{{ popularUser.followers }} followers</p>
    </div>
  </div>
</template>
