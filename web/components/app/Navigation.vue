<template>
  <nav>
    <ul class="flex gap-2">
      <li>
        <AppLink to="/">
          {{ $t("app.header.home") }}
        </AppLink>
      </li>

      <li>
        <AppLink to="/auth/login">
          {{ $t("app.header.login") }}
        </AppLink>
      </li>
      <li>
        <AppLink to="/auth/register">
          {{ $t("app.header.register") }}
        </AppLink>
      </li>
    </ul>
  </nav>
</template>
