<script setup lang="ts">
const options = [
  {
    name: "English",
    value: "en",
  },
  {
    name: "Español",
    value: "es",
  },
];

type Props = {
  class?: string;
};

defineProps<Props>();

const { localeProperties, locales, setLocale } = useI18n();

const currentCode = computed(() => localeProperties.value.code);
</script>

<template>
  <USelect
    icon="i-mdi-web"
    :value="currentCode"
    :options="options"
    option-attribute="name"
    :class="cn($props.class)"
    @update:model-value="setLocale"
  />

  <button @click="setLocale('en')">En</button>
  <button @click="setLocale('es')">Es</button>
</template>
