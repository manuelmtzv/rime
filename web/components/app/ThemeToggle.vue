<script setup lang="ts">
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
  <ColorScheme placeholder="..." tag="span" class="w-8 h-8">
    <button
      @click="toggleColorMode"
      class="flex hover:bg-gray-100 dark:hover:bg-gray-900 p-1.5 rounded-md"
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
  </ColorScheme>
</template>
