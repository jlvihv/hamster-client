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

export function nodePolyfillAlias() {
  const polyfills = {
    url: 'url',
    stream: 'stream',
    assert: 'assert',
    querystring: 'qs',
    // buffer: 'buffer-es6',
    // util: 'util',
    // sys: 'util',
    // events: 'events',
    // path: 'path',
    // punycode: 'punycode',
    // string_decoder: 'string-decoder',
    // http: 'http',
    // https: 'http',
    // os: 'os',
    // constants: 'constants',
    // _stream_duplex: 'readable-stream/duplex',
    // _stream_passthrough: 'readable-stream/passthrough',
    // _stream_readable: 'readable-stream/readable',
    // _stream_writable: 'readable-stream/writable',
    // _stream_transform: 'readable-stream/transform',
    // timers: 'timers',
    // console: 'console',
    // vm: 'vm',
    // zlib: 'zlib',
    // tty: 'tty',
    // domain: 'domain',
  };

  const aliases: Recordable = {};

  Object.entries(polyfills).forEach(([name, lib]) => {
    aliases[name] = `rollup-plugin-node-polyfills/polyfills/${lib}`;
  });

  return aliases;
}
