<script setup lang="ts">
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
  <Card class="space-y-2">
    <h3>Quick Actions</h3>

    <nav>
      <ul class="space-y-2">
        <li v-for="action in actions">
          <NuxtLink :to="action.to">
            <UButton variant="link" class="w-full" :icon="action.icon">
              {{ action.title }}
            </UButton>
          </NuxtLink>
        </li>
      </ul>
    </nav>
  </Card>
</template>
