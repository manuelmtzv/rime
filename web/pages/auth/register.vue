<script setup lang="ts">
const showPassword = ref(false);

const form = reactive({
  name: "",
  email: "",
  username: "",
  password: "",
  confirmPassword: "",
});

function handleSubmit() {
  console.log("submit");
}

function togglePasswordVisibility() {
  showPassword.value = !showPassword.value;
}
</script>

<template>
  <Page class="">
    <form class="auth-form" @submit.prevent="handleSubmit">
      <div class="text-center space-y-2 mb-2">
        <h2 class="font-semibold text-lg">Join the community!</h2>
        <p>Create an account to start writing</p>
      </div>
      <UInput size="md" placeholder="Name" v-model="form.name" />

      <UInput size="md" placeholder="Username" v-model="form.username" />

      <UInput size="md" placeholder="Email" v-model="form.email" />

      <UInput size="md" placeholder="Password" v-model="form.password" />

      <UInput
        size="md"
        placeholder="Confirm password"
        :type="showPassword ? 'text' : 'password'"
        :ui="{ icon: { trailing: { pointer: '' } } }"
        v-model="form.confirmPassword"
      >
        <template #trailing>
          <UButton
            size="sm"
            variant="ghost"
            :padded="false"
            :icon="showPassword ? 'heroicons:eye' : 'heroicons:eye-slash'"
            @click.prevent="togglePasswordVisibility"
          ></UButton>
        </template>
      </UInput>

      <div class="flex justify-between items-center my-1">
        <span class="text-sm">Already have an account?</span>

        <NuxtLink to="/auth/login" class="text-sm">Go to login</NuxtLink>
      </div>

      <UButton type="submit" class="justify-center font-semibold" size="md"
        >Register</UButton
      >
    </form>
  </Page>
</template>
