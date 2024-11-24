<script setup lang="ts">
const { showAuthRequiredModal, setShowAuthRequiredModal } = useAppState();
const localePath = useLocalePath();
</script>

<template>
  <UModal
    v-model="showAuthRequiredModal"
    :ui="{
      container: 'max-w-[22rem] mx-auto',
    }"
  >
    <div class="flex flex-col gap-5 p-6">
      <div class="text-center space-y-2">
        <h3 class="text-lg font-semibold">{{ $t("auth.authAction.title") }}</h3>
        <p>{{ $t("auth.authAction.subtitle") }}</p>
      </div>

      <div class="flex flex-col gap-2">
        <NuxtLink
          :to="localePath('/auth/register')"
          @click="setShowAuthRequiredModal(false)"
        >
          <UButton block>
            {{ $t("auth.authAction.register") }}
          </UButton>
        </NuxtLink>

        <NuxtLink
          :to="localePath('/auth/login')"
          @click="setShowAuthRequiredModal(false)"
        >
          <UButton block variant="outline">
            {{ $t("auth.authAction.login") }}
          </UButton>
        </NuxtLink>
      </div>

      <UButton
        variant="ghost"
        size="sm"
        icon="i-heroicons-x-mark-20-solid"
        class="flex z-10 absolute right-4 top-4"
        square
        padded
        @click="setShowAuthRequiredModal(false)"
      />
    </div>
  </UModal>
</template>
