<script setup lang="ts">
import type { LikeEntity } from "@/types";
import { likeRepository } from "~/repositories/like.repository";

type Props = {
  entity: LikeEntity;
  entityId: string;
};

const props = defineProps<Props>();
const authAction = useAuthAction();
const { likeEntity } = likeRepository();

async function likeHandler() {
  authAction.run(async () => {
    try {
      await likeEntity(props.entity, props.entityId);
    } catch (err) {
      console.log(getErrorMessage(err));
    }
  });
}
</script>

<template>
  <button @click="likeHandler">
    <Icon name="heroicons:heart" class="w-6 h-6" />
  </button>
</template>
