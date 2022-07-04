import { resultSuccess, resultPageSuccess, resultError } from '../../helper';
import * as Application from '/@wails/go/app/Application';
import { isEqualWith, lowerCase, remove } from 'lodash-es';

const applications = [
  { id: 1, name: 'Example', abbreviation: 'ex', describe: 'I am an example record', status: 0 },
  { id: 2, name: 'Test', abbreviation: 'test', describe: 'I am an test record', status: 1 },
];

let idCursor = applications.length;

export default {
  ApplicationList: (page: number, pageSize: number, name: string, status: number) => {
    let data = applications;
    if (name) data = applications.filter((x) => isEqualWith(x.name, name, lowerCase));
    if (status) data = applications.filter((x) => x.status === status);

    return resultPageSuccess(page, pageSize, data);
  },
  AddApplication: (application) => {
    idCursor++;

    const newApplication = { id: idCursor, ...application };
    applications.push(newApplication);

    return resultSuccess(newApplication);
  },
  UpdateApplication: (application) => {
    const { id } = application as unknown as any;
    const foundApplication = applications.find((x) => x.id === id);

    if (!foundApplication) {
      return resultError('Record not found');
    }

    // Update application
    Object.assign(foundApplication, application);

    return resultSuccess(foundApplication);
  },
  DeleteApplication: (id) => {
    remove(applications, (x) => x.id === id);
    return resultSuccess();
  },
  QueryApplicationById: (id) => {
    const foundApplication = applications.find((x) => x.id === id);
    return resultSuccess(foundApplication);
  },
} as Partial<typeof Application>;
