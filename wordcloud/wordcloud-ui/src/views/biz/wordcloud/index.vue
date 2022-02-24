<template>
  <div class="p-4">
    <h1>词云图建设中</h1>
  </div>
</template>
<script lang="ts" setup>
import { ref } from 'vue';
const loading = ref(true);

setTimeout(() => {
  loading.value = false;
}, 1500);
</script>
