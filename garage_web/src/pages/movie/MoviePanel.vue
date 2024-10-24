<template>
  <template v-if="movieStore.selectMovieInfo">
    <div class="movie-cover" v-if="movieStore.selectMovieInfo">
      <img :src="coverImageUrl" alt="movie cover" />
    </div>
    <div class="movie-info">
      <span>{{ movieStore.selectMovieInfo.movie.code }}</span>
      <span>{{ movieStore.selectMovieInfo.movie.title }}</span>
      <span>{{ movieStore.selectMovieInfo.movie.publishDate }}</span>
    </div>
    <div class="movie-actor">
      <template v-for="i in movieStore.selectMovieInfo.actors">
        <span>{{ i.actorName }}</span>
      </template>
    </div>

    <div class="movie-tag" ref="movieTagDivRef" @click.stop="onMovieTagClick">
      <input ref="movieTagRef" />
    </div>
  </template>
  <TagTree :show="showTagTree" :position="tagTreePosition" ref="tagTreeRef" />
</template>

<script setup lang="ts">
import { Ref, ref, useTemplateRef, watchEffect } from "vue";
import { onClickOutside } from "@vueuse/core";
import { useMovieStore } from "@/stores/movie";
import TagTree from "@/pages/TagTree.vue";
import Tagify, { TagData } from "@yaireo/tagify";
import "@yaireo/tagify/dist/tagify.css";
import {
  computePosition,
  flip,
  shift,
  offset,
  autoPlacement,
  platform,
  inline,
} from "@floating-ui/vue";

const movieStore = useMovieStore();
const coverImageUrl = ref("");
const movieTagRef = useTemplateRef<HTMLInputElement>("movieTagRef");
const movieTagDivRef = useTemplateRef<HTMLDivElement>("movieTagDivRef");
const tagTreeRef = useTemplateRef("tagTreeRef");
const showTagTree = ref(false);
const tagTreePosition = ref({ x: 0, y: 0 });

let tagify: Tagify | null = null;

movieStore.$subscribe((_mutation, state) => {
  coverImageUrl.value = `/api/v1/img/${state.selectMovieInfo?.movie.id}/${state.selectMovieInfo?.movie.cover}`;
});

onClickOutside(tagTreeRef, (_event) => {
  if (showTagTree.value) {
    showTagTree.value = false;
  }
});

watchEffect(() => {
  if (movieTagRef.value == null) return;
  tagify = new Tagify(movieTagRef.value, {
    userInput: false,
  });
  const tags: Array<TagData> = [];

  movieStore.selectMovieInfo?.tags.forEach((tag) => {
    let data: TagData = {
      value: tag.tag_name,
      editable: false,
      source: tag,
    };
    if (tag.tag_name) tags.push(data);
  });

  tagify.addTags(tags);

  tagify
    .on("remove", (e) => {
      console.log(e);
      console.log(movieTagRef.value?.value);
    })
    .on("add", (e) => {
      console.log(e);
    })
    .on("focus", (_e) => {
      console.log("focus");
    });
});

function onMovieTagClick() {
  if (!movieTagDivRef.value || !tagTreeRef.value) return;
  const empty = document.createElement("div");
  computePosition(movieTagDivRef.value, empty, {
    placement: "left",
    middleware: [offset(30), flip(), shift()],
    strategy: "fixed",
  }).then(({ x, y }) => {
    showTagTree.value = true;
    tagTreePosition.value = { x, y };
  });
}
</script>

<style scoped lang="scss">
.movie-tag {
  width: fit-content;
}
</style>
