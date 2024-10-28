<template>
  <template v-if="movieStore.selectMovieInfo">
    <div class="movie-cover" v-if="movieStore.selectMovieInfo">
      <img
        :src="coverImageUrl"
        alt="movie cover"
        onerror="this.src='/img_not_fount.svg'; this.onerror=null;"
      />
    </div>
    <div class="movie-info">
      <input class="input" v-model="movieStore.selectMovieInfo.movie.code"></input>
      <span>{{ movieStore.selectMovieInfo.movie.title }}</span>
      <span>{{ movieStore.selectMovieInfo.movie.publishDate }}</span>
    </div>
    <div class="movie-actor">
      <template v-for="i in movieStore.selectMovieInfo.actors">
        <span>{{ i.actorName }}</span>
      </template>
    </div>

    <div class="movie-tag" ref="movieTagDivRef">
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

onClickOutside(
  tagTreeRef,
  (_event) => {
    if (showTagTree.value) {
      showTagTree.value = false;
    }
  },
  {
    ignore: [movieTagDivRef],
  }
);

watchEffect(() => {
  if (movieTagRef.value == null) return;
  tagify = new Tagify(movieTagRef.value, {
    userInput: false,
    hooks: {},
    dropdown: {},
    templates: {},
  });
  const tags: Array<TagData> = [];

  if (
    movieStore.selectMovieInfo &&
    movieStore.selectMovieInfo.tags?.length > 0
  ) {
    movieStore.selectMovieInfo?.tags.forEach((tag) => {
      let data: TagData = {
        value: tag.tag_name,
        editable: false,
        source: tag,
      };
      if (tag.tag_name) tags.push(data);
    });

    tagify.addTags(tags);
  }

  tagify
    .on("remove", (_e) => {})
    .on("add", (_e) => {})
    .on("focus", (_e) => {
      console.log("foucs");
      onMovieTagClick();
    })
    .on("dropdown:show", (_e) => {
      console.log("dropdown:show");
      // tagify?.dropdown.hide();
    })
    .on("input", (_e) => {});
});

function onMovieTagClick() {
  if (showTagTree.value) return;
  if (!movieTagDivRef.value || !tagTreeRef.value) return;

  computePosition(movieTagDivRef.value, tagTreeRef.value.$el, {
    placement: "left",
    middleware: [offset({ mainAxis: 250, crossAxis: -20 }), flip(), shift()],
    strategy: "fixed",
  }).then(({ x, y }) => {
    showTagTree.value = true;
    tagTreePosition.value = { x, y };
  });
}
</script>

<style scoped lang="scss">
.movie-cover {
  img {
    max-width: 300px;
    max-height: 300px;
  }
}
.movie-info {
  display: flex;
  flex-direction: column;
}

.movie-tag {
  width: fit-content;
}
</style>
