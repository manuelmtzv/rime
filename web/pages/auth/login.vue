<script setup lang="ts">
function handleSubmit() {
  console.log("submit");
}
</script>

<template>
  <Page class="">
    <form class="auth-form mt-20 mb-10" @submit.prevent="handleSubmit">
      <div class="text-center space-y-2 mb-2">
        <h2 class="font-semibold text-lg">Welcome back!</h2>
        <p>Enter your credentials to access your writings</p>
      </div>

      <UInput size="md" placeholder="Your email or username" />

      <UInput size="md" placeholder="Your password" />

      <div class="flex justify-between items-center">
        <ULink href="/auth/register" class="text-sm"
          >Don't have an account?</ULink
        >

        <ULink href="/auth/register" class="text-sm">Create an account</ULink>
      </div>

      <UButton type="submit" class="justify-center font-semibold" size="md"
        >Login</UButton
      >
    </form>
  </Page>
</template>
