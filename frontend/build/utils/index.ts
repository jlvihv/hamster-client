import glob from 'glob';
import path from 'path';

// Read all environment variable configuration files to process.env
export function wrapperEnv(envConf: Recordable): ViteEnv {
  const ret: any = {};

  for (const envName of Object.keys(envConf)) {
    let value = envConf[envName];
    value = { true: true, false: false }[value] || value;

    if (envName === 'VITE_PORT') {
      value = Number(value);
    }

    ret[envName] = value;
    if (typeof value === 'string') {
      process.env[envName] = value;
    } else if (typeof value === 'object') {
      process.env[envName] = JSON.stringify(value);
    }
  }

  return ret;
}

export function packedFileAlias(isBuild: boolean) {
  const alias = {};

  if (isBuild) {
    const packsFolder = path.resolve(__dirname, '../../src/packs');
    const distFolder = path.resolve(__dirname, '../webpack/dist');

    glob.sync('/**/*.{js,ts}', { root: packsFolder }).map((file) => {
      const filePath = path.parse(file);
      const libName = filePath.name.replace('--', '/');
      alias[libName] = path.resolve(distFolder, `${filePath.name}.js`);
    });
  }

  return alias;
}
