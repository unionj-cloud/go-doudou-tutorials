import { UploadApiResult } from './model/uploadModel';
import { defHttp } from '/@/utils/http/axios';
import { UploadFileParams } from '/#/axios';
import { useGlobSetting } from '/@/hooks/setting';

const { uploadUrl = '' } = useGlobSetting();

/**
 * @description: Upload interface
 */
export function uploadApi(
  params: UploadFileParams,
  onUploadProgress: (progressEvent: ProgressEvent) => void,
) {
  let url = '';
  if (params.data) {
    url = params.data.url;
    if (!url) {
      url = uploadUrl;
    }
    delete params.data.url;
  }
  return defHttp.uploadFile<UploadApiResult>(
    {
      url,
      onUploadProgress,
    },
    params,
  );
}
