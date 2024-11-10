<script setup lang="ts">
type Props = {
  mobile?: boolean;
};

defineProps<Props>();
const localePath = useLocalePath();
const { t } = useI18n();

const actions = [
  {
    title: t("writing.quickActions.create"),
    to: localePath("/writings/create"),
    icon: "i-heroicons-pencil",
  },
  {
    title: t("writing.quickActions.viewDrafts"),
    to: localePath("/writings/drafts"),
    icon: "i-heroicons-document-duplicate",
  },
  {
    title: t("writing.quickActions.myCollection"),
    to: localePath("/writings/collection"),
    icon: "i-heroicons-folder-open",
  },
];
</script>

<template>
  <Card :class="cn('space-y-2', mobile && 'border-none p-0')">
    <h3 :class="cn('mb-1', mobile && 'mb-2')">
      {{ $t("writing.quickActions.title") }}
    </h3>

    <nav>
      <ul :class="cn('space-y-2', mobile && 'space-y-4')">
        <li v-for="action in actions">
          <NuxtLink :to="action.to">
            <UButton
              :variant="mobile ? 'outline' : 'link'"
              class="w-full"
              :size="mobile ? 'lg' : undefined"
              :icon="action.icon"
              :ui="{
                variant: {
                  outline: mobile
                    ? 'text-black dark:text-white opacity-75 dark:opacity-75'
                    : '',
                },
              }"
            >
              {{ action.title }}
            </UButton>
          </NuxtLink>
        </li>
      </ul>
    </nav>
  </Card>
</template>
