<template>
  <div class="p-4">
    <h1>{{ t('sys.app.underConstruction') }}</h1>
  </div>
</template>
<script lang="ts" setup>
  import { ref } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  const loading = ref(true);
  const { t } = useI18n();

  setTimeout(() => {
    loading.value = false;
  }, 1500);
</script>
