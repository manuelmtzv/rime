<script setup lang="ts">
const localePath = useLocalePath();

const { user } = useUserState();
</script>

<template>
  <nav>
    <ul class="flex gap-2">
      <li>
        <AppLink :to="localePath('/')">
          {{ $t("app.header.home") }}
        </AppLink>
      </li>

      <template v-if="user">
        <li>
          <AppLink :to="localePath('/profile')">
            {{ $t("app.header.profile") }}
          </AppLink>
        </li>
      </template>

      <template v-else>
        <li>
          <AppLink :to="localePath('/auth/login')">
            {{ $t("app.header.login") }}
          </AppLink>
        </li>
        <li>
          <AppLink :to="localePath('/auth/register')">
            {{ $t("app.header.register") }}
          </AppLink>
        </li>
      </template>
    </ul>
  </nav>
</template>
