<script setup lang="ts">
const localePath = useLocalePath();

function handleSubmit() {
  console.log("submit");
}
</script>

<template>
  <Page class="">
    <form class="auth-form" @submit.prevent="handleSubmit">
      <div class="text-center space-y-2 mb-2">
        <h2 class="font-semibold text-lg font-poetry">
          {{ $t("auth.login.title") }}
        </h2>
        <p>{{ $t("auth.login.subtitle") }}</p>
      </div>

      <UInput size="md" :placeholder="$t('auth.login.identifier')" />

      <UInput size="md" :placeholder="$t('auth.login.password')" />

      <div class="flex justify-between items-center my-1">
        <span class="text-sm">{{ $t("auth.login.noAccount") }}</span>

        <NuxtLink :to="localePath('/auth/register')" class="text-sm">
          {{ $t("auth.login.registerLink") }}
        </NuxtLink>
      </div>

      <UButton type="submit" class="justify-center font-semibold" size="md">{{
        $t("auth.login.submit")
      }}</UButton>
    </form>
  </Page>
</template>
