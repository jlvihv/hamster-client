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
