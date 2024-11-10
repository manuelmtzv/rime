<script setup lang="ts">
type Props = {
  mobile?: boolean;
};

defineProps<Props>();
const localePath = useLocalePath();

const actions = [
  {
    title: "Create Writing",
    to: localePath("/writings/create"),
    icon: "i-heroicons-pencil",
  },
  {
    title: "View Drafts",
    to: localePath("/writings/drafts"),
    icon: "i-heroicons-document-duplicate",
  },
  {
    title: "My Collection",
    to: localePath("/writings/collection"),
    icon: "i-heroicons-folder-open",
  },
];
</script>

<template>
  <Card class="space-y-2 border-none p-0">
    <h3 :class="cn('mb-1', mobile && 'mb-2')">Quick Actions</h3>

    <nav>
      <ul :class="cn('space-y-2', mobile && 'space-y-4')">
        <li v-for="action in actions">
          <NuxtLink :to="action.to">
            <UButton
              :variant="mobile ? 'outline' : 'link'"
              class="w-full"
              :size="mobile ? 'lg' : undefined"
              :icon="action.icon"
            >
              {{ action.title }}
            </UButton>
          </NuxtLink>
        </li>
      </ul>
    </nav>
  </Card>
</template>
