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
      <BasicTable
        :columns="columns"
        :dataSource="data"
        :canResize="canResize"
        :loading="loading"
        :striped="striped"
        :bordered="border"
        showTableSetting
        :pagination="pagination"
        @columns-change="handleColumnChange"
      />
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
  import { BasicColumn, BasicTable, ColumnChangeParam } from '/@/components/Table';
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
        title: '姓名',
        dataIndex: 'name',
        width: 150,
        filters: [
          { text: 'Male', value: 'male' },
          { text: 'Female', value: 'female' },
        ],
      },
      {
        title: '地址',
        dataIndex: 'address',
      },
      {
        title: '编号',
        dataIndex: 'no',
        width: 150,
        sorter: true,
        defaultHidden: true,
      },
      {
        title: '开始时间',
        width: 150,
        sorter: true,
        dataIndex: 'beginTime',
      },
      {
        title: '结束时间',
        width: 150,
        sorter: true,
        dataIndex: 'endTime',
      },
    ];
  }

  function getBasicData() {
    const data: any = (() => {
      const arr: any = [];
      for (let index = 0; index < 40; index++) {
        arr.push({
          id: `${index}`,
          name: 'John Brown',
          age: `1${index}`,
          no: `${index + 10}`,
          address: 'New York No. 1 Lake ParkNew York No. 1 Lake Park',
          beginTime: new Date().toLocaleString(),
          endTime: new Date().toLocaleString(),
        });
      }
      return arr;
    })();
    return data;
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

      const canResize = ref(false);
      const loading = ref(false);
      const striped = ref(true);
      const border = ref(true);
      const pagination = ref<any>(false);
      function toggleCanResize() {
        canResize.value = !canResize.value;
      }
      function toggleStriped() {
        striped.value = !striped.value;
      }
      function toggleLoading() {
        loading.value = true;
        setTimeout(() => {
          loading.value = false;
          pagination.value = { pageSize: 20 };
        }, 3000);
      }
      function toggleBorder() {
        border.value = !border.value;
      }

      function handleColumnChange(data: ColumnChangeParam[]) {
        console.log('ColumnChanged', data);
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

        columns: getBasicColumns(),
        data: getBasicData(),
        canResize,
        loading,
        striped,
        border,
        toggleStriped,
        toggleCanResize,
        toggleLoading,
        toggleBorder,
        pagination,
        handleColumnChange,
      };
    },
  });
</script>
