<script setup lang="ts">
const localePath = useLocalePath();
</script>

<template>
  <nav>
    <ul class="flex gap-2">
      <li>
        <AppLink :to="localePath('/')">
          {{ $t("app.header.home") }}
        </AppLink>
      </li>

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
    </ul>
  </nav>
</template>
