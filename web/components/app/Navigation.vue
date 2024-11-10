<script setup lang="ts">
type Props = {
  class?: string;
  listClass?: string;
  mobile?: boolean;
  onLinkClick?: () => void;
};

type Emits = {
  (event: "link-click"): void;
};

defineProps<Props>();
defineEmits<Emits>();
const localePath = useLocalePath();

const { user } = useUserState();
</script>

<template>
  <nav :class="cn($props.class)">
    <ul
      :class="
        cn('flex', mobile ? 'flex flex-col gap-4' : 'flex-row gap-2', listClass)
      "
    >
      <li>
        <AppLink :to="localePath('/')" :mobile="mobile" @click="onLinkClick">
          {{ $t("app.header.home") }}
        </AppLink>
      </li>

      <template v-if="user">
        <li>
          <AppLink
            :to="localePath('/profile')"
            :mobile="mobile"
            @click="onLinkClick"
          >
            {{ $t("app.header.profile") }}
          </AppLink>
        </li>
      </template>

      <template v-else>
        <li>
          <AppLink
            :to="localePath('/auth/login')"
            :mobile="mobile"
            @click="onLinkClick"
          >
            {{ $t("app.header.login") }}
          </AppLink>
        </li>
        <li>
          <AppLink
            :to="localePath('/auth/register')"
            :mobile="mobile"
            @click="onLinkClick"
          >
            {{ $t("app.header.register") }}
          </AppLink>
        </li>
      </template>
    </ul>
  </nav>
</template>
