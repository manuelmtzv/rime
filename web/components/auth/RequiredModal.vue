<script setup lang="ts">
const { showAuthRequiredModal, setShowAuthRequiredModal } = useAppState();
const localePath = useLocalePath();
</script>

<template>
  <UModal
    v-model="showAuthRequiredModal"
    :ui="{
      container: 'max-w-[20rem] mx-auto',
    }"
  >
    <div class="flex flex-col gap-4 p-6">
      <div class="text-center">
        <h3 class="text-lg font-semibold">Login or register</h3>
        <p>You need to login or register to continue.</p>
      </div>

      <div class="flex flex-col gap-2">
        <NuxtLink
          :to="localePath('/auth/register')"
          @click="setShowAuthRequiredModal(false)"
        >
          <UButton block>Register</UButton>
        </NuxtLink>

        <NuxtLink
          :to="localePath('/auth/login')"
          @click="setShowAuthRequiredModal(false)"
        >
          <UButton block variant="outline">Login</UButton>
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
