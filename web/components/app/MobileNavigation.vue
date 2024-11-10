<script setup lang="ts">
const isOpen = ref(false);

function closeSlideover() {
  isOpen.value = false;
}
</script>

<template>
  <div>
    <UButton
      icon="i-heroicons-bars-3"
      variant="outline"
      @click="isOpen = true"
    />

    <USlideover v-model="isOpen" side="left">
      <div class="flex flex-col gap-3 p-4 flex-1 overflow-y-auto">
        <div class="flex gap-4 justify-between items-center pt-3 pb-2">
          <AppLogo />

          <UButton
            variant="outline"
            size="sm"
            icon="i-heroicons-x-mark-20-solid"
            class="flex sm:hidden z-10"
            square
            padded
            @click="isOpen = false"
          />
        </div>

        <section class="flex-1 space-y-6 py-2 overflow-y-auto">
          <div>
            <h3 class="mb-2">Navigation</h3>
            <AppNavigation mobile @link-click="closeSlideover" />
          </div>

          <WritingQuickActions mobile />
        </section>

        <div
          class="flex items-center justify-between gap-4 w-full bg-white dark:bg-grey-900 py-2"
        >
          <AppLocaleSelect />

          <AppThemeToggle button-class="rounded-md border" />
        </div>
      </div>
    </USlideover>
  </div>
</template>
