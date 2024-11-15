<script setup lang="ts">
const form = reactive({
  title: '',
  content: ''
})
</script>

<template>
  <Page>
    <form action="" class="form max-w-4xl mt-8 border-none shadow-none padding-0">
      <h1 class="text-lg md:text-xl font-medium">Nuevo Escrito</h1>

      <ClientOnly>
        <TiptapEditor v-model="form.content" :tiptap-props="{ minHeight: '24rem' }" />
      </ClientOnly>

      <SimpleAlert
        message="Los saltos de línea múltiples de convertiran en uno sólo, esto para garantizar la consistencia de los escritos."
        class="my-2" />

      <nav class="flex items-center justify-between">
        <UButton type="button" class="btn-secondary" variant="ghost" color="black">Cancelar</UButton>

        <div class="flex gap-4">
          <UButton type="submit" class="btn-primary" variant="soft" @click.prevent="">Crear borrador</UButton>
          <UButton type="submit" class="btn-primary" color="black" @click.prevent="">Guardar y publicar</UButton>
        </div>
      </nav>
    </form>
  </Page>
</template>
