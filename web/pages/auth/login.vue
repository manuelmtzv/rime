<script setup lang="ts">
import { useVuelidate } from "@vuelidate/core";
import { getErrorMessage } from "~/utils/errors";

const { login } = useAuth();
const localePath = useLocalePath();
const { required } = useValidation();
const toast = useToast();
const { t } = useI18n();
const authError = ref<string>();

const form = reactive({
  username: "",
  password: "",
});

const rules = {
  username: { required },
  password: { required },
};

const v$ = useVuelidate(rules, form);

async function handleSubmit() {
  await v$.value.$validate();
  if (v$.value.$error) return;

  try {
    await login(form);
    toast.add({ title: t("auth.login.success") });
    await navigateTo(localePath("/"));
  } catch (error) {
    authError.value = getErrorMessage(error);
    toast.add({ title: t("auth.login.failed") });
  }
}
</script>

<template>
  <Page class="">
    <form class="form" @submit.prevent="handleSubmit">
      <div class="text-center space-y-2 mb-2">
        <h2 class="font-semibold text-lg font-poetry">
          {{ $t("auth.login.title") }}
        </h2>
        <p>{{ $t("auth.login.subtitle") }}</p>
      </div>

      <UFormGroup :error="toValue(v$.username.$errors[0]?.$message)">
        <UInput v-model="form.username" size="md" :placeholder="$t('auth.login.username')" />
      </UFormGroup>

      <UFormGroup :error="toValue(v$.password.$errors[0]?.$message)">
        <UInput v-model="form.password" size="md" :placeholder="$t('auth.login.password')" />
      </UFormGroup>

      <div v-if="authError" class="text-red-500 text-sm">{{ authError }}</div>

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
