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

type LocaleSelectProps = {
  minimal?: boolean;
};

defineProps<LocaleSelectProps>();

const id = useId();
const { localeProperties, locales, setLocale } = useI18n();

const displayableLocales = computed(() =>
  locales.value.filter((locale) => locale.code !== localeProperties.value.code)
);
</script>

<template>
  <USelect
    icon="i-mdi-web"
    :model-value="localeProperties.code"
    :options="options"
    option-attribute="name"
    @update:model-value="setLocale"
  />
</template>
