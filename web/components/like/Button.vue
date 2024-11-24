<script setup lang="ts">
import type { LikeEntity } from "@/types";
import { likeRepository } from "@/repositories/like.repository";

type Props = {
  liked: boolean;
  entity: LikeEntity;
  entityId: string;
};

const props = defineProps<Props>();
const authAction = useAuthAction();
const toast = useToast();
const { likeEntity, unlikeEntity } = likeRepository();
const { t } = useI18n();
const loading = ref(false);

const optimisticLike = ref(props.liked);

async function likeHandler() {
  authAction.run(async () => {
    if (loading.value) return;
    loading.value = true;

    try {
      optimisticLike.value = !optimisticLike.value;

      if (optimisticLike.value) {
        await likeEntity(props.entity, props.entityId);
      } else {
        await unlikeEntity(props.entity, props.entityId);
      }
    } catch (err) {
      toast.add({
        color: "yellow",
        title: t(
          optimisticLike.value
            ? "writing.likes.likeFailed"
            : "writing.likes.unlikeFailed"
        ),
      });
      optimisticLike.value = !optimisticLike.value;
    } finally {
      loading.value = false;
    }
  });
}

const liked = computed(() => optimisticLike.value);
</script>

<template>
  <button @click="likeHandler" aria-label="Toggle like" :disabled="loading">
    <Icon v-show="liked" name="heroicons:heart-solid" class="w-6 h-6" />
    <Icon
      v-show="!liked"
      name="heroicons:heart"
      class="w-6 h-6 hover:text-gray-600 hover:dark:text-gray-400"
    />
  </button>
</template>
