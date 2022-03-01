<template>
  <PageWrapper :title="pageTitle" contentFullHeight>
    <CollapseContainer :title="containerTitle">
      <BasicForm
        autoFocusFirstItem
        :labelWidth="200"
        :schemas="schemas"
        :actionColOptions="{ span: 24 }"
        :showActionButtonGroup="false"
      />

      <BasicUpload
        :maxSize="100"
        :maxNumber="100"
        @change="handleChange"
        :api="uploadApi"
        :uploadParams="uploadParams"
      />
      <br />
      <a v-if="downloadUrl" :href="downloadUrl"><a-button type="success">下载</a-button></a>
    </CollapseContainer>
  </PageWrapper>
</template>
<script lang="ts">
  import { useI18n } from '/@/hooks/web/useI18n';
  const { t } = useI18n();
  import { defineComponent, ref, unref, computed } from 'vue';
  import { BasicForm, FormSchema } from '/@/components/Form/index';
  import { BasicUpload } from '/@/components/Upload';
  import { CollapseContainer } from '/@/components/Container';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { PageWrapper } from '/@/components/Page';
  import { uploadApi } from '/@/api/sys/upload';
  import { useGlobSetting } from '/@/hooks/setting';

  const { bffUrl = '' } = useGlobSetting();

  export default defineComponent({
    components: { BasicForm, CollapseContainer, PageWrapper, BasicUpload },
    setup() {
      const langRef = ref(0);

      const schemas: FormSchema[] = [
        {
          field: 'commentType',
          component: 'RadioGroup',
          label: t('sys.wordcloud.lang'),
          colProps: {
            span: 8,
          },
          componentProps: {
            options: [
              {
                label: '中文',
                value: 'zh',
              },
              {
                label: 'English',
                value: 'en',
              },
            ],
            onChange: (e: any) => {
              console.log(e.target.value);
              langRef.value = e.target.value;
            },
          },
        },
      ];

      const { createMessage } = useMessage();
      const uploadParams = computed<Recordable>(() => {
        return {
          url: bffUrl + '/upload',
          lang: unref(langRef),
        };
      });

      const downloadUrl = ref('');
      const pageTitle = t('sys.wordcloud.pageTitle');
      const containerTitle = t('sys.wordcloud.containerTitle');
      return {
        schemas,
        uploadParams,
        uploadApi,
        downloadUrl,
        pageTitle,
        containerTitle,
        handleChange: (list: string[]) => {
          createMessage.info(`已上传文件${JSON.stringify(list)}`);
          debugger
          downloadUrl.value = list.length > 0 ? list[list.length - 1] : '';
        },
      };
    },
  });
</script>
