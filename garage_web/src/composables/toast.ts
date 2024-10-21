import { ref, onMounted, onUnmounted } from "vue";

export function useMouse() {
  const x = ref(0);
  const y = ref(0);

  onMounted(() => {});
  onUnmounted(() => {});

  return { x, y };
}
