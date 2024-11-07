<script setup lang="ts">
const route = useRoute();
</script>

<template>
  <Page>
    <pre>{{ route.params }}</pre>
  </Page>
</template>
