<script setup lang="ts">
const showPassword = ref(false);
const localePath = useLocalePath();

const form = reactive({
  name: "",
  lastname: "",
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
        <h2 class="font-semibold text-lg font-poetry">
          {{ $t("auth.register.title") }}
        </h2>
        <p>{{ $t("auth.register.subtitle") }}</p>
      </div>
      <UInput
        size="md"
        :placeholder="$t('auth.register.name')"
        v-model="form.name"
      />

      <UInput
        size="md"
        :placeholder="$t('auth.register.lastname')"
        v-model="form.lastname"
      />

      <UInput
        size="md"
        :placeholder="$t('auth.register.username')"
        v-model="form.username"
      />

      <UInput
        size="md"
        :placeholder="$t('auth.register.email')"
        v-model="form.email"
      />

      <UInput
        size="md"
        :placeholder="$t('auth.register.password')"
        v-model="form.password"
      />

      <UInput
        size="md"
        :placeholder="$t('auth.register.confirmPassword')"
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
        <span class="text-sm">{{ $t("auth.register.hasAccount") }}</span>

        <NuxtLink :to="localePath('/auth/login')" class="text-sm">{{
          $t("auth.register.loginLink")
        }}</NuxtLink>
      </div>

      <UButton type="submit" class="justify-center font-semibold" size="md">{{
        $t("auth.register.submit")
      }}</UButton>
    </form>
  </Page>
</template>
