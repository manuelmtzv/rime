<script setup lang="ts">
type Props = {
  buttonClass?: string;
};

defineProps<Props>();
const colorMode = useColorMode();

const toggleColorMode = () => {
  if (colorMode.preference === "light") {
    colorMode.preference = "dark";
  } else {
    colorMode.preference = "light";
  }
};
</script>

<template>
  <span class="min-w-8">
    <ClientOnly>
      <button
        @click="toggleColorMode"
        :class="
          cn(
            'flex hover:bg-gray-100 dark:hover:bg-gray-900 p-1.5 rounded-md',
            buttonClass
          )
        "
      >
        <Icon
          v-show="colorMode.preference == 'dark'"
          name="heroicons:sun"
          size="20"
        />
        <Icon
          v-show="colorMode.preference == 'light'"
          name="heroicons:moon"
          size="20"
        />
      </button>

      <template #fallback>
        <div class="min-w-8 max-w-8 min-h-8 max-h-8 overflow-hidden"></div>
      </template>
    </ClientOnly>
  </span>
</template>
