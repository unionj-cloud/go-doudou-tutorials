<template>
  <PageWrapper :title="t('sys.wordcloud.pageTitle')" contentFullHeight>
    <CollapseContainer :title="t('sys.wordcloud.containerTitle')">
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
    <CollapseContainer :title="t('sys.wordcloud.tableTitle')">
      <BasicTable @register="registerTable">
        <template #toolbar>
          <a-button type="primary" @click="handleReloadCurrent">
            {{ t('sys.table.refreshCurrentPage') }}
          </a-button>
          <a-button type="primary" @click="handleReload">
            {{ t('sys.table.refreshCurrentPageReturnFirst') }}
          </a-button>
        </template>
        <template #srcUrl="{ record }">
          <a :href="filefix(record.srcUrl)"> {{ filenamefix(record.srcUrl) }} </a>
        </template>
        <template #imgUrl="{ record }">
          <a :href="filefix(record.imgUrl)" target="_blank">
            <img :src="filefix(record.imgUrl)" />
          </a>
        </template>
        <template #lang="{ record }">
          {{ langfix(record.lang) }}
        </template>
        <template #status="{ record }">
          {{ statusfix(record.status) }}
        </template>
      </BasicTable>
    </CollapseContainer>
  </PageWrapper>
</template>
<script lang="ts">
  import { useI18n } from '/@/hooks/web/useI18n';
  import { defineComponent, ref, unref, computed } from 'vue';
  import { BasicForm, FormSchema } from '/@/components/Form/index';
  import { BasicUpload } from '/@/components/Upload';
  import { CollapseContainer } from '/@/components/Container';
  import { PageWrapper } from '/@/components/Page';
  import { uploadApi } from '/@/api/sys/upload';
  import { useGlobSetting } from '/@/hooks/setting';
  import { BasicColumn, BasicTable, useTable } from '/@/components/Table';
  const { bffUrl = '', ossUrl = '' } = useGlobSetting();
  import { demoListApi } from '/@/api/demo/table';

  export default defineComponent({
    components: { BasicForm, CollapseContainer, PageWrapper, BasicUpload, BasicTable },
    setup() {
      const { t } = useI18n();
      const langRef = ref('zh');

      const schemas: FormSchema[] = [
        {
          field: 'commentType',
          component: 'RadioGroup',
          label: t('sys.wordcloud.lang'),
          colProps: {
            span: 8,
          },
          defaultValue: 'zh',
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
      const columns: BasicColumn[] = [
        {
          title: 'ID',
          dataIndex: 'id',
          fixed: 'left',
        },
        {
          title: t('sys.wordcloud.inputFile'),
          dataIndex: 'srcUrl',
          slots: { customRender: 'srcUrl' },
        },
        {
          title: t('sys.wordcloud.lang'),
          dataIndex: 'lang',
          slots: { customRender: 'lang' },
        },
        {
          title: t('sys.wordcloud.imgUrl'),
          dataIndex: 'imgUrl',
          slots: { customRender: 'imgUrl' },
        },
        {
          title: t('sys.wordcloud.status'),
          dataIndex: 'status',
          slots: { customRender: 'status' },
        },
        {
          title: t('sys.wordcloud.error'),
          dataIndex: 'error',
        },
        {
          title: t('sys.wordcloud.createAt'),
          dataIndex: 'createAt',
        },
      ];
      const [registerTable, { reload }] = useTable({
        api: demoListApi,
        columns,
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
        t,
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

    methods: {
      filefix: (input: string) => {
        if (!input) {
          return '';
        }
        return input.replaceAll(new URL(input).origin, ossUrl);
      },
      filenamefix: (input: string) => {
        if (!input) {
          return '';
        }
        return input.slice(input.lastIndexOf('/') + 1);
      },
      statusfix: function (input: number) {
        let ret = '';
        switch (input) {
          case 0:
            ret = this.t('sys.wordcloud.waiting');
            break;
          case 1:
            ret = this.t('sys.wordcloud.doing');
            break;
          case 2:
            ret = this.t('sys.wordcloud.success');
            break;
          case 3:
            ret = this.t('sys.wordcloud.fail');
            break;
          default:
            ret = this.t('sys.wordcloud.unknown');
        }
        return ret;
      },
      langfix: function (input: string) {
        let ret = '';
        switch (input) {
          case 'zh':
            ret = this.t('sys.wordcloud.chinese');
            break;
          case 'en':
            ret = this.t('sys.wordcloud.english');
            break;
          default:
            ret = this.t('sys.wordcloud.unknown');
        }
        return ret;
      },
    },
  });
</script>
