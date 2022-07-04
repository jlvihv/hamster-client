import { set } from 'lodash-es';

const modules = import.meta.globEager('./go/**/*.ts');
const mockModules: { name: string; module: Object }[] = [];
Object.keys(modules).forEach((key) => {
  const filename = key.split(/(\\|\/)/g).pop();
  const resourceName = filename?.split('.')[0];

  if (resourceName) {
    mockModules.push({ name: resourceName, module: modules[key].default });
  }
});

const setupApiMock = () => {
  mockModules.forEach(({ name, module }) => {
    // Set APIs in window namespace
    set(window, `go.app.${name}`, module);
  });
};

setupApiMock();
export default setupApiMock;
