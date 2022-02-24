import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const dashboard: AppRouteModule = {
  path: '/wordcloud',
  name: 'Wordcloud',
  component: LAYOUT,
  redirect: '/wordcloud/index',
  meta: {
    hideChildrenInMenu: true,
    icon: 'simple-icons:icloud',
    title: t('routes.dashboard.wordcloud'),
    orderNo: 100000,
  },
  children: [
    {
      path: 'index',
      name: 'WordcloudPage',
      component: () => import('/@/views/biz/wordcloud/index.vue'),
      meta: {
        title: t('routes.dashboard.wordcloud'),
        icon: 'simple-icons:icloud',
        hideMenu: true,
      },
    },
  ],
};

export default dashboard;
