<script setup lang="ts">
const { user, fullName } = useUserState();
const localePath = useLocalePath();
</script>

<template>
  <Card v-if="user">
    <div class="flex flex-col gap-2 items-center">
      <UAvatar :alt="fullName" size="lg" />
      <div class="w-full text-center">
        <h3 class="font-semibold">{{ fullName }}</h3>
        <p class="text-sm text-gray-500">{{ user.username }}</p>
      </div>

      <div class="grid grid-cols-2 gap-4 mt-1 w-full">
        <NuxtLink :to="localePath('/profile')" class="flex-1">
          <UButton
            class="w-full justify-center"
            size="xs"
            icon="i-heroicons-user-circle-solid"
          >
            {{ $t("user.card.profile") }}
          </UButton>
        </NuxtLink>

        <NuxtLink :to="localePath('/settings')" class="flex-1">
          <UButton
            class="w-full justify-center"
            size="xs"
            icon="i-heroicons-cog-6-tooth-solid"
            variant="soft"
          >
            {{ $t("user.card.settings") }}
          </UButton>
        </NuxtLink>
      </div>
    </div>
  </Card>
</template>
