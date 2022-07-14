import { resultSuccess } from '../../helper';
import * as Deploy from '/@wails/go/app/Deploy';

const deployInfoCollection = {};

export default {
  GetDeployInfo(applicationId: number) {
    return resultSuccess({ id: applicationId, data: deployInfoCollection[applicationId] });
  },
  SaveDeployInfo(applicationId: number, data: any) {
    deployInfoCollection[applicationId] = data;
    return resultSuccess(true);
  },
  DeployTheGraph(applicationId: number, data: any) {
    console.log('DeployTheGraph', applicationId, data);
    return resultSuccess(true);
  },
} as Partial<typeof Deploy>;
