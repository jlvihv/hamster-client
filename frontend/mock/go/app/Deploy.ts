import { resultSuccess } from '../../helper';
import * as Deploy from '/@wails/go/app/Deploy';

const deployInfoCollection = {};

export default {
  GetDeployInfo(applicationId: number) {
    const jsonString = deployInfoCollection[applicationId];
    const json = jsonString && JSON.parse(jsonString);

    return resultSuccess(json);
  },
  SaveDeployInfo(applicationId: number, data: string) {
    deployInfoCollection[applicationId] = data;
    return resultSuccess(true);
  },
  DeployTheGraph(applicationId: number, data: string) {
    console.log('DeployTheGraph', applicationId, data);
    return resultSuccess(true);
  },
} as Partial<typeof Deploy>;
