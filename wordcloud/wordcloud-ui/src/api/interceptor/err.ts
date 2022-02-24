import { useMessage } from '/@/hooks/web/useMessage';
import { useI18n } from '/@/hooks/web/useI18n';
const { createErrorModal } = useMessage();

export const err = (error: any) => {
  if (error.response) {
    const { t } = useI18n();
    if (error.response.status === 403) {
      createErrorModal({
        title: t('sys.api.errorTip'),
        content: t('sys.api.errMsg403') || t('sys.api.networkExceptionMsg'),
        getContainer: () => document.body,
      });
    } else if (error.response.status === 401) {
      createErrorModal({
        title: t('sys.api.errorTip'),
        content: t('sys.api.errMsg401') || t('sys.api.networkExceptionMsg'),
        getContainer: () => document.body,
      });
      setTimeout(() => {
        window.location.href = `${import.meta.env.VITE_APP_LOGIN_PAGE}`;
      }, 2000);
    } else {
      createErrorModal({
        title: t('sys.api.errorTip'),
        content: error.response.data || t('sys.api.networkExceptionMsg'),
        getContainer: () => document.body,
      });
    }
  }
  return Promise.reject(error);
};
