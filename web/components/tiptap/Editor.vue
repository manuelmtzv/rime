<script setup lang="ts">
import { useEditor, EditorContent } from "@tiptap/vue-3";
import StarterKit from "@tiptap/starter-kit";
import Placeholder from "@tiptap/extension-placeholder";
import Underline from "@tiptap/extension-underline";

type TiptapEditorProps = {
  placeholder?: string;
  title?: string;
  tiptapProps?: {
    minHeight?: string;
  };
  preview?: boolean;
  hideEditor?: boolean;
};

const props = withDefaults(defineProps<TiptapEditorProps>(), {
  preview: true,
});
const value = defineModel<string>({ required: true });
const showPreview = ref(false);

const title = computed(() => props.title ?? "");

const editor = useEditor({
  content: value.value,
  onCreate({ editor }) {
    value.value = editor.getHTML();
  },
  onUpdate({ editor }) {
    value.value = cleanTiptapOutput(editor.getHTML());
  },
  extensions: [
    StarterKit.configure({
      heading: {
        levels: [2, 3],
      },
    }),
    Underline,
    Placeholder.configure({
      placeholder: props.placeholder ?? "Start typing...",
    }),
  ],
  editorProps: {
    attributes: {
      class: cn(
        "block w-full px-4 py-4 rounded-md shadow-sm focus:outline-none text-sm  dark:bg-dark-800 dark:border-dark-700 rounded-t-none min-h-44 overflow-y-auto"
      ),
    },
  },
});
</script>

<template>
  <div v-if="editor" :class="cn('rounded-md', !hideEditor && 'min-h-80')">
    <nav
      :class="
        cn(
          'border rounded-md px-2 py-1.5 flex gap-2 min-h-10',
          !hideEditor && 'mb-4'
        )
      "
    >
      <TiptapButtonWrapper
        :is-active="editor.isActive('paragraph')"
        @click.prevent="editor?.chain().focus().setParagraph().run()"
      >
        <Icon name="heroicons:pencil-solid" size="16" />
      </TiptapButtonWrapper>

      <TiptapButtonWrapper
        :is-active="editor.isActive('bold')"
        @click.prevent="editor?.chain().focus().toggleBold().run()"
      >
        <Icon name="heroicons:bold" size="16" />
      </TiptapButtonWrapper>

      <TiptapButtonWrapper
        :is-active="editor.isActive('italic')"
        @click.prevent="editor?.chain().focus().toggleItalic().run()"
      >
        <Icon name="heroicons:italic" size="16" />
      </TiptapButtonWrapper>

      <TiptapButtonWrapper
        :is-active="editor.isActive('underline')"
        @click.prevent="editor?.chain().focus().toggleUnderline().run()"
      >
        <Icon name="heroicons:underline" size="16" />
      </TiptapButtonWrapper>

      <TiptapButtonWrapper
        :is-active="editor.isActive('strike')"
        @click.prevent="editor?.chain().focus().toggleStrike().run()"
      >
        <Icon name="heroicons:strikethrough" size="16" />
      </TiptapButtonWrapper>

      <template v-if="preview">
        <TiptapButtonWrapper
          :is-active="showPreview"
          @click.prevent="showPreview = !showPreview"
          class="ml-auto rounded-lg px-2"
        >
          <span class="text-xs"> Preview </span>
        </TiptapButtonWrapper>
      </template>

      <slot name="actions" />
    </nav>

    <slot name="before-editor" />

    <EditorContent
      v-if="!showPreview"
      v-show="!hideEditor"
      :editor="editor"
      v-model="value"
      class="min-h-96 font-poetry border border-gray-300 dark:border-gray-300 rounded-md"
    />

    <TiptapContent
      v-else
      :content="value"
      :title="title"
      class="gap-2"
      content-class="block w-full rounded-md focus:outline-none text-sm dark:bg-dark-900 dark:border-dark-700 rounded-t-none min-h-96 border border-gray-100 dark:border-gray-500 p-4"
    />

    <slot name="after-editor" />
  </div>
</template>

<style lang="postcss">
.tiptap {
  @apply space-y-1.5;
}
</style>

<style lang="scss">
.tiptap {
  min-height: v-bind("props.tiptapProps?.minHeight");
  box-shadow: none;
  resize: vertical;

  :first-child {
    margin-top: 0;
  }

  h1,
  h2,
  h3 {
    text-wrap: pretty;
    font-weight: 500;
  }

  p {
    font-size: 0.9rem;
  }

  h1 {
    font-size: 1.125rem;
  }

  h2 {
    font-size: 1rem;
  }

  h3 {
    font-size: 0.9rem;
  }

  ul {
    list-style-type: disc;
  }

  ol {
    list-style-type: decimal;
  }

  ul,
  ol {
    padding: 0 1rem;
    margin: 0.75rem 0.75rem 0.75rem 0.4rem;

    li p {
      margin-top: 0.25em;
      margin-bottom: 0.25em;
    }
  }
}
</style>
