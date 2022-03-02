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
        :maxSize="1"
        :maxNumber="1"
        @change="handleChange"
        :api="uploadApi"
        :uploadParams="uploadParams"
      />
      <br />
      <img v-if="downloadUrl" :src="downloadUrl" />
    </CollapseContainer>
    <br />
    <CollapseContainer :title="tableTitle">
      <BasicTable @register="registerTable">
        <template #toolbar>
          <a-button type="primary" @click="handleReloadCurrent"> 刷新当前页 </a-button>
          <a-button type="primary" @click="handleReload"> 刷新并返回第一页 </a-button>
        </template>
      </BasicTable>
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
  import { PageWrapper } from '/@/components/Page';
  import { uploadApi } from '/@/api/sys/upload';
  import { useGlobSetting } from '/@/hooks/setting';
  import { BasicColumn, BasicTable, useTable } from '/@/components/Table';
  const { bffUrl = '', ossUrl = '' } = useGlobSetting();

  function getBasicColumns(): BasicColumn[] {
    return [
      {
        title: 'ID',
        dataIndex: 'id',
        fixed: 'left',
        width: 200,
      },
      {
        title: '上传文件',
        dataIndex: 'srcUrl',
        width: 150,
      },
      {
        title: '语种',
        dataIndex: 'lang',
        width: 150,
      },
      {
        title: '词云图',
        dataIndex: 'imgUrl',
      },
      {
        title: '状态',
        dataIndex: 'status',
        width: 150,
      },
      {
        title: '错误信息',
        width: 150,
        dataIndex: 'error',
      },
      {
        title: '创建时间',
        width: 150,
        dataIndex: 'createAt',
      },
    ];
  }

  export default defineComponent({
    components: { BasicForm, CollapseContainer, PageWrapper, BasicUpload, BasicTable },
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

      const uploadParams = computed<Recordable>(() => {
        return {
          url: bffUrl + '/upload',
          lang: unref(langRef),
        };
      });

      const downloadUrl = ref('');
      const pageTitle = t('sys.wordcloud.pageTitle');
      const containerTitle = t('sys.wordcloud.containerTitle');
      const tableTitle = t('sys.wordcloud.tableTitle');

      const [registerTable, { reload }] = useTable({
        api: demoListApi,
        columns: getBasicColumns(),
        pagination: { pageSize: 10 },
      });
      function handleReloadCurrent() {
        reload();
      }

      function handleReload() {
        reload({
          page: 1,
        });
      }

      return {
        schemas,
        uploadParams,
        uploadApi,
        downloadUrl,
        pageTitle,
        containerTitle,
        tableTitle,
        handleChange: (list: string[]) => {
          const imgUrl = list.length > 0 ? list[list.length - 1] : '';
          if (imgUrl) {
            downloadUrl.value = imgUrl.replaceAll(new URL(imgUrl).origin, ossUrl);
          }
        },
        registerTable,
        handleReloadCurrent,
        handleReload,
      };
    },
  });
</script>
