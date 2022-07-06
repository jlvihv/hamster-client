import { resultSuccess, resultError, getTimestamps } from '../../helper';
import * as Graph from '/@wails/go/app/Graph';

const graphInfo = [
  {
    nodeEthereumUrl: 'string',
    ethereumUrl: 'string',
    ethereumNetwork: 'string',
    indexerAddress: 'string',
    mnemonic: 'string',
    application: [
      {
        id: 1,
        name: 'Example',
        describe: 'I am an example record',
        status: 0,
        ...getTimestamps(),
      },
    ],
    applicationId: 2,
  },
];

export default {
  QueryApplyAndParams: (id: number) => {
    const foundGraphInfo = graphInfo.find((x) => x.applicationId === id);
    return foundGraphInfo ? resultSuccess(graphInfo) : resultError('Infomation record not found');
  },
} as Partial<typeof Graph>;
