<template>
  <div class="tag">
    <ag-grid-vue
      :rowData="rowData"
      :columnDefs="colDefs"
      style="height: 500px"
      class="ag-theme-quartz"
    >
    </ag-grid-vue>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import "ag-grid-community/styles/ag-grid.css"; // Mandatory CSS required by the Data Grid
import "ag-grid-community/styles/ag-theme-quartz.css"; // Optional Theme applied to the Data Grid
import { AgGridVue } from "ag-grid-vue3"; // Vue Data Grid Component

const rowData = ref([
  { make: "Tesla", model: "Model Y", price: 64950, electric: true },
  { make: "Ford", model: "F-Series", price: 33850, electric: false },
  { make: "Toyota", model: "Corolla", price: 29600, electric: false },
]);

// Column Definitions: Defines the columns to be displayed.
const colDefs = ref([
  { field: "make" },
  { field: "model" },
  { field: "price" },
  { field: "electric" },
]);
</script>

<style scoped lang="scss">
.tag {
  flex: 1;
  padding: 40px;
}
</style>
