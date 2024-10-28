<template>
  <div
    ref="target"
    class="tag-tree"
    :style="{
      top: position.y + 'px',
      left: position.x + 'px',
      display: show ? 'block' : 'none',
    }"
  >
    <Draggable
      v-model="tagList"
      ref="tree"
      virtualization
      class="mtl-tree"
      treeLine
      updateBehavior="modify"
    >
      <template #default="{ node, stat }: { node: TreeTag, stat: Stat<Tag> }">
        <i
          v-if="!stat.open && node.children.length > 0"
          @click.native="stat.open = !stat.open"
          class="pi pi-angle-right"
          style="font-size: 1rem"
        ></i>
        <i
          v-if="stat.open && node.children.length > 0"
          @click.native="stat.open = !stat.open"
          class="pi pi-angle-down"
          style="font-size: 1rem"
        ></i>

        <div
          class="tag"
          :class="{ check: stat.checked }"
          @click.native="stat.checked = !stat.checked"
        >
          <i class="pi pi-tag" style="font-size: 1rem"></i>
          <span class="mtl-ml">{{ node.name }} {{ stat.checked }}</span>
        </div>
      </template>
    </Draggable>
  </div>
</template>

<script setup lang="ts">
import { ref, Ref, watchEffect } from "vue";
import { Tag, GetTag } from "@/api/tag";
import { Draggable } from "@he-tree/vue";
import { Stat } from "@he-tree/tree-utils";
import "@he-tree/vue/style/default.css";
import "@he-tree/vue/style/material-design.css";

interface TreeTag {
  id: number;
  name: string;
  children: TreeTag[];
}

const tree = ref<InstanceType<typeof Draggable>>();
const tagList: Ref<Array<TreeTag>> = ref([]);
const { show, position = { x: 0, y: 0 } } = defineProps<{
  show?: boolean;
  position?: {
    x: number;
    y: number;
  };
}>();

watchEffect(async () => {
  if (show) {
    await GetTag().then((d) => {
      let f = convertToTree(d.data.data);
      tree.value?.addMulti(f);
      checkTags();
    });
  } else {
    console.log("remove data");
    console.log(tree.value?.getData());
    console.log(tagList.value);
    tagList.value = [];
    // tree.value?.removeMulti(tree.value?.getData());
    // tree.value?.remove(tree.value?.getData()[0]);
  }
});

function convertToTree(tags: Array<Tag>): Array<TreeTag> {
  const map: { [key: number]: TreeTag } = {};
  const roots: Array<TreeTag> = [];

  tags.forEach((tag) => {
    map[tag.id] = {
      id: tag.id,
      name: tag.name,
      children: [],
    };
  });

  tags.forEach((tag) => {
    if (tag.pid === 0) {
      roots.push(map[tag.id]);
    } else {
      if (map[tag.pid]) {
        map[tag.pid].children.push(map[tag.id]);
      }
    }
  });
  return roots;
}

function checkTags() {}
</script>

<style scoped lang="scss">
.tag-tree {
  position: fixed;
  background-color: #ffffff;
  box-shadow: 0px 1px 5px 1px rgba(0, 0, 0, 0.1);
  border-radius: 5px;
  padding: 10px;
  min-width: 240px;
  min-height: 320px;
  max-width: 240px;
  max-height: 320px;
  z-index: 99;
  overflow: auto;
  .mtl-tree {
    .tag {
      padding: 5px;
      border-radius: 5px;
      display: flex;
      align-items: center;
      &.check {
        background: #8f8ff3;
      }
    }
  }
}
</style>
