<script setup lang="ts">
const { user } = useUserState();
</script>

<template>
  <Page>
    <pre>
      {{ user }}
    </pre>
  </Page>
</template>
