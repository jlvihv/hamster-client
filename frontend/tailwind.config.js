// tailwindcss.config.js not support ES module import,
// we need use a const here to instead of import.
// import { primaryColor } from './src/settings/themeSetting';
const primaryColor = '#db730b';

module.exports = {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        primary: primaryColor,
      },
    },
  },
};
