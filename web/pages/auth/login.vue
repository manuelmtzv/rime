<script setup lang="ts">
function handleSubmit() {
  console.log("submit");
}
</script>

<template>
  <Page class="">
    <form class="auth-form" @submit.prevent="handleSubmit">
      <div class="text-center space-y-2 mb-2">
        <h2 class="font-semibold text-lg font-poetry">Welcome back!</h2>
        <p>Enter your credentials to access your writings</p>
      </div>

      <UInput size="md" placeholder="Your email or username" />

      <UInput size="md" placeholder="Your password" />

      <div class="flex justify-between items-center my-1">
        <span class="text-sm">Don't have an account?</span>

        <NuxtLink to="/auth/register" class="text-sm"
          >Create an account</NuxtLink
        >
      </div>

      <UButton type="submit" class="justify-center font-semibold" size="md"
        >Login</UButton
      >
    </form>
  </Page>
</template>
